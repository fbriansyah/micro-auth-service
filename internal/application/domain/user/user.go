package dmuser

import (
	dmsession "github.com/fbriansyah/micro-auth-service/internal/application/domain/session"
	"github.com/google/uuid"
)

type User struct {
	ID       uuid.UUID
	Username string
	Fullname string
	Session  dmsession.Session
}
