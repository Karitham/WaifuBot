package discord

import (
	"errors"
	"time"

	"github.com/Karitham/WaifuBot/internal/anilist"
	"github.com/Karitham/corde"
	"github.com/rs/zerolog/log"
)

type randomCharGetter interface {
	RandomChar(notIn ...int) (anilist.Character, error)
}

type Character struct {
	Date   time.Time       `json:"date"`
	Image  string          `json:"image"`
	Name   string          `json:"name"`
	Type   string          `json:"type"`
	UserID corde.Snowflake `json:"user_id"`
	ID     int64           `json:"id"`
}

type User struct {
	Date     time.Time       `json:"date"`
	Quote    string          `json:"quote"`
	Favorite uint64          `json:"favorite"`
	UserID   corde.Snowflake `json:"user_id"`
	Tokens   int32           `json:"tokens"`
}

func (b *Bot) roll(w corde.ResponseWriter, i *corde.InteractionRequest) {
	var char anilist.Character

	if err := b.Store.Tx(func(s Store) error {
		user, err := s.User(i.Context, i.Member.User.ID)
		if err != nil {
			return err
		}

		var toUpdate int = 0
		switch {
		case time.Now().After(user.Date.Add(b.RollTimeout)):
			toUpdate = 1 // Time
		case user.Tokens > b.TokensNeeded:
			toUpdate = 2 // Tokens
		default:
			w.Respond(corde.NewResp().
				Contentf("Invalid roll.\nYou need %d tokens to roll, you have %d, or you can wait %s until next free roll.",
					b.TokensNeeded,
					user.Tokens,
					time.Until(user.Date.Add(b.RollTimeout)).Round(time.Second),
				))
			return errors.New("not enough tokens")
		}

		chars, err := s.Chars(i.Context, i.Member.User.ID)
		if err != nil {
			log.Err(err).Msg("error with db service")
			w.Respond(corde.NewResp().Content("An error occurred dialing the database, please try again later").Ephemeral())
			return err
		}

		c, err := b.AnimeService.RandomChar(IDs(chars)...)
		if err != nil {
			log.Err(err).Msg("error with anime service")
			w.Respond(corde.NewResp().Content("An error getting a random character occurred, please try again later").Ephemeral())
			return err
		}
		char = c

		if err := s.PutChar(
			i.Context,
			i.Member.User.ID,
			Character{
				Date:   time.Now(),
				Image:  c.Image.Large,
				Name:   c.Name.Full,
				Type:   "ROLL",
				UserID: i.Member.User.ID,
				ID:     int64(c.ID),
			}); err != nil {
			log.Err(err).Msg("error with db service")
			w.Respond(corde.NewResp().Content("An error occurred dialing the database, please try again later").Ephemeral())
			return err
		}

		switch toUpdate {
		case 1:
			if err := s.SetUserDate(i.Context, i.Member.User.ID, time.Now()); err != nil {
				log.Err(err).Msg("error with db service")
				w.Respond(corde.NewResp().Content("An error occurred dialing the database, please try again later").Ephemeral())
				return err
			}
		case 2: // TODO: doesn't exist yet
			// if err := s.SetUserTokens(i.Context, i.Member.User.ID, user.Tokens-b.TokensNeeded); err != nil {
			// 	log.Err(err).Msg("error with db service")
			// 	w.Respond(corde.NewResp().Content("An error occurred dialing the database, please try again later").Ephemeral())
			// 	return err
			// }
		}

		return nil
	}); err != nil {
		return
	}

	w.Respond(corde.NewEmbed().
		Title(char.Name.Full).
		URL(char.SiteURL).
		Color(anilist.Color).
		Footer(corde.Footer{IconURL: anilist.IconURL, Text: "View them on anilist"}).
		Thumbnail(corde.Image{URL: char.Image.Large}).
		Descriptionf("You rolled %s.\nCongratulations!", char.Name.Full),
	)
}

func IDs(c []Character) []int {
	ids := make([]int, len(c))
	for i, v := range c {
		ids[i] = int(v.ID)
	}

	return ids
}
