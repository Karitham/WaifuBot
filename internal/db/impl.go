package db

import (
	"context"
	"database/sql"
	"errors"
	"strings"
	"time"

	"github.com/Karitham/WaifuBot/internal/discord"
	"github.com/Karitham/corde"
	"github.com/Masterminds/squirrel"
)

var _ discord.Store = (*Queries)(nil)

// PutChar a char in the database
func (q *Queries) PutChar(ctx context.Context, userID corde.Snowflake, c discord.Character) error {
	if q.tx == nil {
		return q.asTx(func(q *Queries) error {
			return q.PutChar(ctx, userID, c)
		})
	}

	_, err := q.getUser(ctx, uint64(userID))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			if err := q.createUser(ctx, uint64(userID)); err != nil {
				return err
			}
		} else {
			return err
		}
	}

	p := insertCharParams{
		ID:     c.ID,
		UserID: uint64(c.UserID),
		Image:  c.Image,
		Name:   strings.Join(strings.Fields(c.Name), " "),
		Type:   c.Type,
	}

	return q.insertChar(ctx, p)
}

func (q *Queries) SetUserAnilistURL(ctx context.Context, userID corde.Snowflake, url string) error {
	return q.updateUser(ctx, userID, withAnilistURL(url))
}

// Chars returns the user's characters
func (q *Queries) Chars(ctx context.Context, userID corde.Snowflake) ([]discord.Character, error) {
	dbchs, err := q.getChars(ctx, uint64(userID))
	if err != nil {
		return nil, err
	}

	chars := make([]discord.Character, 0, len(dbchs))
	for _, c := range dbchs {
		chars = append(chars, discord.Character{
			Date:   c.Date,
			Image:  c.Image,
			Name:   c.Name,
			Type:   c.Type,
			UserID: corde.Snowflake(c.UserID),
			ID:     c.ID,
		})
	}

	return chars, nil
}

// CharsIDs returns the user's character's ID
func (q *Queries) CharsIDs(ctx context.Context, userID corde.Snowflake) ([]int64, error) {
	return q.getCharsID(ctx, uint64(userID))
}

// User returns a user
func (q *Queries) User(ctx context.Context, userID corde.Snowflake) (discord.User, error) {
	u, err := q.getUser(ctx, uint64(userID))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			if err := q.createUser(ctx, uint64(userID)); err != nil {
				return discord.User{}, err
			}
		} else {
			return discord.User{}, err
		}
	}

	return discord.User{
		Date:     u.Date,
		Quote:    u.Quote,
		UserID:   corde.Snowflake(u.UserID),
		Favorite: uint64(u.Favorite.Int64),
		Tokens:   u.Tokens,
	}, nil
}

// updateUser updates a user's properties
func (q *Queries) updateUser(ctx context.Context, userID corde.Snowflake, opts ...func(*squirrel.UpdateBuilder)) error {
	if q.tx == nil {
		return q.asTx(func(q *Queries) error {
			return q.updateUser(ctx, userID, opts...)
		})
	}

	_, err := q.getUser(ctx, uint64(userID))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			if err := q.createUser(ctx, uint64(userID)); err != nil {
				return err
			}
		} else {
			return err
		}
	}

	builder := squirrel.Update("users").Where(squirrel.Eq{
		"user_id": userID,
	}).PlaceholderFormat(squirrel.Dollar)

	for _, opt := range opts {
		opt(&builder)
	}

	str, args, err := builder.ToSql()
	if err != nil {
		return err
	}

	if _, err := q.exec(ctx, nil, str, args...); err != nil {
		return err
	}

	return nil
}

// withFavorite sets user favorite
func withFav(f int64) func(*squirrel.UpdateBuilder) {
	return func(s *squirrel.UpdateBuilder) {
		*s = s.Set("favorite", f)
	}
}

// withQuote sets user quote
func withQuote(q string) func(*squirrel.UpdateBuilder) {
	return func(s *squirrel.UpdateBuilder) {
		*s = s.Set("quote", q)
	}
}

// withDate sets the date
func withDate(d time.Time) func(*squirrel.UpdateBuilder) {
	return func(s *squirrel.UpdateBuilder) {
		*s = s.Set("date", d.UTC())
	}
}

// withAnilistURL sets the anilist url
func withAnilistURL(url string) func(*squirrel.UpdateBuilder) {
	return func(s *squirrel.UpdateBuilder) {
		*s = s.Set("anilist_url", url)
	}
}

// withToken sets the token
func withToken(t string) func(*squirrel.UpdateBuilder) {
	return func(s *squirrel.UpdateBuilder) {
		*s = s.Set("token", t)
	}
}

// SetUserDate sets the user's date
func (q *Queries) SetUserDate(ctx context.Context, userID corde.Snowflake, d time.Time) error {
	return q.updateUser(ctx, userID, withDate(d))
}

// SetUserToken sets the user's token
func (q *Queries) SetUserToken(ctx context.Context, userID corde.Snowflake, token string) error {
	return q.updateUser(ctx, userID, withToken(token))
}

// SetUserFavorite sets the user's favorite
func (q *Queries) SetUserFavorite(ctx context.Context, userID corde.Snowflake, c int64) error {
	return q.updateUser(ctx, userID, withFav(c))
}

// SetUserQuote sets the user's quote
func (q *Queries) SetUserQuote(ctx context.Context, userID corde.Snowflake, quote string) error {
	return q.updateUser(ctx, userID, withQuote(quote))
}

// CharsStartingWith returns characters starting with the given string
func (q *Queries) CharsStartingWith(ctx context.Context, userID corde.Snowflake, s string) ([]discord.Character, error) {
	_, err := q.getUser(ctx, uint64(userID))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			if err := q.createUser(ctx, uint64(userID)); err != nil {
				return nil, err
			}
		} else {
			return nil, err
		}
	}

	dbchs, err := q.getCharsWhoseIDStartWith(ctx, getCharsWhoseIDStartWithParams{
		UserID:  uint64(userID),
		Lim:     50,
		Off:     0,
		LikeStr: s + "%",
	})
	if err != nil {
		return nil, err
	}

	chars := make([]discord.Character, 0, len(dbchs))
	for _, c := range dbchs {
		chars = append(chars, discord.Character{
			Date:   c.Date,
			Image:  c.Image,
			Name:   c.Name,
			Type:   c.Type,
			UserID: corde.Snowflake(c.UserID),
			ID:     c.ID,
		})
	}

	return chars, nil
}

// Profile returns the user's profile
func (q *Queries) Profile(ctx context.Context, userID corde.Snowflake) (discord.Profile, error) {
	_, err := q.getUser(ctx, uint64(userID))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			if err := q.createUser(ctx, uint64(userID)); err != nil {
				return discord.Profile{}, err
			}
		} else {
			return discord.Profile{}, err
		}
	}

	p, err := q.getProfile(ctx, uint64(userID))
	if err != nil {
		return discord.Profile{}, err
	}

	return discord.Profile{
		User: discord.User{
			Date:       p.UserDate,
			Quote:      p.UserQuote,
			UserID:     corde.Snowflake(p.UserID),
			Tokens:     p.UserTokens,
			AnilistURL: p.UserAnilistUrl,
			Favorite:   uint64(p.FavoriteID.Int64),
		},
		CharacterCount: int(p.Count),
		Favorite: discord.Character{
			Image:  p.FavoriteImage.String,
			Name:   p.FavoriteName.String,
			UserID: userID,
			ID:     p.FavoriteID.Int64,
		},
	}, nil
}

func (q *Queries) GiveUserChar(ctx context.Context, dst corde.Snowflake, src corde.Snowflake, charID int64) error {
	_, err := q.giveChar(ctx, giveCharParams{
		Given: int64(dst),
		ID:    charID,
		Giver: int64(src),
	})
	return err
}

func (q *Queries) VerifyChar(ctx context.Context, userID corde.Snowflake, charID int64) (discord.Character, error) {
	c, err := q.getChar(ctx, getCharParams{
		ID:     charID,
		UserID: uint64(userID),
	})
	if err != nil {
		return discord.Character{}, err
	}

	return discord.Character{
		Date:   c.Date,
		Image:  c.Image,
		Name:   c.Name,
		Type:   c.Type,
		UserID: corde.Snowflake(c.UserID),
		ID:     c.ID,
	}, nil
}

func (q *Queries) ConsumeDropTokens(ctx context.Context, userID corde.Snowflake, count int32) (discord.User, error) {
	u, err := q.consumeDropTokens(ctx, consumeDropTokensParams{
		Tokens: count,
		UserID: uint64(userID),
	})
	if err != nil {
		return discord.User{}, err
	}
	return discord.User{
		Date:     u.Date,
		Quote:    u.Quote,
		Favorite: uint64(u.Favorite.Int64),
		UserID:   userID,
		Tokens:   u.Tokens,
	}, nil
}

func (q *Queries) AddDropToken(ctx context.Context, userID corde.Snowflake) error {
	return q.addDropToken(ctx, uint64(userID))
}

func (q *Queries) DeleteChar(ctx context.Context, userID corde.Snowflake, charID int64) (discord.Character, error) {
	c, err := q.deleteChar(ctx, deleteCharParams{
		UserID: uint64(userID),
		ID:     charID,
	})
	if err != nil {
		return discord.Character{}, err
	}

	return discord.Character{
		Date:   c.Date,
		Image:  c.Image,
		Name:   c.Name,
		Type:   c.Type,
		UserID: corde.Snowflake(c.UserID),
		ID:     c.ID,
	}, nil
}
