package grpcserver

import (
	"context"
	"fmt"

	dmuser "github.com/fbriansyah/micro-auth-service/internal/application/domain/user"
	"github.com/fbriansyah/micro-payment-proto/protogen/go/auth"
	"google.golang.org/grpc/codes"
)

// Login method is implementation of rpc Login in AuthService. It calls Login method in service
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
	}, nil
}

// CreateUser is implementation of rpc CreateUser in AuthService. It calls Register methon in service
func (a *GrpcServerAdapter) CreateUser(ctx context.Context, req *auth.CreateUserRequest) (*auth.CreateUserResponse, error) {
	user, err := a.service.Register(
		ctx,
		dmuser.User{
			Username: req.Username,
			Fullname: req.Name,
		},
		req.Password,
	)
	if err != nil {
		return nil, generateError(
			codes.FailedPrecondition,
			fmt.Sprintf("error create user: %v", err),
		)
	}

	return &auth.CreateUserResponse{
		Userid:   user.ID.String(),
		Username: user.Username,
		Name:     user.Fullname,
	}, nil
}
