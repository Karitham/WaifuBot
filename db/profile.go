package db

import (
	"context"
	"database/sql"
	"time"
)

func (q *Queries) GetProfile(ctx context.Context, userID int64) (Profile, error) {
	row := q.db.QueryRowContext(ctx, getProfile, userID)
	var p Profile
	var ft favTmp
	err := row.Scan(
		&ft.Image,
		&ft.Name,
		&ft.ID,
		&p.Date,
		&p.Quote,
		&p.UserID,
		&p.Count,
	)
	if err != nil {
		return Profile{}, err
	}

	p.Favorite.ID = ft.ID.Int64
	p.Favorite.Image = ft.Image.String
	p.Favorite.Name = ft.Name.String

	return p, nil
}

type Profile struct {
	Date     time.Time `json:"user_date"`
	Quote    string    `json:"user_quote"`
	Favorite Favorite  `json:"favorite"`
	UserID   int64     `json:"user_id"`
	Count    int64     `json:"count"`
}

type favTmp struct {
	Name  sql.NullString `json:"name"`
	Image sql.NullString `json:"image"`
	ID    sql.NullInt64  `json:"id"`
}

type Favorite struct {
	Name  string `json:"name"`
	Image string `json:"image"`
	ID    int64  `json:"id"`
}

const getProfile = `-- name: getProfile :one
SELECT characters.image as favorite_image,
    characters.name as favorite_name,
    characters.id as favorite_id,
    users.date as user_date,
    users.quote as user_quote,
    users.user_id as user_id,
    (
        SELECT count(*)
        FROM characters
        WHERE characters.user_id = $1
    ) as count
FROM users
    LEFT JOIN characters ON characters.id = users.favorite
WHERE users.user_id = $1
`
