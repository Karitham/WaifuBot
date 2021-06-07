package db

import (
	"context"

	sq "github.com/Masterminds/squirrel"
)

// UpdateUser updates the database user from the passed in user
func (q *Queries) UpdateUser(ctx context.Context, user User) error {
	s := sq.Update("users").Where(sq.Eq{"user_id": user.UserID}).PlaceholderFormat(sq.Dollar)

	if !user.Date.IsZero() {
		s = s.Set("date", user.Date)
	}
	if user.Favorite.Valid {
		s = s.Set("favorite", user.Favorite)
	}
	if user.Quote != "" {
		s = s.Set("quote", user.Quote)
	}

	sql, args, err := s.ToSql()
	if err != nil {
		return err
	}

	_, err = q.db.ExecContext(ctx, sql, args...)
	if err != nil {
		return err
	}

	return nil
}
