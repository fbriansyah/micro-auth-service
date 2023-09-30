package grpcclient

import (
	"context"

	dmsession "github.com/fbriansyah/micro-auth-service/internal/application/domain/session"
	"github.com/fbriansyah/micro-auth-service/internal/port"
	"github.com/fbriansyah/micro-payment-proto/protogen/go/session"
	"google.golang.org/grpc"
)

type SessionAdapterClient struct {
	sessionClient port.SessionPort
}

func NewSessionAdapterClient(conn *grpc.ClientConn) *SessionAdapterClient {
	client := session.NewSessionServiceClient(conn)

	return &SessionAdapterClient{sessionClient: client}
}

// CreateSession create rpc call to session micro service
func (a *SessionAdapterClient) CreateSession(userID string) (dmsession.Session, error) {
	session, err := a.sessionClient.CreateSession(context.Background(), &session.UserID{
		UserId: userID,
	})

	if err != nil {
		return dmsession.Session{}, nil
	}

	return dmsession.Session{
		Id:                    session.Id,
		UserId:                userID,
		AccessToken:           session.AccessToken,
		RefreshToken:          session.AccessToken,
		AccessTokenExpiresAt:  session.AccessTokenExpiresAt.String(),
		RefreshTokenExpiresAt: session.RefreshTokenExpiresAt.String(),
	}, nil
}
