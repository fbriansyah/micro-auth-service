package grpcserver

import (
	"context"
	"fmt"

	"github.com/fbriansyah/micro-payment-proto/protogen/go/auth"
	"github.com/fbriansyah/micro-payment-proto/protogen/go/session"
	"google.golang.org/grpc/codes"
)

func (a *GrpcServerAdapter) Login(ctx context.Context, req *auth.LoginRequest) (*auth.LoginResponse, error) {
	user, err := a.service.Login(ctx, req.Username, req.Password)
	if err != nil {
		return nil, generateError(
			codes.FailedPrecondition,
			fmt.Sprintf("error login: %v", err),
		)
	}

	return &auth.LoginResponse{
		Userid: user.ID.String(),
		Name:   user.Fullname,
		Session: &session.Session{
			Id:                    user.Session.Id,
			UserId:                user.ID.String(),
			AccessToken:           user.Session.AccessToken,
			RefreshToken:          user.Session.RefreshToken,
			AccessTokenExpiresAt:  user.Session.AccessTokenExpiresAt,
			RefreshTokenExpiresAt: user.Session.RefreshTokenExpiresAt,
		},
	}, nil
}

func (a *GrpcServerAdapter) CreateUser(context.Context, *auth.CreateUserRequest) (*auth.CreateUserResponse, error) {

	return &auth.CreateUserResponse{
		Userid:   "",
		Username: "",
		Name:     "",
	}, nil
}
