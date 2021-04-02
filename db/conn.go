package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"

	"github.com/Karitham/WaifuBot/config"
)

// Init initialises the connetion with the db
func Init(conf config.Database) (*Queries, error) {
	connStr := fmt.Sprintf("user=%s dbname=%s password=%s host=%s sslmode=disable", conf.User, conf.Dbname, conf.Password, conf.Host)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	return New(db), nil
}
