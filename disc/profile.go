package disc

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/Karitham/WaifuBot/anilist"
	"github.com/Karitham/WaifuBot/db"
	"github.com/diamondburned/arikawa/v2/bot/extras/arguments"
	"github.com/diamondburned/arikawa/v2/discord"
	"github.com/diamondburned/arikawa/v2/gateway"
	"github.com/lib/pq"
	"github.com/rs/zerolog/log"
)

// Profile displays user profile
func (b *Bot) Profile(m *gateway.MessageCreateEvent, _ ...*arguments.UserMention) (*discord.Embed, error) {
	user := parseUser(m)

	data, err := b.conn.GetUserProfile(context.Background(), int64(user.ID))
	if err == sql.ErrNoRows {
		err = b.conn.CreateUser(b.Ctx.Context(), int64(user.ID))
		if err != nil {
			log.Err(err).Msg("Error creating user")
			return nil, err
		}
	} else if err != nil {
		log.Debug().
			Err(err).
			Str("Type", "PROFILE").
			Msg("Error getting user profile")

		return nil, err
	}

	log.Trace().
		Str("Type", "PROFILE").
		Int("User", int(user.ID)).
		Str("Quote", data.Quote).
		Str("Name", data.Name.String).
		Msg("sent profile embed")

	return &discord.Embed{
		Title: fmt.Sprintf("%s's profile", user.Username),

		Description: fmt.Sprintf(
			"%s\n%s last rolled %s ago.\nThey own %d waifus.\nTheir Favorite waifu is %s",
			data.Quote,
			user.Username,
			time.Since(data.Date.UTC()).Truncate(time.Second),
			data.Count,
			data.Name.String,
		),

		Thumbnail: &discord.EmbedThumbnail{URL: data.Image.String},
	}, nil
}

// Favorite sets a waifu as favorite
func (b *Bot) Favorite(m *gateway.MessageCreateEvent, name ...Name) (string, error) {
	if len(name) == 0 {
		return "", errors.New("no character name entered")
	}
	n := strings.Join(name, " ")

	id := parseArgs(n)
	searchArgs := anilist.CharSearchInput{
		ID:   id,
		Name: n,
	}

	char, err := anilist.CharSearch(searchArgs)
	if err != nil {
		log.Debug().
			Err(err).
			Str("Type", "FAVORITE").
			Msg("Error Searching anilist")

		return "", err
	}

	err = b.conn.SetFavorite(context.Background(), db.SetFavoriteParams{
		ID: int64(m.Author.ID),
		Favorite: sql.NullInt64{
			Int64: int64(char.Character.ID),
			Valid: true,
		},
	})
	if err, ok := err.(*pq.Error); ok && err.Code == "23503" {
		log.Debug().
			Err(err).
			Str("Type", "FAVORITE").
			Msg("Error setting favorite")

		return "", errors.New("You do not own this character")

	} else if err != nil {
		log.Err(err).
			Str("Type", "FAVORITE").
			Msg("Error setting favorite")

		return "", err
	}

	log.Trace().
		Str("Type", "FAVORITE").
		Int("User", int(m.Author.ID)).
		Int("Character", int(char.Character.ID)).
		Msg("updated favorite")

	return fmt.Sprintf("New waifu set, check your profile\n<%s>", char.Character.SiteURL), nil
}

// Quote sets a quote on the user profile
func (b *Bot) Quote(m *gateway.MessageCreateEvent, quote ...string) (string, error) {
	if quote == nil {
		return "", errors.New("no quote entered")
	}
	q := strings.Join(quote, " ")

	err := b.conn.SetQuote(context.Background(), db.SetQuoteParams{
		ID:    int64(m.Author.ID),
		Quote: q,
	})
	if err != nil {
		log.Debug().
			Err(err).
			Str("Type", "QUOTE").
			Msg("Error setting quote")

		return "", err
	}

	log.Trace().
		Str("Type", "QUOTE").
		Int("User", int(m.Author.ID)).
		Str("Quote", q).
		Msg("updated quote")

	return fmt.Sprintf("New quote set :\n%s", q), nil
}
