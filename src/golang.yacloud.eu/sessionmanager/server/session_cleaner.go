package main

import (
	"fmt"
	"golang.conradwood.net/go-easyops/authremote"
	"golang.yacloud.eu/sessionmanager/db"
	"time"
)

const (
	OLDEST_LASTUSED_NOUSER = time.Duration(30) * time.Minute
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
	cutoff := now.Add(0 - OLDEST_LASTUSED_NOUSER).Unix()
	ctx := authremote.Context()
	entries, err := db.DefaultDBSessionLog().FromQuery(ctx, "userid = '' and lastused < $1", cutoff)
	if err != nil {
		return err
	}
	fmt.Printf("Deleting %d entries with no userid and lastused < %d\n", len(entries), cutoff)
	ct := 0
	// not re-creating context here to avoid starvation.
	for _, e := range entries {
		err = db.DefaultDBSessionLog().DeleteByID(ctx, e.ID)
		if err != nil {
			fmt.Printf("After deleting %d entries, error: %s\n", ct, err)
			return err
		}
		ct++
	}
	return nil
}
