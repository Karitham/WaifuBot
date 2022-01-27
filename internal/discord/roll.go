package discord

import (
	"errors"
	"time"

	"github.com/Karitham/WaifuBot/internal/anilist"
	"github.com/Karitham/corde"
	"github.com/rs/zerolog/log"
)

type randomCharGetter interface {
	RandomChar(notIn ...int64) (anilist.CharAndMedia, error)
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
	var char anilist.CharAndMedia

	if err := b.Store.Tx(func(s Store) error {
		user, err := s.User(i.Context, i.Member.User.ID)
		if err != nil {
			return err
		}

		var toUpdate int = 0
		switch {
		case time.Now().After(user.Date.Add(b.RollTimeout)):
			toUpdate = 1 // Time
		case user.Tokens >= b.TokensNeeded:
			toUpdate = 2 // Tokens
		default:
			w.Respond(newErrf("Invalid roll.\nYou need %d tokens to roll, you have %d, or you can wait %s until next free roll.",
				b.TokensNeeded,
				user.Tokens,
				time.Until(user.Date.Add(b.RollTimeout)).Round(time.Second),
			))
			return errors.New("not enough tokens")
		}

		charsIDs, err := s.CharsIDs(i.Context, i.Member.User.ID)
		if err != nil {
			log.Err(err).Msg("error with db service")
			w.Respond(rspErr("An error occurred dialing the database, please try again later"))
			return err
		}

		c, err := b.AnimeService.RandomChar(charsIDs...)
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
				Image:  c.Image.Large,
				Name:   c.Name.Full,
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

	w.Respond(corde.NewEmbed().
		Title(char.Name.Full).
		URL(char.SiteURL).
		Footer(corde.Footer{IconURL: anilist.IconURL, Text: "View them on anilist"}).
		Thumbnail(corde.Image{URL: char.Image.Large}).
		Descriptionf("You rolled %s, id: %d\nIt appears in :\n- %s", char.Name.Full, char.ID, char.MediaTitle),
	)
}
