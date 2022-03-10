package discord

import (
	"context"
	"errors"
	"time"

	"github.com/Karitham/corde"
	"github.com/Karitham/corde/components"
	"github.com/Karitham/corde/snowflake"
	"github.com/rs/zerolog/log"
)

type RandomCharer interface {
	RandomChar(ctx context.Context, notIn ...int64) (MediaCharacter, error)
}

type MediaCharacter struct {
	ID          int64
	Name        string
	ImageURL    string
	URL         string
	Description string
	MediaTitle  string
}

type Character struct {
	Date   time.Time           `json:"date"`
	Image  string              `json:"image"`
	Name   string              `json:"name"`
	Type   string              `json:"type"`
	UserID snowflake.Snowflake `json:"user_id"`
	ID     int64               `json:"id"`
}

type User struct {
	Date     time.Time           `json:"date"`
	Quote    string              `json:"quote"`
	Favorite uint64              `json:"favorite"`
	UserID   snowflake.Snowflake `json:"user_id"`
	Tokens   int32               `json:"tokens"`
}

func (b *Bot) roll(w corde.ResponseWriter, i *corde.Request[components.SlashCommandInteractionData]) {
	var char MediaCharacter

	if err := b.Store.Tx(func(s Store) error {
		user, err := s.User(i.Context, i.Member.User.ID)
		if err != nil {
			return err
		}

		toUpdate := 0
		switch {
		case time.Now().After(user.Date.Add(b.RollCooldown)):
			toUpdate = 1 // Time
		case user.Tokens >= b.TokensNeeded:
			toUpdate = 2 // Tokens
		default:
			w.Respond(newErrf("Invalid roll.\nYou need %d tokens to roll, you have %d, or you can wait %s until next free roll.",
				b.TokensNeeded,
				user.Tokens,
				time.Until(user.Date.Add(b.RollCooldown)).Round(time.Second),
			))
			return errors.New("not enough tokens")
		}

		charsIDs, err := s.CharsIDs(i.Context, i.Member.User.ID)
		if err != nil {
			log.Err(err).Msg("error with db service")
			w.Respond(rspErr("An error occurred dialing the database, please try again later"))
			return err
		}

		c, err := b.AnimeService.RandomChar(i.Context, charsIDs...)
		if err != nil {
			log.Err(err).Msg("error with anime service")
			w.Respond(rspErr("An error getting a random character occurred, please try again later"))
			return err
		}
		char = c

		if err := s.PutChar(
			i.Context,
			i.Member.User.ID,
			Character{
				Date:   time.Now(),
				Image:  c.ImageURL,
				Name:   c.Name,
				Type:   "ROLL",
				UserID: i.Member.User.ID,
				ID:     int64(c.ID),
			}); err != nil {
			log.Err(err).Msg("error with db service")
			w.Respond(rspErr("An error occurred dialing the database, please try again later"))
			return err
		}

		switch toUpdate {
		case 1:
			if err := s.SetUserDate(i.Context, i.Member.User.ID, time.Now()); err != nil {
				log.Err(err).Msg("error with db service")
				w.Respond(rspErr("An error occurred dialing the database, please try again later"))
				return err
			}
		case 2: // TODO: doesn't exist yet
			// if err := s.SetUserTokens(i.Context, i.Member.User.ID, user.Tokens-b.TokensNeeded); err != nil {
			// 	log.Err(err).Msg("error with db service")
			// 	w.Respond(newErr("An error occurred dialing the database, please try again later"))
			// 	return err
			// }
		}

		return nil
	}); err != nil {
		return
	}

	w.Respond(components.NewEmbed().
		Title(char.Name).
		URL(char.URL).
		Footer(components.Footer{IconURL: AnilistIconURL, Text: "View them on anilist"}).
		Thumbnail(components.Image{URL: char.ImageURL}).
		Descriptionf("You rolled %s, id: %d\nIt appears in :\n- %s", char.Name, char.ID, char.MediaTitle),
	)
}
