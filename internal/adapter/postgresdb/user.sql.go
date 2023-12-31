// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0
// source: user.sql

package postgresdb

import (
	"context"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users (username, password, fullname)
VALUES ($1, $2, $3)
RETURNING id, fullname, username, password, is_active
`

type CreateUserParams struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Fullname string `json:"fullname"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, createUser, arg.Username, arg.Password, arg.Fullname)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Fullname,
		&i.Username,
		&i.Password,
		&i.IsActive,
	)
	return i, err
}

const getUserByUsername = `-- name: GetUserByUsername :one
SELECT id, fullname, username, password, is_active FROM users WHERE username = $1 LIMIT 1
`

func (q *Queries) GetUserByUsername(ctx context.Context, username string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserByUsername, username)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Fullname,
		&i.Username,
		&i.Password,
		&i.IsActive,
	)
	return i, err
}
