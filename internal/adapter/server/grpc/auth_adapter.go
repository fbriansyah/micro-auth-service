package grpcserver

import (
	"context"

	"github.com/fbriansyah/micro-payment-proto/protogen/go/auth"
	"github.com/fbriansyah/micro-payment-proto/protogen/go/session"
)

func (a *GrpcServerAdapter) Login(context.Context, *auth.LoginRequest) (*auth.LoginResponse, error) {

	return &auth.LoginResponse{
		Userid:  "",
		Name:    "",
		Session: &session.Session{},
	}, nil
}

func (a *GrpcServerAdapter) CreateUser(context.Context, *auth.CreateUserRequest) (*auth.CreateUserResponse, error) {

	return &auth.CreateUserResponse{
		Userid:   "",
		Username: "",
		Name:     "",
	}, nil
}
