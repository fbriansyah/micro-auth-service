package postgresdb

import (
	"database/sql"
)

type DatabaseAdapter interface {
	Querier
}

type DatabaseStore struct {
	db *sql.DB
	*Queries
}

func NewDatabaseAdapter(db *sql.DB) DatabaseAdapter {
	return &DatabaseStore{
		db:      db,
		Queries: New(db),
	}
}
