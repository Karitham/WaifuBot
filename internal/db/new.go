package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

// NewDB initialises the connetion with the db
func NewDB(connstr string) (*Queries, error) {
	db, err := sql.Open("postgres", connstr)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return New(db), nil
}
