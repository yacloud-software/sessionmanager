package main

import (
	"context"

	"golang.conradwood.net/apis/common"
	"golang.conradwood.net/go-easyops/auth"
	"golang.conradwood.net/go-easyops/errors"
	pb "golang.yacloud.eu/apis/sessionmanager"
)

func (e *echoServer) WhoAmI(ctx context.Context, req *common.Void) (*pb.WhoAmIResponse, error) {
	u := auth.GetUser(ctx)
	res := &pb.WhoAmIResponse{
		User: u,
	}

	return res, nil
}
func (e *echoServer) WhoAmIWithTriggerAuth(ctx context.Context, req *common.Void) (*pb.WhoAmIResponse, error) {
	u := auth.GetUser(ctx)
	if u == nil {
		return nil, errors.Unauthenticated(ctx, "please log in")
	}
	res := &pb.WhoAmIResponse{
		User: u,
	}

	return res, nil
}
