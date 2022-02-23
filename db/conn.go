package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

// Init initialises the connetion with the db
func Init(url string) (*Queries, error) {
	db, err := sql.Open("postgres", url)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return New(db), nil
}
