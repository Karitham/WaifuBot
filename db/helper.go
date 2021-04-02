package db

import (
	"context"
	"database/sql"
	"errors"
	"time"
)

func (q *Queries) RollChar(ctx context.Context, userID, charID int64, image, name string) error {
	return q.AsTx(func(q *Queries) error {
		err := q.InsertChar(ctx, InsertCharParams{
			UserID: userID,
			ID:     charID,
			Image:  sql.NullString{String: image, Valid: true},
			Name:   sql.NullString{String: name, Valid: true},
		})
		if err != nil {
			return err
		}

		err = q.SetDate(ctx, SetDateParams{ID: userID, Date: time.Now().UTC()})
		if err != nil {
			return err
		}
		return nil
	})
}

// AsTx creates a new transaction from the Queries struct, executes and commits it.
// Pass in a function that does multiple queries.
// It will be rolled back on error, or commited if all goes well.
func (q *Queries) AsTx(queries func(*Queries) error) error {
	db, ok := q.db.(*sql.DB)
	if !ok {
		return errors.New("db is not a *sql.DB:")
	}

	tx, err := db.Begin()
	if err != nil {
		return err
	}

	defer func() {
		if err != nil {
			_ = tx.Rollback()
			return
		}
		_ = tx.Commit()
	}()

	return queries(New(db).WithTx(tx))
}
