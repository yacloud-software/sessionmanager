// Code generated by protoc-gen-go.
// source: protos/golang.yacloud.eu/apis/sessionmanager/sessionmanager.proto
// DO NOT EDIT!

/*
Package sessionmanager is a generated protocol buffer package.

It is generated from these files:
	protos/golang.yacloud.eu/apis/sessionmanager/sessionmanager.proto

It has these top-level messages:
	PingResponse
	SessionLog
	NewSessionRequest
	SessionResponse
	SessionVerifyResponse
	SessionToken
	SessionAliveResponse
	User2SessionRequest
*/
package sessionmanager

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import common "golang.conradwood.net/apis/common"
import session "golang.yacloud.eu/apis/session"
import auth "golang.conradwood.net/apis/auth"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// comment: message pingresponse
type PingResponse struct {
	// comment: field pingresponse.response
	Response string `protobuf:"bytes,1,opt,name=Response" json:"Response,omitempty"`
}

func (m *PingResponse) Reset()                    { *m = PingResponse{} }
func (m *PingResponse) String() string            { return proto.CompactTextString(m) }
func (*PingResponse) ProtoMessage()               {}
func (*PingResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *PingResponse) GetResponse() string {
	if m != nil {
		return m.Response
	}
	return ""
}

type SessionLog struct {
	ID           uint64 `protobuf:"varint,1,opt,name=ID" json:"ID,omitempty"`
	UserID       string `protobuf:"bytes,2,opt,name=UserID" json:"UserID,omitempty"`
	Username     string `protobuf:"bytes,3,opt,name=Username" json:"Username,omitempty"`
	Useremail    string `protobuf:"bytes,4,opt,name=Useremail" json:"Useremail,omitempty"`
	IP           string `protobuf:"bytes,5,opt,name=IP" json:"IP,omitempty"`
	UserAgent    string `protobuf:"bytes,6,opt,name=UserAgent" json:"UserAgent,omitempty"`
	Created      uint32 `protobuf:"varint,7,opt,name=Created" json:"Created,omitempty"`
	BrowserID    string `protobuf:"bytes,8,opt,name=BrowserID" json:"BrowserID,omitempty"`
	SessionToken string `protobuf:"bytes,9,opt,name=SessionToken" json:"SessionToken,omitempty"`
	LastUsed     uint32 `protobuf:"varint,10,opt,name=LastUsed" json:"LastUsed,omitempty"`
	TriggerHost  string `protobuf:"bytes,11,opt,name=TriggerHost" json:"TriggerHost,omitempty"`
}

func (m *SessionLog) Reset()                    { *m = SessionLog{} }
func (m *SessionLog) String() string            { return proto.CompactTextString(m) }
func (*SessionLog) ProtoMessage()               {}
func (*SessionLog) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *SessionLog) GetID() uint64 {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *SessionLog) GetUserID() string {
	if m != nil {
		return m.UserID
	}
	return ""
}

func (m *SessionLog) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *SessionLog) GetUseremail() string {
	if m != nil {
		return m.Useremail
	}
	return ""
}

func (m *SessionLog) GetIP() string {
	if m != nil {
		return m.IP
	}
	return ""
}

func (m *SessionLog) GetUserAgent() string {
	if m != nil {
		return m.UserAgent
	}
	return ""
}

func (m *SessionLog) GetCreated() uint32 {
	if m != nil {
		return m.Created
	}
	return 0
}

func (m *SessionLog) GetBrowserID() string {
	if m != nil {
		return m.BrowserID
	}
	return ""
}

func (m *SessionLog) GetSessionToken() string {
	if m != nil {
		return m.SessionToken
	}
	return ""
}

func (m *SessionLog) GetLastUsed() uint32 {
	if m != nil {
		return m.LastUsed
	}
	return 0
}

func (m *SessionLog) GetTriggerHost() string {
	if m != nil {
		return m.TriggerHost
	}
	return ""
}

type NewSessionRequest struct {
	IPAddress   string `protobuf:"bytes,1,opt,name=IPAddress" json:"IPAddress,omitempty"`
	UserAgent   string `protobuf:"bytes,2,opt,name=UserAgent" json:"UserAgent,omitempty"`
	BrowserID   string `protobuf:"bytes,3,opt,name=BrowserID" json:"BrowserID,omitempty"`
	UserID      string `protobuf:"bytes,4,opt,name=UserID" json:"UserID,omitempty"`
	Username    string `protobuf:"bytes,5,opt,name=Username" json:"Username,omitempty"`
	Useremail   string `protobuf:"bytes,6,opt,name=Useremail" json:"Useremail,omitempty"`
	TriggerHost string `protobuf:"bytes,7,opt,name=TriggerHost" json:"TriggerHost,omitempty"`
}

func (m *NewSessionRequest) Reset()                    { *m = NewSessionRequest{} }
func (m *NewSessionRequest) String() string            { return proto.CompactTextString(m) }
func (*NewSessionRequest) ProtoMessage()               {}
func (*NewSessionRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *NewSessionRequest) GetIPAddress() string {
	if m != nil {
		return m.IPAddress
	}
	return ""
}

func (m *NewSessionRequest) GetUserAgent() string {
	if m != nil {
		return m.UserAgent
	}
	return ""
}

func (m *NewSessionRequest) GetBrowserID() string {
	if m != nil {
		return m.BrowserID
	}
	return ""
}

func (m *NewSessionRequest) GetUserID() string {
	if m != nil {
		return m.UserID
	}
	return ""
}

func (m *NewSessionRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *NewSessionRequest) GetUseremail() string {
	if m != nil {
		return m.Useremail
	}
	return ""
}

func (m *NewSessionRequest) GetTriggerHost() string {
	if m != nil {
		return m.TriggerHost
	}
	return ""
}

type SessionResponse struct {
	Token                string           `protobuf:"bytes,1,opt,name=Token" json:"Token,omitempty"`
	LastSessionTimestamp uint32           `protobuf:"varint,2,opt,name=LastSessionTimestamp" json:"LastSessionTimestamp,omitempty"`
	NewDevice            bool             `protobuf:"varint,3,opt,name=NewDevice" json:"NewDevice,omitempty"`
	Session              *session.Session `protobuf:"bytes,4,opt,name=Session" json:"Session,omitempty"`
}

func (m *SessionResponse) Reset()                    { *m = SessionResponse{} }
func (m *SessionResponse) String() string            { return proto.CompactTextString(m) }
func (*SessionResponse) ProtoMessage()               {}
func (*SessionResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *SessionResponse) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *SessionResponse) GetLastSessionTimestamp() uint32 {
	if m != nil {
		return m.LastSessionTimestamp
	}
	return 0
}

func (m *SessionResponse) GetNewDevice() bool {
	if m != nil {
		return m.NewDevice
	}
	return false
}

func (m *SessionResponse) GetSession() *session.Session {
	if m != nil {
		return m.Session
	}
	return nil
}

type SessionVerifyResponse struct {
	IsValid    bool        `protobuf:"varint,1,opt,name=IsValid" json:"IsValid,omitempty"`
	SessionLog *SessionLog `protobuf:"bytes,2,opt,name=SessionLog" json:"SessionLog,omitempty"`
}

func (m *SessionVerifyResponse) Reset()                    { *m = SessionVerifyResponse{} }
func (m *SessionVerifyResponse) String() string            { return proto.CompactTextString(m) }
func (*SessionVerifyResponse) ProtoMessage()               {}
func (*SessionVerifyResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *SessionVerifyResponse) GetIsValid() bool {
	if m != nil {
		return m.IsValid
	}
	return false
}

func (m *SessionVerifyResponse) GetSessionLog() *SessionLog {
	if m != nil {
		return m.SessionLog
	}
	return nil
}

type SessionToken struct {
	Token string `protobuf:"bytes,1,opt,name=Token" json:"Token,omitempty"`
}

func (m *SessionToken) Reset()                    { *m = SessionToken{} }
func (m *SessionToken) String() string            { return proto.CompactTextString(m) }
func (*SessionToken) ProtoMessage()               {}
func (*SessionToken) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *SessionToken) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type SessionAliveResponse struct {
	IsValid bool             `protobuf:"varint,1,opt,name=IsValid" json:"IsValid,omitempty"`
	Session *session.Session `protobuf:"bytes,2,opt,name=Session" json:"Session,omitempty"`
}

func (m *SessionAliveResponse) Reset()                    { *m = SessionAliveResponse{} }
func (m *SessionAliveResponse) String() string            { return proto.CompactTextString(m) }
func (*SessionAliveResponse) ProtoMessage()               {}
func (*SessionAliveResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *SessionAliveResponse) GetIsValid() bool {
	if m != nil {
		return m.IsValid
	}
	return false
}

func (m *SessionAliveResponse) GetSession() *session.Session {
	if m != nil {
		return m.Session
	}
	return nil
}

type User2SessionRequest struct {
	Session *session.Session `protobuf:"bytes,1,opt,name=Session" json:"Session,omitempty"`
	User    *auth.User       `protobuf:"bytes,2,opt,name=User" json:"User,omitempty"`
}

func (m *User2SessionRequest) Reset()                    { *m = User2SessionRequest{} }
func (m *User2SessionRequest) String() string            { return proto.CompactTextString(m) }
func (*User2SessionRequest) ProtoMessage()               {}
func (*User2SessionRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *User2SessionRequest) GetSession() *session.Session {
	if m != nil {
		return m.Session
	}
	return nil
}

func (m *User2SessionRequest) GetUser() *auth.User {
	if m != nil {
		return m.User
	}
	return nil
}

func init() {
	proto.RegisterType((*PingResponse)(nil), "sessionmanager.PingResponse")
	proto.RegisterType((*SessionLog)(nil), "sessionmanager.SessionLog")
	proto.RegisterType((*NewSessionRequest)(nil), "sessionmanager.NewSessionRequest")
	proto.RegisterType((*SessionResponse)(nil), "sessionmanager.SessionResponse")
	proto.RegisterType((*SessionVerifyResponse)(nil), "sessionmanager.SessionVerifyResponse")
	proto.RegisterType((*SessionToken)(nil), "sessionmanager.SessionToken")
	proto.RegisterType((*SessionAliveResponse)(nil), "sessionmanager.SessionAliveResponse")
	proto.RegisterType((*User2SessionRequest)(nil), "sessionmanager.User2SessionRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for SessionManager service

type SessionManagerClient interface {
	// comment: rpc ping
	Ping(ctx context.Context, in *common.Void, opts ...grpc.CallOption) (*PingResponse, error)
	// create a new session
	NewSession(ctx context.Context, in *NewSessionRequest, opts ...grpc.CallOption) (*SessionResponse, error)
	// verify a session (by token)
	VerifySession(ctx context.Context, in *SessionToken, opts ...grpc.CallOption) (*SessionVerifyResponse, error)
	// keep a session alive
	KeepAliveSession(ctx context.Context, in *SessionToken, opts ...grpc.CallOption) (*SessionAliveResponse, error)
	// update a session (e.g. to set a user)
	User2Session(ctx context.Context, in *User2SessionRequest, opts ...grpc.CallOption) (*SessionAliveResponse, error)
}

type sessionManagerClient struct {
	cc *grpc.ClientConn
}

func NewSessionManagerClient(cc *grpc.ClientConn) SessionManagerClient {
	return &sessionManagerClient{cc}
}

func (c *sessionManagerClient) Ping(ctx context.Context, in *common.Void, opts ...grpc.CallOption) (*PingResponse, error) {
	out := new(PingResponse)
	err := grpc.Invoke(ctx, "/sessionmanager.SessionManager/Ping", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sessionManagerClient) NewSession(ctx context.Context, in *NewSessionRequest, opts ...grpc.CallOption) (*SessionResponse, error) {
	out := new(SessionResponse)
	err := grpc.Invoke(ctx, "/sessionmanager.SessionManager/NewSession", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sessionManagerClient) VerifySession(ctx context.Context, in *SessionToken, opts ...grpc.CallOption) (*SessionVerifyResponse, error) {
	out := new(SessionVerifyResponse)
	err := grpc.Invoke(ctx, "/sessionmanager.SessionManager/VerifySession", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sessionManagerClient) KeepAliveSession(ctx context.Context, in *SessionToken, opts ...grpc.CallOption) (*SessionAliveResponse, error) {
	out := new(SessionAliveResponse)
	err := grpc.Invoke(ctx, "/sessionmanager.SessionManager/KeepAliveSession", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sessionManagerClient) User2Session(ctx context.Context, in *User2SessionRequest, opts ...grpc.CallOption) (*SessionAliveResponse, error) {
	out := new(SessionAliveResponse)
	err := grpc.Invoke(ctx, "/sessionmanager.SessionManager/User2Session", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for SessionManager service

type SessionManagerServer interface {
	// comment: rpc ping
	Ping(context.Context, *common.Void) (*PingResponse, error)
	// create a new session
	NewSession(context.Context, *NewSessionRequest) (*SessionResponse, error)
	// verify a session (by token)
	VerifySession(context.Context, *SessionToken) (*SessionVerifyResponse, error)
	// keep a session alive
	KeepAliveSession(context.Context, *SessionToken) (*SessionAliveResponse, error)
	// update a session (e.g. to set a user)
	User2Session(context.Context, *User2SessionRequest) (*SessionAliveResponse, error)
}

func RegisterSessionManagerServer(s *grpc.Server, srv SessionManagerServer) {
	s.RegisterService(&_SessionManager_serviceDesc, srv)
}

func _SessionManager_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.Void)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SessionManagerServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sessionmanager.SessionManager/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SessionManagerServer).Ping(ctx, req.(*common.Void))
	}
	return interceptor(ctx, in, info, handler)
}

func _SessionManager_NewSession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewSessionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SessionManagerServer).NewSession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sessionmanager.SessionManager/NewSession",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SessionManagerServer).NewSession(ctx, req.(*NewSessionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SessionManager_VerifySession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SessionToken)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SessionManagerServer).VerifySession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sessionmanager.SessionManager/VerifySession",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SessionManagerServer).VerifySession(ctx, req.(*SessionToken))
	}
	return interceptor(ctx, in, info, handler)
}

func _SessionManager_KeepAliveSession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SessionToken)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SessionManagerServer).KeepAliveSession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sessionmanager.SessionManager/KeepAliveSession",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SessionManagerServer).KeepAliveSession(ctx, req.(*SessionToken))
	}
	return interceptor(ctx, in, info, handler)
}

func _SessionManager_User2Session_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User2SessionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SessionManagerServer).User2Session(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sessionmanager.SessionManager/User2Session",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SessionManagerServer).User2Session(ctx, req.(*User2SessionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _SessionManager_serviceDesc = grpc.ServiceDesc{
	ServiceName: "sessionmanager.SessionManager",
	HandlerType: (*SessionManagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _SessionManager_Ping_Handler,
		},
		{
			MethodName: "NewSession",
			Handler:    _SessionManager_NewSession_Handler,
		},
		{
			MethodName: "VerifySession",
			Handler:    _SessionManager_VerifySession_Handler,
		},
		{
			MethodName: "KeepAliveSession",
			Handler:    _SessionManager_KeepAliveSession_Handler,
		},
		{
			MethodName: "User2Session",
			Handler:    _SessionManager_User2Session_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/golang.yacloud.eu/apis/sessionmanager/sessionmanager.proto",
}

func init() {
	proto.RegisterFile("protos/golang.yacloud.eu/apis/sessionmanager/sessionmanager.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 669 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x55, 0xcf, 0x6e, 0xd3, 0x4e,
	0x10, 0x56, 0xdc, 0x34, 0x49, 0x27, 0x69, 0x7f, 0xfd, 0x2d, 0x05, 0x19, 0xab, 0x2a, 0x21, 0x14,
	0xa9, 0xaa, 0x2a, 0x57, 0x0a, 0x37, 0x6e, 0x29, 0x39, 0x10, 0x51, 0xaa, 0xc8, 0xb4, 0x91, 0x90,
	0xb8, 0x2c, 0xf1, 0x60, 0x2c, 0x62, 0x6f, 0xf0, 0x3a, 0x8d, 0x7a, 0xe5, 0x45, 0x38, 0x71, 0xe1,
	0x29, 0x78, 0x15, 0xde, 0x04, 0xed, 0x1f, 0xff, 0xdb, 0x36, 0x6d, 0x2e, 0xc9, 0xce, 0xec, 0xcc,
	0x37, 0xf3, 0xcd, 0xb7, 0xbb, 0x86, 0xc1, 0x3c, 0x61, 0x29, 0xe3, 0xa7, 0x01, 0x9b, 0xd1, 0x38,
	0x70, 0x6f, 0xe8, 0x74, 0xc6, 0x16, 0xbe, 0x8b, 0x8b, 0x53, 0x3a, 0x0f, 0xf9, 0x29, 0x47, 0xce,
	0x43, 0x16, 0x47, 0x34, 0xa6, 0x01, 0x26, 0x86, 0xe9, 0xca, 0x5c, 0xb2, 0x53, 0xf5, 0x3a, 0xae,
	0xc6, 0x9a, 0xb2, 0x38, 0xa1, 0xfe, 0x92, 0x31, 0xdf, 0x8d, 0x31, 0x55, 0x78, 0x53, 0x16, 0x45,
	0x2c, 0xd6, 0x7f, 0x2a, 0xdf, 0x39, 0xb9, 0xbf, 0x76, 0xf6, 0xaf, 0xa3, 0x8f, 0xef, 0x41, 0xa7,
	0x8b, 0xf4, 0xab, 0xfc, 0x51, 0xb1, 0xbd, 0x63, 0xe8, 0x8c, 0xc3, 0x38, 0xf0, 0x90, 0xcf, 0x59,
	0xcc, 0x91, 0x38, 0xd0, 0xca, 0xd6, 0x76, 0xad, 0x5b, 0x3b, 0xda, 0xf2, 0x72, 0xbb, 0xf7, 0xc7,
	0x02, 0xf8, 0xa0, 0x2a, 0x9d, 0xb3, 0x80, 0xec, 0x80, 0x35, 0x1a, 0xca, 0xa0, 0xba, 0x67, 0x8d,
	0x86, 0xe4, 0x09, 0x34, 0xae, 0x38, 0x26, 0xa3, 0xa1, 0x6d, 0xc9, 0x44, 0x6d, 0x09, 0x48, 0xb1,
	0x8a, 0x69, 0x84, 0xf6, 0x86, 0x82, 0xcc, 0x6c, 0xb2, 0x0f, 0x5b, 0x62, 0x8d, 0x11, 0x0d, 0x67,
	0x76, 0x5d, 0x6e, 0x16, 0x0e, 0x59, 0x61, 0x6c, 0x6f, 0x4a, 0xb7, 0x35, 0x1a, 0x67, 0xd1, 0x83,
	0x00, 0xe3, 0xd4, 0x6e, 0x14, 0xd1, 0xd2, 0x41, 0x6c, 0x68, 0xbe, 0x49, 0x90, 0xa6, 0xe8, 0xdb,
	0xcd, 0x6e, 0xed, 0x68, 0xdb, 0xcb, 0x4c, 0x91, 0x77, 0x96, 0xb0, 0xa5, 0x6a, 0xae, 0xa5, 0xf2,
	0x72, 0x07, 0x39, 0x81, 0x8e, 0x66, 0x75, 0xc9, 0xbe, 0x61, 0x6c, 0x6f, 0x89, 0x80, 0xb3, 0xd6,
	0xef, 0x1f, 0x4f, 0xeb, 0x69, 0xb2, 0x40, 0xaf, 0xb2, 0x2b, 0xd8, 0x9c, 0x53, 0x9e, 0x5e, 0x71,
	0xf4, 0x6d, 0x90, 0x65, 0x72, 0x9b, 0x74, 0xa1, 0x7d, 0x99, 0x84, 0x41, 0x80, 0xc9, 0x5b, 0xc6,
	0x53, 0xbb, 0x2d, 0x2b, 0x95, 0x5d, 0xbd, 0xbf, 0x35, 0xf8, 0xff, 0x02, 0x97, 0x1a, 0xd1, 0xc3,
	0xef, 0x0b, 0xe4, 0xa9, 0xe8, 0x6f, 0x34, 0x1e, 0xf8, 0x7e, 0x82, 0x9c, 0xeb, 0xa9, 0x17, 0x8e,
	0x2a, 0x6b, 0xcb, 0x64, 0x5d, 0xe1, 0xb6, 0x61, 0x72, 0x2b, 0x34, 0xa9, 0xaf, 0xd4, 0x64, 0xf3,
	0x3e, 0x4d, 0x1a, 0xa6, 0x26, 0x06, 0xc7, 0xe6, 0x6d, 0x8e, 0xbf, 0x6a, 0xf0, 0x5f, 0x4e, 0x50,
	0x1f, 0xab, 0x3d, 0xd8, 0x54, 0xc3, 0x55, 0xec, 0x94, 0x41, 0xfa, 0xb0, 0x27, 0x66, 0x97, 0xcd,
	0x37, 0x8c, 0x90, 0xa7, 0x34, 0x9a, 0x4b, 0x92, 0xdb, 0xde, 0x9d, 0x7b, 0xa2, 0xbb, 0x0b, 0x5c,
	0x0e, 0xf1, 0x3a, 0x9c, 0xaa, 0xe3, 0xd4, 0xf2, 0x0a, 0x07, 0x39, 0x86, 0xa6, 0xce, 0x90, 0x84,
	0xdb, 0xfd, 0x5d, 0x37, 0xbb, 0x1b, 0x59, 0x4b, 0x59, 0x40, 0x2f, 0x82, 0xc7, 0x7a, 0x39, 0xc1,
	0x24, 0xfc, 0x72, 0x93, 0x37, 0x6b, 0x43, 0x73, 0xc4, 0x27, 0x74, 0x16, 0xfa, 0xb2, 0xdd, 0x96,
	0x97, 0x99, 0xe4, 0x75, 0xf9, 0x02, 0xc8, 0x36, 0xdb, 0x7d, 0xc7, 0x35, 0xae, 0x7c, 0x11, 0xe1,
	0x95, 0xa2, 0x7b, 0x87, 0xd5, 0x63, 0x76, 0xf7, 0x48, 0x7a, 0x9f, 0x60, 0x4f, 0x47, 0x0d, 0x66,
	0xe1, 0x35, 0xae, 0xd1, 0x53, 0x89, 0xb2, 0xf5, 0x10, 0x65, 0x0a, 0x8f, 0x84, 0x92, 0x7d, 0xe3,
	0xfc, 0x95, 0x20, 0x6a, 0x0f, 0x40, 0x90, 0x03, 0xa8, 0x0b, 0x08, 0x5d, 0x0b, 0x5c, 0xf9, 0x96,
	0x08, 0x8f, 0x27, 0xfd, 0xfd, 0x9f, 0x1b, 0xb0, 0xa3, 0x63, 0xdf, 0xab, 0x81, 0x90, 0x3e, 0xd4,
	0xc5, 0x1b, 0x43, 0x3a, 0xae, 0x7e, 0xd4, 0x26, 0x2c, 0xf4, 0x9d, 0x7d, 0x73, 0x6e, 0x95, 0x77,
	0x68, 0x0c, 0x50, 0xdc, 0x13, 0xf2, 0xdc, 0x8c, 0xbd, 0x75, 0x87, 0x9c, 0x67, 0x2b, 0x64, 0xc8,
	0x11, 0x2f, 0x61, 0x5b, 0xe9, 0x9c, 0x81, 0xee, 0xaf, 0xc8, 0x90, 0x42, 0x38, 0x2f, 0x57, 0xec,
	0x1a, 0x67, 0x65, 0x02, 0xbb, 0xef, 0x10, 0xe7, 0x52, 0xac, 0xf5, 0x80, 0x0f, 0x57, 0xec, 0x56,
	0xf5, 0xfe, 0x08, 0x9d, 0xb2, 0x52, 0xe4, 0x85, 0x99, 0x75, 0x87, 0x8e, 0xeb, 0x41, 0x9f, 0x75,
	0xe1, 0x00, 0x17, 0xf9, 0xa7, 0x44, 0x7c, 0x15, 0x8c, 0xb4, 0xcf, 0x0d, 0xf9, 0x6d, 0x78, 0xf5,
	0x2f, 0x00, 0x00, 0xff, 0xff, 0xa9, 0xdf, 0xf6, 0xa1, 0xfa, 0x06, 0x00, 0x00,
}
