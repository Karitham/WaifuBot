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

	data, err := b.DB.GetProfile(context.Background(), int64(user.ID))
	if err == sql.ErrNoRows {
		err = b.DB.CreateUser(b.Ctx.Context(), int64(user.ID))
		if err != nil {
			log.Err(err).Msg("Error creating user")
			return nil, errors.New("error creating your profile")
		}
	} else if err != nil {
		log.Debug().
			Err(err).
			Str("Type", "PROFILE").
			Msg("Error getting user profile")

		return nil, errors.New("error getting your profile")
	}

	return &discord.Embed{
		Title: fmt.Sprintf("%s's profile", user.Username),
		Description: fmt.Sprintf(
			"%s\n%s last rolled %s ago.\nThey own %d waifus.\nTheir Favorite waifu is %s",
			data.Quote,
			user.Username,
			time.Since(data.Date.UTC()).Truncate(time.Second),
			data.Count,
			data.Favorite.Name,
		),
		Thumbnail: &discord.EmbedThumbnail{URL: data.Favorite.Image},
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

		return "", errors.New("error getting a character, please check the ID/Name entered is right and retry later")
	}

	err = b.DB.UpdateUser(context.Background(), db.User{
		Favorite: sql.NullInt64{
			Int64: int64(char.Character.ID),
			Valid: true,
		},
		UserID: int64(m.Author.ID),
	})
	if err, ok := err.(*pq.Error); ok && err.Code == "23503" {
		log.Debug().
			Err(err).
			Str("Type", "FAVORITE").
			Msg("Error setting favorite")

		return "", errors.New("you do not own this character")
	} else if err != nil {
		log.Err(err).
			Str("Type", "FAVORITE").
			Msg("Error setting favorite")

		return "", errors.New("an error occured setting this character as favorite, please retry later or raise an issue on https://github.com/Karitham/WaifuBot")
	}

	return fmt.Sprintf("New waifu set, check your profile\n<%s>", char.Character.SiteURL), nil
}

// Quote sets a quote on the user profile
func (b *Bot) Quote(m *gateway.MessageCreateEvent, quote ...string) (string, error) {
	if quote == nil {
		return "", errors.New("no quote entered")
	}
	q := strings.Join(quote, " ")
	if len(q) > 1048 {
		return "", errors.New("quote too long. please submit a quote with a length lower than 1048 characters")
	}

	err := b.DB.UpdateUser(context.Background(), db.User{
		Quote:  q,
		UserID: int64(m.Author.ID),
	})
	if err != nil {
		log.Debug().
			Err(err).
			Str("Type", "QUOTE").
			Msg("Error setting quote")
		return "", errors.New("an error occured setting this as your profile quote, please retry later or raise an issue on https://github.com/Karitham/WaifuBot")
	}

	return fmt.Sprintf("New quote set :\n%s", q), nil
}
