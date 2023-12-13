package main

import (
	"context"
	"flag"
	"fmt"
	"golang.conradwood.net/apis/common"
	"golang.conradwood.net/go-easyops/authremote"
	"golang.conradwood.net/go-easyops/errors"
	"golang.conradwood.net/go-easyops/server"
	"golang.conradwood.net/go-easyops/utils"
	"golang.yacloud.eu/apis/session"
	pb "golang.yacloud.eu/apis/sessionmanager"
	"golang.yacloud.eu/sessionmanager/db"
	"google.golang.org/grpc"
	"net"
	"os"
	"sort"
	"time"
)

const (
	// if request from same ip for same user, it's not really noteworthy
	MAX_TIME_ON_IP = time.Duration(24) * time.Hour
)

var (
	port = flag.Int("port", 4100, "The grpc server port")
)

type echoServer struct {
}

func main() {
	var err error
	flag.Parse()
	fmt.Printf("Starting SessionManagerServer...\n")
	db.DefaultDBSessionLog()
	go session_cleaner()
	sd := server.NewServerDef()
	sd.SetPort(*port)
	sd.SetRegister(server.Register(
		func(server *grpc.Server) error {
			e := new(echoServer)
			pb.RegisterSessionManagerServer(server, e)
			return nil
		},
	))
	err = server.ServerStartup(sd)
	utils.Bail("Unable to start server", err)
	os.Exit(0)
}

/************************************
* grpc functions
************************************/

func (e *echoServer) Ping(ctx context.Context, req *common.Void) (*pb.PingResponse, error) {
	resp := &pb.PingResponse{Response: "pingresponse"}
	return resp, nil
}
func (e *echoServer) NewSession(ctx context.Context, req *pb.NewSessionRequest) (*pb.SessionResponse, error) {
	sls, err := db.DefaultDBSessionLog().ByUserID(ctx, req.UserID)
	if err != nil {
		return nil, err
	}
	sort.Slice(sls, func(i, j int) bool {
		return sls[i].Created >= sls[j].Created
	})
	ip := req.IPAddress
	host, _, err := net.SplitHostPort(ip)
	if err != nil {
		fmt.Printf("Invalid ip address?: %s\n", err)
	} else {
		ip = host
	}

	isnew := true
	for _, s := range sls {
		if s.BrowserID == req.BrowserID {
			isnew = false
			break
		}
		created := time.Unix(int64(s.Created), 0)
		if s.IP == ip && time.Since(created) < MAX_TIME_ON_IP {
			isnew = false
			break
		}
	}

	st := utils.RandomString(128)
	fmt.Printf("SessionRequest: %#v\n", req)
	now := uint32(time.Now().Unix())
	sl := &pb.SessionLog{
		UserID:       req.UserID,
		Username:     req.Username,
		Useremail:    req.Useremail,
		IP:           ip,
		UserAgent:    req.UserAgent,
		BrowserID:    req.BrowserID,
		Created:      now,
		SessionToken: st,
		LastUsed:     now,
		TriggerHost:  req.TriggerHost,
	}
	_, err = db.DefaultDBSessionLog().Save(ctx, sl)
	if err != nil {
		return nil, err
	}

	res := &pb.SessionResponse{
		Token:     st,
		NewDevice: isnew,
	}
	if len(sls) > 0 {
		res.LastSessionTimestamp = sls[0].Created
	}
	fmt.Printf("User %s session created, newdevice=%v, lastsession=%s\n", req.UserID, res.NewDevice, utils.TimestampString(res.LastSessionTimestamp))
	s, err := create_session_from_log(ctx, sl)
	if err != nil {
		return nil, err
	}
	res.Session = s
	fmt.Printf("Created session for user %s: %s\n", s.UserID, res.Token)
	return res, nil
}

func (e *echoServer) VerifySession(ctx context.Context, req *pb.SessionToken) (*pb.SessionVerifyResponse, error) {
	res := &pb.SessionVerifyResponse{
		IsValid: false,
	}
	sls, err := db.DefaultDBSessionLog().BySessionToken(ctx, req.Token)
	if err != nil {
		return nil, err
	}
	if len(sls) == 0 {
		return res, nil
	}
	sl := sls[0]
	res.IsSessionToken = true

	t := time.Unix(int64(sl.LastUsed), 0)
	if time.Since(t) > time.Duration(120)*time.Minute {
		return res, nil
	}
	res.SessionLog = sl
	res.IsValid = true

	now := uint32(time.Now().Unix())
	sl.LastUsed = now
	err = db.DefaultDBSessionLog().Update(ctx, sl)
	if err != nil {
		fmt.Printf("failed to update sessionlog: %s\n", err)
	}
	user, err := authremote.GetSignedUserByID(ctx, res.SessionLog.UserID)
	if err != nil {
		return nil, err
	}
	res.User = user
	return res, nil
}
func (e *echoServer) KeepAliveSession(ctx context.Context, req *pb.SessionToken) (*pb.SessionAliveResponse, error) {
	res := &pb.SessionAliveResponse{
		IsValid: false,
	}
	sl, err := get_sessionlog_by_sessionid(ctx, req.Token)
	if err != nil {
		return nil, err
	}
	if sl == nil {
		return nil, errors.InvalidArgs(ctx, "invalid sessiontoken", "invalid sessiontoken")
	}
	s, err := create_session_from_log(ctx, sl)
	if err != nil {
		return nil, err
	}
	res.Session = s
	res.IsValid = true
	return res, nil
}

func (e *echoServer) DisassociateUserFromSession(ctx context.Context, req *pb.SessionToken) (*pb.SessionAliveResponse, error) {
	res := &pb.SessionAliveResponse{
		IsValid: false,
	}
	sid := req.Token
	sl, err := get_sessionlog_by_sessionid(ctx, sid)
	if err != nil {
		return nil, err
	}
	if sl == nil {
		return res, nil
	}
	wasuser := sl.UserID
	sl.UserID = ""
	sl.LastUsed = uint32(time.Now().Unix())
	err = db.DefaultDBSessionLog().Update(ctx, sl)
	if err != nil {
		return nil, err
	}

	s, err := create_session_from_log(ctx, sl)
	if err != nil {
		return nil, err
	}
	res.IsValid = true
	res.Session = s
	fmt.Printf("Disassociated user \"%s\" from session %s\n", wasuser, sid)
	return res, nil
}

// update a session, e.g. with a new user
func (e *echoServer) User2Session(ctx context.Context, req *pb.User2SessionRequest) (*pb.SessionAliveResponse, error) {
	res := &pb.SessionAliveResponse{
		IsValid: false,
	}
	sid := req.Session.SessionID
	u := req.User
	sl, err := get_sessionlog_by_sessionid(ctx, sid)
	if err != nil {
		return nil, err
	}
	if sl == nil {
		return nil, errors.InvalidArgs(ctx, "invalid sessiontoken", "invalid sessiontoken")
	}
	if sl.UserID != "" && u.ID != "" && sl.UserID != u.ID {
		return nil, errors.InvalidArgs(ctx, "auth error", "session \"%s\" changed user from \"%s\" to \"%s\"", sid, sl.UserID, u.ID)
	}

	sl.UserID = u.ID
	sl.Useremail = u.Email
	sl.Username = fmt.Sprintf("%s %s", u.FirstName, u.LastName)
	err = db.DefaultDBSessionLog().Update(ctx, sl)
	if err != nil {
		return nil, err
	}

	res.IsValid = true
	fmt.Printf("Associated user %s with session %s\n", sl.UserID, req.Session.SessionID)
	return res, nil
}

func get_sessionlog_by_sessionid(ctx context.Context, sessionid string) (*pb.SessionLog, error) {
	sls, err := db.DefaultDBSessionLog().BySessionToken(ctx, sessionid)
	if err != nil {
		return nil, err
	}
	if len(sls) == 0 {
		return nil, nil
	}
	sl := sls[0]
	return sl, nil

}
func create_session_from_log(ctx context.Context, req *pb.SessionLog) (*session.Session, error) {
	res := &session.Session{
		SessionID: req.SessionToken,
	}
	return res, nil
}





