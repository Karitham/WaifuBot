package db

import (
	"database/sql"
	"errors"
)

func (q *Queries) Tx(fn func(Querier) error) error {
	db, ok := q.db.(*sql.DB)
	if !ok {
		return errors.New("invalid database provided")
	}

	tx, err := db.Begin()
	if err != nil {
		return err
	}

	err = fn(q.WithTx(tx))
	if err != nil {
		if err = tx.Rollback(); err != nil {
			return err
		}
		return err
	}

	return tx.Commit()
}
