package db

/*
 This file was created by mkdb-client.
 The intention is not to modify thils file, but you may extend the struct DBSessionLog
 in a seperate file (so that you can regenerate this one from time to time)
*/

/*
 PRIMARY KEY: ID
*/

/*
 postgres:
 create sequence sessionlog_seq;

Main Table:

 CREATE TABLE sessionlog (id integer primary key default nextval('sessionlog_seq'),userid text not null  ,username text not null  ,useremail text not null  ,ip text not null  ,useragent text not null  ,created integer not null  ,browserid text not null  ,sessiontoken text not null  unique  ,lastused integer not null  ,triggerhost text not null  );

Alter statements:
ALTER TABLE sessionlog ADD COLUMN IF NOT EXISTS userid text not null default '';
ALTER TABLE sessionlog ADD COLUMN IF NOT EXISTS username text not null default '';
ALTER TABLE sessionlog ADD COLUMN IF NOT EXISTS useremail text not null default '';
ALTER TABLE sessionlog ADD COLUMN IF NOT EXISTS ip text not null default '';
ALTER TABLE sessionlog ADD COLUMN IF NOT EXISTS useragent text not null default '';
ALTER TABLE sessionlog ADD COLUMN IF NOT EXISTS created integer not null default 0;
ALTER TABLE sessionlog ADD COLUMN IF NOT EXISTS browserid text not null default '';
ALTER TABLE sessionlog ADD COLUMN IF NOT EXISTS sessiontoken text not null unique  default '';
ALTER TABLE sessionlog ADD COLUMN IF NOT EXISTS lastused integer not null default 0;
ALTER TABLE sessionlog ADD COLUMN IF NOT EXISTS triggerhost text not null default '';


Archive Table: (structs can be moved from main to archive using Archive() function)

 CREATE TABLE sessionlog_archive (id integer unique not null,userid text not null,username text not null,useremail text not null,ip text not null,useragent text not null,created integer not null,browserid text not null,sessiontoken text not null,lastused integer not null,triggerhost text not null);
*/

import (
	"context"
	gosql "database/sql"
	"fmt"
	"golang.conradwood.net/go-easyops/sql"
	savepb "golang.yacloud.eu/apis/sessionmanager"
	"os"
)

var (
	default_def_DBSessionLog *DBSessionLog
)

type DBSessionLog struct {
	DB                  *sql.DB
	SQLTablename        string
	SQLArchivetablename string
}

func DefaultDBSessionLog() *DBSessionLog {
	if default_def_DBSessionLog != nil {
		return default_def_DBSessionLog
	}
	psql, err := sql.Open()
	if err != nil {
		fmt.Printf("Failed to open database: %s\n", err)
		os.Exit(10)
	}
	res := NewDBSessionLog(psql)
	ctx := context.Background()
	err = res.CreateTable(ctx)
	if err != nil {
		fmt.Printf("Failed to create table: %s\n", err)
		os.Exit(10)
	}
	default_def_DBSessionLog = res
	return res
}
func NewDBSessionLog(db *sql.DB) *DBSessionLog {
	foo := DBSessionLog{DB: db}
	foo.SQLTablename = "sessionlog"
	foo.SQLArchivetablename = "sessionlog_archive"
	return &foo
}

// archive. It is NOT transactionally save.
func (a *DBSessionLog) Archive(ctx context.Context, id uint64) error {

	// load it
	p, err := a.ByID(ctx, id)
	if err != nil {
		return err
	}

	// now save it to archive:
	_, e := a.DB.ExecContext(ctx, "archive_DBSessionLog", "insert into "+a.SQLArchivetablename+" (id,userid, username, useremail, ip, useragent, created, browserid, sessiontoken, lastused, triggerhost) values ($1,$2, $3, $4, $5, $6, $7, $8, $9, $10, $11) ", p.ID, p.UserID, p.Username, p.Useremail, p.IP, p.UserAgent, p.Created, p.BrowserID, p.SessionToken, p.LastUsed, p.TriggerHost)
	if e != nil {
		return e
	}

	// now delete it.
	a.DeleteByID(ctx, id)
	return nil
}

// Save (and use database default ID generation)
func (a *DBSessionLog) Save(ctx context.Context, p *savepb.SessionLog) (uint64, error) {
	qn := "DBSessionLog_Save"
	rows, e := a.DB.QueryContext(ctx, qn, "insert into "+a.SQLTablename+" (userid, username, useremail, ip, useragent, created, browserid, sessiontoken, lastused, triggerhost) values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) returning id", p.UserID, p.Username, p.Useremail, p.IP, p.UserAgent, p.Created, p.BrowserID, p.SessionToken, p.LastUsed, p.TriggerHost)
	if e != nil {
		return 0, a.Error(ctx, qn, e)
	}
	defer rows.Close()
	if !rows.Next() {
		return 0, a.Error(ctx, qn, fmt.Errorf("No rows after insert"))
	}
	var id uint64
	e = rows.Scan(&id)
	if e != nil {
		return 0, a.Error(ctx, qn, fmt.Errorf("failed to scan id after insert: %s", e))
	}
	p.ID = id
	return id, nil
}

// Save using the ID specified
func (a *DBSessionLog) SaveWithID(ctx context.Context, p *savepb.SessionLog) error {
	qn := "insert_DBSessionLog"
	_, e := a.DB.ExecContext(ctx, qn, "insert into "+a.SQLTablename+" (id,userid, username, useremail, ip, useragent, created, browserid, sessiontoken, lastused, triggerhost) values ($1,$2, $3, $4, $5, $6, $7, $8, $9, $10, $11) ", p.ID, p.UserID, p.Username, p.Useremail, p.IP, p.UserAgent, p.Created, p.BrowserID, p.SessionToken, p.LastUsed, p.TriggerHost)
	return a.Error(ctx, qn, e)
}

func (a *DBSessionLog) Update(ctx context.Context, p *savepb.SessionLog) error {
	qn := "DBSessionLog_Update"
	_, e := a.DB.ExecContext(ctx, qn, "update "+a.SQLTablename+" set userid=$1, username=$2, useremail=$3, ip=$4, useragent=$5, created=$6, browserid=$7, sessiontoken=$8, lastused=$9, triggerhost=$10 where id = $11", p.UserID, p.Username, p.Useremail, p.IP, p.UserAgent, p.Created, p.BrowserID, p.SessionToken, p.LastUsed, p.TriggerHost, p.ID)

	return a.Error(ctx, qn, e)
}

// delete by id field
func (a *DBSessionLog) DeleteByID(ctx context.Context, p uint64) error {
	qn := "deleteDBSessionLog_ByID"
	_, e := a.DB.ExecContext(ctx, qn, "delete from "+a.SQLTablename+" where id = $1", p)
	return a.Error(ctx, qn, e)
}

// get it by primary id
func (a *DBSessionLog) ByID(ctx context.Context, p uint64) (*savepb.SessionLog, error) {
	qn := "DBSessionLog_ByID"
	rows, e := a.DB.QueryContext(ctx, qn, "select id,userid, username, useremail, ip, useragent, created, browserid, sessiontoken, lastused, triggerhost from "+a.SQLTablename+" where id = $1", p)
	if e != nil {
		return nil, a.Error(ctx, qn, fmt.Errorf("ByID: error querying (%s)", e))
	}
	defer rows.Close()
	l, e := a.FromRows(ctx, rows)
	if e != nil {
		return nil, a.Error(ctx, qn, fmt.Errorf("ByID: error scanning (%s)", e))
	}
	if len(l) == 0 {
		return nil, a.Error(ctx, qn, fmt.Errorf("No SessionLog with id %v", p))
	}
	if len(l) != 1 {
		return nil, a.Error(ctx, qn, fmt.Errorf("Multiple (%d) SessionLog with id %v", len(l), p))
	}
	return l[0], nil
}

// get it by primary id (nil if no such ID row, but no error either)
func (a *DBSessionLog) TryByID(ctx context.Context, p uint64) (*savepb.SessionLog, error) {
	qn := "DBSessionLog_TryByID"
	rows, e := a.DB.QueryContext(ctx, qn, "select id,userid, username, useremail, ip, useragent, created, browserid, sessiontoken, lastused, triggerhost from "+a.SQLTablename+" where id = $1", p)
	if e != nil {
		return nil, a.Error(ctx, qn, fmt.Errorf("TryByID: error querying (%s)", e))
	}
	defer rows.Close()
	l, e := a.FromRows(ctx, rows)
	if e != nil {
		return nil, a.Error(ctx, qn, fmt.Errorf("TryByID: error scanning (%s)", e))
	}
	if len(l) == 0 {
		return nil, nil
	}
	if len(l) != 1 {
		return nil, a.Error(ctx, qn, fmt.Errorf("Multiple (%d) SessionLog with id %v", len(l), p))
	}
	return l[0], nil
}

// get all rows
func (a *DBSessionLog) All(ctx context.Context) ([]*savepb.SessionLog, error) {
	qn := "DBSessionLog_all"
	rows, e := a.DB.QueryContext(ctx, qn, "select id,userid, username, useremail, ip, useragent, created, browserid, sessiontoken, lastused, triggerhost from "+a.SQLTablename+" order by id")
	if e != nil {
		return nil, a.Error(ctx, qn, fmt.Errorf("All: error querying (%s)", e))
	}
	defer rows.Close()
	l, e := a.FromRows(ctx, rows)
	if e != nil {
		return nil, fmt.Errorf("All: error scanning (%s)", e)
	}
	return l, nil
}

/**********************************************************************
* GetBy[FIELD] functions
**********************************************************************/

// get all "DBSessionLog" rows with matching UserID
func (a *DBSessionLog) ByUserID(ctx context.Context, p string) ([]*savepb.SessionLog, error) {
	qn := "DBSessionLog_ByUserID"
	rows, e := a.DB.QueryContext(ctx, qn, "select id,userid, username, useremail, ip, useragent, created, browserid, sessiontoken, lastused, triggerhost from "+a.SQLTablename+" where userid = $1", p)
	if e != nil {
		return nil, a.Error(ctx, qn, fmt.Errorf("ByUserID: error querying (%s)", e))
	}
	defer rows.Close()
	l, e := a.FromRows(ctx, rows)
	if e != nil {
		return nil, a.Error(ctx, qn, fmt.Errorf("ByUserID: error scanning (%s)", e))
	}
	return l, nil
}

// the 'like' lookup
func (a *DBSessionLog) ByLikeUserID(ctx context.Context, p string) ([]*savepb.SessionLog, error) {
	qn := "DBSessionLog_ByLikeUserID"
	rows, e := a.DB.QueryContext(ctx, qn, "select id,userid, username, useremail, ip, useragent, created, browserid, sessiontoken, lastused, triggerhost from "+a.SQLTablename+" where userid ilike $1", p)
	if e != nil {
		return nil, a.Error(ctx, qn, fmt.Errorf("ByUserID: error querying (%s)", e))
	}
	defer rows.Close()
	l, e := a.FromRows(ctx, rows)
	if e != nil {
		return nil, a.Error(ctx, qn, fmt.Errorf("ByUserID: error scanning (%s)", e))
	}
	return l, nil
}

// get all "DBSessionLog" rows with matching Username
func (a *DBSessionLog) ByUsername(ctx context.Context, p string) ([]*savepb.SessionLog, error) {
	qn := "DBSessionLog_ByUsername"
	rows, e := a.DB.QueryContext(ctx, qn, "select id,userid, username, useremail, ip, useragent, created, browserid, sessiontoken, lastused, triggerhost from "+a.SQLTablename+" where username = $1", p)
	if e != nil {
		return nil, a.Error(ctx, qn, fmt.Errorf("ByUsername: error querying (%s)", e))
	}
	defer rows.Close()
	l, e := a.FromRows(ctx, rows)
	if e != nil {
		return nil, a.Error(ctx, qn, fmt.Errorf("ByUsername: error scanning (%s)", e))
	}
	return l, nil
}

// the 'like' lookup
func (a *DBSessionLog) ByLikeUsername(ctx context.Context, p string) ([]*savepb.SessionLog, error) {
	qn := "DBSessionLog_ByLikeUsername"
	rows, e := a.DB.QueryContext(ctx, qn, "select id,userid, username, useremail, ip, useragent, created, browserid, sessiontoken, lastused, triggerhost from "+a.SQLTablename+" where username ilike $1", p)
	if e != nil {
		return nil, a.Error(ctx, qn, fmt.Errorf("ByUsername: error querying (%s)", e))
	}
	defer rows.Close()
	l, e := a.FromRows(ctx, rows)
	if e != nil {
		return nil, a.Error(ctx, qn, fmt.Errorf("ByUsername: error scanning (%s)", e))
	}
	return l, nil
}

// get all "DBSessionLog" rows with matching Useremail
func (a *DBSessionLog) ByUseremail(ctx context.Context, p string) ([]*savepb.SessionLog, error) {
	qn := "DBSessionLog_ByUseremail"
	rows, e := a.DB.QueryContext(ctx, qn, "select id,userid, username, useremail, ip, useragent, created, browserid, sessiontoken, lastused, triggerhost from "+a.SQLTablename+" where useremail = $1", p)
	if e != nil {
		return nil, a.Error(ctx, qn, fmt.Errorf("ByUseremail: error querying (%s)", e))
	}
	defer rows.Close()
	l, e := a.FromRows(ctx, rows)
	if e != nil {
		return nil, a.Error(ctx, qn, fmt.Errorf("ByUseremail: error scanning (%s)", e))
	}
	return l, nil
}

// the 'like' lookup
func (a *DBSessionLog) ByLikeUseremail(ctx context.Context, p string) ([]*savepb.SessionLog, error) {
	qn := "DBSessionLog_ByLikeUseremail"
	rows, e := a.DB.QueryContext(ctx, qn, "select id,userid, username, useremail, ip, useragent, created, browserid, sessiontoken, lastused, triggerhost from "+a.SQLTablename+" where useremail ilike $1", p)
	if e != nil {
		return nil, a.Error(ctx, qn, fmt.Errorf("ByUseremail: error querying (%s)", e))
	}
	defer rows.Close()
	l, e := a.FromRows(ctx, rows)
	if e != nil {
		return nil, a.Error(ctx, qn, fmt.Errorf("ByUseremail: error scanning (%s)", e))
	}
	return l, nil
}

// get all "DBSessionLog" rows with matching IP
func (a *DBSessionLog) ByIP(ctx context.Context, p string) ([]*savepb.SessionLog, error) {
	qn := "DBSessionLog_ByIP"
	rows, e := a.DB.QueryContext(ctx, qn, "select id,userid, username, useremail, ip, useragent, created, browserid, sessiontoken, lastused, triggerhost from "+a.SQLTablename+" where ip = $1", p)
	if e != nil {
		return nil, a.Error(ctx, qn, fmt.Errorf("ByIP: error querying (%s)", e))
	}
	defer rows.Close()
	l, e := a.FromRows(ctx, rows)
	if e != nil {
		return nil, a.Error(ctx, qn, fmt.Errorf("ByIP: error scanning (%s)", e))
	}
	return l, nil
}

// the 'like' lookup
func (a *DBSessionLog) ByLikeIP(ctx context.Context, p string) ([]*savepb.SessionLog, error) {
	qn := "DBSessionLog_ByLikeIP"
	rows, e := a.DB.QueryContext(ctx, qn, "select id,userid, username, useremail, ip, useragent, created, browserid, sessiontoken, lastused, triggerhost from "+a.SQLTablename+" where ip ilike $1", p)
	if e != nil {
		return nil, a.Error(ctx, qn, fmt.Errorf("ByIP: error querying (%s)", e))
	}
	defer rows.Close()
	l, e := a.FromRows(ctx, rows)
	if e != nil {
		return nil, a.Error(ctx, qn, fmt.Errorf("ByIP: error scanning (%s)", e))
	}
	return l, nil
}

// get all "DBSessionLog" rows with matching UserAgent
func (a *DBSessionLog) ByUserAgent(ctx context.Context, p string) ([]*savepb.SessionLog, error) {
	qn := "DBSessionLog_ByUserAgent"
	rows, e := a.DB.QueryContext(ctx, qn, "select id,userid, username, useremail, ip, useragent, created, browserid, sessiontoken, lastused, triggerhost from "+a.SQLTablename+" where useragent = $1", p)
	if e != nil {
		return nil, a.Error(ctx, qn, fmt.Errorf("ByUserAgent: error querying (%s)", e))
	}
	defer rows.Close()
	l, e := a.FromRows(ctx, rows)
	if e != nil {
		return nil, a.Error(ctx, qn, fmt.Errorf("ByUserAgent: error scanning (%s)", e))
	}
	return l, nil
}

// the 'like' lookup
func (a *DBSessionLog) ByLikeUserAgent(ctx context.Context, p string) ([]*savepb.SessionLog, error) {
	qn := "DBSessionLog_ByLikeUserAgent"
	rows, e := a.DB.QueryContext(ctx, qn, "select id,userid, username, useremail, ip, useragent, created, browserid, sessiontoken, lastused, triggerhost from "+a.SQLTablename+" where useragent ilike $1", p)
	if e != nil {
		return nil, a.Error(ctx, qn, fmt.Errorf("ByUserAgent: error querying (%s)", e))
	}
	defer rows.Close()
	l, e := a.FromRows(ctx, rows)
	if e != nil {
		return nil, a.Error(ctx, qn, fmt.Errorf("ByUserAgent: error scanning (%s)", e))
	}
	return l, nil
}

// get all "DBSessionLog" rows with matching Created
func (a *DBSessionLog) ByCreated(ctx context.Context, p uint32) ([]*savepb.SessionLog, error) {
	qn := "DBSessionLog_ByCreated"
	rows, e := a.DB.QueryContext(ctx, qn, "select id,userid, username, useremail, ip, useragent, created, browserid, sessiontoken, lastused, triggerhost from "+a.SQLTablename+" where created = $1", p)
	if e != nil {
		return nil, a.Error(ctx, qn, fmt.Errorf("ByCreated: error querying (%s)", e))
	}
	defer rows.Close()
	l, e := a.FromRows(ctx, rows)
	if e != nil {
		return nil, a.Error(ctx, qn, fmt.Errorf("ByCreated: error scanning (%s)", e))
	}
	return l, nil
}

// the 'like' lookup
func (a *DBSessionLog) ByLikeCreated(ctx context.Context, p uint32) ([]*savepb.SessionLog, error) {
	qn := "DBSessionLog_ByLikeCreated"
	rows, e := a.DB.QueryContext(ctx, qn, "select id,userid, username, useremail, ip, useragent, created, browserid, sessiontoken, lastused, triggerhost from "+a.SQLTablename+" where created ilike $1", p)
	if e != nil {
		return nil, a.Error(ctx, qn, fmt.Errorf("ByCreated: error querying (%s)", e))
	}
	defer rows.Close()
	l, e := a.FromRows(ctx, rows)
	if e != nil {
		return nil, a.Error(ctx, qn, fmt.Errorf("ByCreated: error scanning (%s)", e))
	}
	return l, nil
}

// get all "DBSessionLog" rows with matching BrowserID
func (a *DBSessionLog) ByBrowserID(ctx context.Context, p string) ([]*savepb.SessionLog, error) {
	qn := "DBSessionLog_ByBrowserID"
	rows, e := a.DB.QueryContext(ctx, qn, "select id,userid, username, useremail, ip, useragent, created, browserid, sessiontoken, lastused, triggerhost from "+a.SQLTablename+" where browserid = $1", p)
	if e != nil {
		return nil, a.Error(ctx, qn, fmt.Errorf("ByBrowserID: error querying (%s)", e))
	}
	defer rows.Close()
	l, e := a.FromRows(ctx, rows)
	if e != nil {
		return nil, a.Error(ctx, qn, fmt.Errorf("ByBrowserID: error scanning (%s)", e))
	}
	return l, nil
}

// the 'like' lookup
func (a *DBSessionLog) ByLikeBrowserID(ctx context.Context, p string) ([]*savepb.SessionLog, error) {
	qn := "DBSessionLog_ByLikeBrowserID"
	rows, e := a.DB.QueryContext(ctx, qn, "select id,userid, username, useremail, ip, useragent, created, browserid, sessiontoken, lastused, triggerhost from "+a.SQLTablename+" where browserid ilike $1", p)
	if e != nil {
		return nil, a.Error(ctx, qn, fmt.Errorf("ByBrowserID: error querying (%s)", e))
	}
	defer rows.Close()
	l, e := a.FromRows(ctx, rows)
	if e != nil {
		return nil, a.Error(ctx, qn, fmt.Errorf("ByBrowserID: error scanning (%s)", e))
	}
	return l, nil
}

// get all "DBSessionLog" rows with matching SessionToken
func (a *DBSessionLog) BySessionToken(ctx context.Context, p string) ([]*savepb.SessionLog, error) {
	qn := "DBSessionLog_BySessionToken"
	rows, e := a.DB.QueryContext(ctx, qn, "select id,userid, username, useremail, ip, useragent, created, browserid, sessiontoken, lastused, triggerhost from "+a.SQLTablename+" where sessiontoken = $1", p)
	if e != nil {
		return nil, a.Error(ctx, qn, fmt.Errorf("BySessionToken: error querying (%s)", e))
	}
	defer rows.Close()
	l, e := a.FromRows(ctx, rows)
	if e != nil {
		return nil, a.Error(ctx, qn, fmt.Errorf("BySessionToken: error scanning (%s)", e))
	}
	return l, nil
}

// the 'like' lookup
func (a *DBSessionLog) ByLikeSessionToken(ctx context.Context, p string) ([]*savepb.SessionLog, error) {
	qn := "DBSessionLog_ByLikeSessionToken"
	rows, e := a.DB.QueryContext(ctx, qn, "select id,userid, username, useremail, ip, useragent, created, browserid, sessiontoken, lastused, triggerhost from "+a.SQLTablename+" where sessiontoken ilike $1", p)
	if e != nil {
		return nil, a.Error(ctx, qn, fmt.Errorf("BySessionToken: error querying (%s)", e))
	}
	defer rows.Close()
	l, e := a.FromRows(ctx, rows)
	if e != nil {
		return nil, a.Error(ctx, qn, fmt.Errorf("BySessionToken: error scanning (%s)", e))
	}
	return l, nil
}

// get all "DBSessionLog" rows with matching LastUsed
func (a *DBSessionLog) ByLastUsed(ctx context.Context, p uint32) ([]*savepb.SessionLog, error) {
	qn := "DBSessionLog_ByLastUsed"
	rows, e := a.DB.QueryContext(ctx, qn, "select id,userid, username, useremail, ip, useragent, created, browserid, sessiontoken, lastused, triggerhost from "+a.SQLTablename+" where lastused = $1", p)
	if e != nil {
		return nil, a.Error(ctx, qn, fmt.Errorf("ByLastUsed: error querying (%s)", e))
	}
	defer rows.Close()
	l, e := a.FromRows(ctx, rows)
	if e != nil {
		return nil, a.Error(ctx, qn, fmt.Errorf("ByLastUsed: error scanning (%s)", e))
	}
	return l, nil
}

// the 'like' lookup
func (a *DBSessionLog) ByLikeLastUsed(ctx context.Context, p uint32) ([]*savepb.SessionLog, error) {
	qn := "DBSessionLog_ByLikeLastUsed"
	rows, e := a.DB.QueryContext(ctx, qn, "select id,userid, username, useremail, ip, useragent, created, browserid, sessiontoken, lastused, triggerhost from "+a.SQLTablename+" where lastused ilike $1", p)
	if e != nil {
		return nil, a.Error(ctx, qn, fmt.Errorf("ByLastUsed: error querying (%s)", e))
	}
	defer rows.Close()
	l, e := a.FromRows(ctx, rows)
	if e != nil {
		return nil, a.Error(ctx, qn, fmt.Errorf("ByLastUsed: error scanning (%s)", e))
	}
	return l, nil
}

// get all "DBSessionLog" rows with matching TriggerHost
func (a *DBSessionLog) ByTriggerHost(ctx context.Context, p string) ([]*savepb.SessionLog, error) {
	qn := "DBSessionLog_ByTriggerHost"
	rows, e := a.DB.QueryContext(ctx, qn, "select id,userid, username, useremail, ip, useragent, created, browserid, sessiontoken, lastused, triggerhost from "+a.SQLTablename+" where triggerhost = $1", p)
	if e != nil {
		return nil, a.Error(ctx, qn, fmt.Errorf("ByTriggerHost: error querying (%s)", e))
	}
	defer rows.Close()
	l, e := a.FromRows(ctx, rows)
	if e != nil {
		return nil, a.Error(ctx, qn, fmt.Errorf("ByTriggerHost: error scanning (%s)", e))
	}
	return l, nil
}

// the 'like' lookup
func (a *DBSessionLog) ByLikeTriggerHost(ctx context.Context, p string) ([]*savepb.SessionLog, error) {
	qn := "DBSessionLog_ByLikeTriggerHost"
	rows, e := a.DB.QueryContext(ctx, qn, "select id,userid, username, useremail, ip, useragent, created, browserid, sessiontoken, lastused, triggerhost from "+a.SQLTablename+" where triggerhost ilike $1", p)
	if e != nil {
		return nil, a.Error(ctx, qn, fmt.Errorf("ByTriggerHost: error querying (%s)", e))
	}
	defer rows.Close()
	l, e := a.FromRows(ctx, rows)
	if e != nil {
		return nil, a.Error(ctx, qn, fmt.Errorf("ByTriggerHost: error scanning (%s)", e))
	}
	return l, nil
}

/**********************************************************************
* Helper to convert from an SQL Query
**********************************************************************/

// from a query snippet (the part after WHERE)
func (a *DBSessionLog) FromQuery(ctx context.Context, query_where string, args ...interface{}) ([]*savepb.SessionLog, error) {
	rows, err := a.DB.QueryContext(ctx, "custom_query_"+a.Tablename(), "select "+a.SelectCols()+" from "+a.Tablename()+" where "+query_where, args...)
	if err != nil {
		return nil, err
	}
	return a.FromRows(ctx, rows)
}

/**********************************************************************
* Helper to convert from an SQL Row to struct
**********************************************************************/
func (a *DBSessionLog) Tablename() string {
	return a.SQLTablename
}

func (a *DBSessionLog) SelectCols() string {
	return "id,userid, username, useremail, ip, useragent, created, browserid, sessiontoken, lastused, triggerhost"
}
func (a *DBSessionLog) SelectColsQualified() string {
	return "" + a.SQLTablename + ".id," + a.SQLTablename + ".userid, " + a.SQLTablename + ".username, " + a.SQLTablename + ".useremail, " + a.SQLTablename + ".ip, " + a.SQLTablename + ".useragent, " + a.SQLTablename + ".created, " + a.SQLTablename + ".browserid, " + a.SQLTablename + ".sessiontoken, " + a.SQLTablename + ".lastused, " + a.SQLTablename + ".triggerhost"
}

func (a *DBSessionLog) FromRows(ctx context.Context, rows *gosql.Rows) ([]*savepb.SessionLog, error) {
	var res []*savepb.SessionLog
	for rows.Next() {
		foo := savepb.SessionLog{}
		err := rows.Scan(&foo.ID, &foo.UserID, &foo.Username, &foo.Useremail, &foo.IP, &foo.UserAgent, &foo.Created, &foo.BrowserID, &foo.SessionToken, &foo.LastUsed, &foo.TriggerHost)
		if err != nil {
			return nil, a.Error(ctx, "fromrow-scan", err)
		}
		res = append(res, &foo)
	}
	return res, nil
}

/**********************************************************************
* Helper to create table and columns
**********************************************************************/
func (a *DBSessionLog) CreateTable(ctx context.Context) error {
	csql := []string{
		`create sequence if not exists ` + a.SQLTablename + `_seq;`,
		`CREATE TABLE if not exists ` + a.SQLTablename + ` (id integer primary key default nextval('` + a.SQLTablename + `_seq'),userid text not null  ,username text not null  ,useremail text not null  ,ip text not null  ,useragent text not null  ,created integer not null  ,browserid text not null  ,sessiontoken text not null  unique  ,lastused integer not null  ,triggerhost text not null  );`,
		`CREATE TABLE if not exists ` + a.SQLTablename + `_archive (id integer primary key default nextval('` + a.SQLTablename + `_seq'),userid text not null  ,username text not null  ,useremail text not null  ,ip text not null  ,useragent text not null  ,created integer not null  ,browserid text not null  ,sessiontoken text not null  unique  ,lastused integer not null  ,triggerhost text not null  );`,
		`ALTER TABLE sessionlog ADD COLUMN IF NOT EXISTS userid text not null default '';`,
		`ALTER TABLE sessionlog ADD COLUMN IF NOT EXISTS username text not null default '';`,
		`ALTER TABLE sessionlog ADD COLUMN IF NOT EXISTS useremail text not null default '';`,
		`ALTER TABLE sessionlog ADD COLUMN IF NOT EXISTS ip text not null default '';`,
		`ALTER TABLE sessionlog ADD COLUMN IF NOT EXISTS useragent text not null default '';`,
		`ALTER TABLE sessionlog ADD COLUMN IF NOT EXISTS created integer not null default 0;`,
		`ALTER TABLE sessionlog ADD COLUMN IF NOT EXISTS browserid text not null default '';`,
		`ALTER TABLE sessionlog ADD COLUMN IF NOT EXISTS sessiontoken text not null unique  default '';`,
		`ALTER TABLE sessionlog ADD COLUMN IF NOT EXISTS lastused integer not null default 0;`,
		`ALTER TABLE sessionlog ADD COLUMN IF NOT EXISTS triggerhost text not null default '';`,
	}
	for i, c := range csql {
		_, e := a.DB.ExecContext(ctx, fmt.Sprintf("create_"+a.SQLTablename+"_%d", i), c)
		if e != nil {
			return e
		}
	}
	return nil
}

/**********************************************************************
* Helper to meaningful errors
**********************************************************************/
func (a *DBSessionLog) Error(ctx context.Context, q string, e error) error {
	if e == nil {
		return nil
	}
	return fmt.Errorf("[table="+a.SQLTablename+", query=%s] Error: %s", q, e)
}
