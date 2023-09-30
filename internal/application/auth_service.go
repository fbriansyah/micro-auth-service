package application

import (
	"context"
	"errors"

	"github.com/fbriansyah/micro-auth-service/internal/adapter/postgresdb"
	dmsession "github.com/fbriansyah/micro-auth-service/internal/application/domain/session"
	dmuser "github.com/fbriansyah/micro-auth-service/internal/application/domain/user"
	"github.com/fbriansyah/micro-auth-service/internal/port"
	"github.com/fbriansyah/micro-auth-service/util"
	"github.com/fbriansyah/micro-payment-proto/protogen/go/session"
)

var (
	ErrorWrongPassword        = errors.New("wrong password")
	ErrorUserNotActive        = errors.New("user is not active")
	ErrorGenerateHashPassword = errors.New("failed to create password hash")
	ErrorCreatingSession      = errors.New("failed to create session")
)

type AuthService struct {
	db            port.DatabasePort
	sessionClient port.SessionPort
}

func NewAuthService(db port.DatabasePort, sessionClient port.SessionPort) *AuthService {
	return &AuthService{
		db:            db,
		sessionClient: sessionClient,
	}
}

// Login check username, password and . This method call create session rpc.
func (s *AuthService) Login(username, password string) (dmuser.User, error) {

	user, err := s.db.GetUserByUsername(context.Background(), username)
	if err != nil {
		return dmuser.User{}, err
	}

	err = util.CheckPassword(password, user.Password)
	if err != nil {
		return dmuser.User{}, ErrorWrongPassword
	}

	if !user.IsActive {
		return dmuser.User{}, ErrorUserNotActive
	}

	// TODO: Call CreateSession from session micro service
	session, err := s.sessionClient.CreateSession(context.Background(), &session.UserID{
		UserId: user.ID.String(),
	})

	if err != nil {
		return dmuser.User{}, nil
	}

	return dmuser.User{
		ID:       user.ID,
		Username: username,
		Fullname: user.Fullname,
		Session: dmsession.Session{
			Id:                    session.Id,
			UserId:                user.ID.String(),
			AccessToken:           session.AccessToken,
			RefreshToken:          session.AccessToken,
			AccessTokenExpiresAt:  session.AccessTokenExpiresAt.String(),
			RefreshTokenExpiresAt: session.RefreshTokenExpiresAt.String(),
		},
	}, nil
}

// Register method, save user to database
func (s *AuthService) Register(user dmuser.User, password string) (dmuser.User, error) {
	hashedPassword, err := util.HashPassword(password)
	if err != nil {
		return dmuser.User{}, ErrorGenerateHashPassword
	}
	arg := postgresdb.CreateUserParams{
		Username: user.Username,
		Fullname: user.Fullname,
		Password: hashedPassword,
	}

	usr, err := s.db.CreateUser(context.Background(), arg)
	if err != nil {
		return dmuser.User{}, err
	}

	return dmuser.User{
		ID:       usr.ID,
		Username: usr.Username,
		Fullname: usr.Fullname,
	}, nil
}
