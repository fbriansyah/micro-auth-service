// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0

package postgresdb

import (
	"context"
)

type Querier interface {
	CreateUser(ctx context.Context, arg CreateUserParams) (User, error)
	GetUserByUsername(ctx context.Context, username string) (User, error)
}

var _ Querier = (*Queries)(nil)
