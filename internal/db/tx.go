package db

import (
	"database/sql"
	"errors"

	"github.com/Karitham/WaifuBot/internal/discord"
)

// Tx executes a function in a transaction.
func (q *Queries) Tx(fn func(s discord.Store) error) error {
	return q.asTx(func(q *Queries) error {
		return fn(q)
	})
}

func (q *Queries) asTx(fn func(q *Queries) error) error {
	db, ok := q.db.(*sql.DB)
	if !ok {
		return errors.New("invalid database provided")
	}

	tx, err := db.Begin()
	if err != nil {
		return err
	}

	queriesErr := fn(q.WithTx(tx))
	if queriesErr != nil {
		if err = tx.Rollback(); err != nil {
			return queriesErr
		}
		return queriesErr
	}

	return tx.Commit()
}
