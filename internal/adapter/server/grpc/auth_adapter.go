package grpcserver

import (
	"context"
	"fmt"

	dmuser "github.com/fbriansyah/micro-auth-service/internal/application/domain/user"
	"github.com/fbriansyah/micro-payment-proto/protogen/go/auth"
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
	}, nil
}

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
