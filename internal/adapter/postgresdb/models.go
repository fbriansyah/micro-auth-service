// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0

package postgresdb

import (
	"github.com/google/uuid"
)

type User struct {
	ID       uuid.UUID `json:"id"`
	Fullname string    `json:"fullname"`
	Username string    `json:"username"`
	Password string    `json:"password"`
	IsActive bool      `json:"is_active"`
}
