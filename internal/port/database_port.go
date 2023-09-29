package port

import (
	"context"

	"github.com/fbriansyah/micro-auth-service/internal/adapter/postgresdb"
)

type DatabasePort interface {
	CreateUser(ctx context.Context, arg postgresdb.CreateUserParams) (postgresdb.User, error)
	GetUserByUsername(ctx context.Context, username string) (postgresdb.User, error)
}
