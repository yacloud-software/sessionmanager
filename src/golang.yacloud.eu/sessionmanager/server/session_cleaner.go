package main

import (
	"fmt"
	"golang.conradwood.net/go-easyops/authremote"
	"golang.conradwood.net/go-easyops/utils"
	pb "golang.yacloud.eu/apis/sessionmanager"
	"golang.yacloud.eu/sessionmanager/db"
	"time"
)

const (
	OLDEST_LASTUSED_NOUSER   = time.Duration(30) * time.Minute
	OLDEST_LASTUSED_WITHUSER = time.Duration(48) * time.Hour
)

func session_cleaner() {
	t := time.Duration(2) * time.Second
	for {
		time.Sleep(t)
		t = time.Duration(180) * time.Second
		err := clean()
		if err != nil {
			fmt.Printf("Failed to clean: %s\n", err)
		}
	}
}
func clean() error {
	now := time.Now()

	// delete those with no user
	t_cutoff := now.Add(0 - OLDEST_LASTUSED_NOUSER)
	cutoff := t_cutoff.Unix()
	ctx := authremote.Context()
	entries, err := db.DefaultDBSessionLog().FromQuery(ctx, "userid = '' and lastused < $1", cutoff)
	if err != nil {
		return err
	}
	fmt.Printf("[cleaner] Deleting %d entries with no userid and lastused < %s\n", len(entries), utils.TimeString(t_cutoff))
	err = remove(entries)
	if err != nil {
		return err
	}

	// delete those with user
	t_cutoff = now.Add(0 - OLDEST_LASTUSED_WITHUSER)
	cutoff = t_cutoff.Unix()
	ctx = authremote.Context()
	entries, err = db.DefaultDBSessionLog().FromQuery(ctx, "lastused < $1", cutoff)
	if err != nil {
		return err
	}
	fmt.Printf("[cleaner] Deleting %d entries with userid and lastused < %s\n", len(entries), utils.TimeString(t_cutoff))
	err = remove(entries)
	if err != nil {
		return err
	}
	return nil
}
func remove(entries []*pb.SessionLog) error {
	ctx := authremote.Context()
	ct := 0
	// not re-creating context here to avoid starvation.
	for _, e := range entries {
		err := db.DefaultDBSessionLog().DeleteByID(ctx, e.ID)
		if err != nil {
			fmt.Printf("[cleaner] After deleting %d entries, error: %s\n", ct, err)
			return err
		}
		ct++
	}
	fmt.Printf("[cleaner] cleaned %d entries\n", ct)
	return nil
}
