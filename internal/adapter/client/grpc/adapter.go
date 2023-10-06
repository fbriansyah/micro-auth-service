package grpcclient

import (
	"context"

	dmsession "github.com/fbriansyah/micro-auth-service/internal/application/domain/session"
	"github.com/fbriansyah/micro-auth-service/internal/port"
	"github.com/fbriansyah/micro-payment-proto/protogen/go/session"
	"google.golang.org/grpc"
)

type SessionClientAdapter struct {
	sessionClient port.SessionPort
}

func NewSessionClientAdapter(conn *grpc.ClientConn) *SessionClientAdapter {
	client := session.NewSessionServiceClient(conn)

	return &SessionClientAdapter{sessionClient: client}
}

// CreateSession create rpc call to session micro service
func (a *SessionClientAdapter) CreateSession(ctx context.Context, userID string) (dmsession.Session, error) {
	session, err := a.sessionClient.CreateSession(ctx, &session.UserID{
		UserId: userID,
	})

	if err != nil {
		return dmsession.Session{}, nil
	}

	return dmsession.Session{
		Id:                    session.Id,
		UserId:                userID,
		AccessToken:           session.AccessToken,
		RefreshToken:          session.RefreshToken,
		AccessTokenExpiresAt:  session.AccessTokenExpiresAt,
		RefreshTokenExpiresAt: session.RefreshTokenExpiresAt,
	}, nil
}
