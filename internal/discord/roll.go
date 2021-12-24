package discord

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/Karitham/WaifuBot/internal/anilist"
	"github.com/Karitham/corde"
	"github.com/rs/zerolog/log"
)

type randomer interface {
	Random(notIn []int) (anilist.Character, error)
}

type charStore interface {
	Put(userID corde.Snowflake, c Character) error
	Characters(userID corde.Snowflake) ([]Character, error)
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
	Favorite sql.NullInt64   `json:"favorite"`
	UserID   corde.Snowflake `json:"user_id"`
	ID       int32           `json:"id"`
}

func (b *Bot) Roller(w corde.ResponseWriter, i *corde.Interaction) {
	chars, err := b.Store.Characters(i.Member.User.ID)
	if err != nil {
		log.Err(err).Msg("error with db service")
		ephemeral(w, "An error occurred dialing the database, please try again later")
		return
	}

	c, err := b.AnimeService.Random(IDs(chars))
	if err != nil {
		log.Err(err).Msg("error with anime service")
		ephemeral(w, "An error getting a random character occurred, please try again later")
		return
	}

	if err := b.Store.Put(i.Member.User.ID, Character{
		Date:   time.Now(),
		Image:  c.Image.Large,
		Name:   c.Name.Full,
		Type:   "ROLL",
		UserID: i.Member.User.ID,
		ID:     int64(c.ID),
	}); err != nil {
		log.Err(err).Msg("error with db service")
		ephemeral(w, "An error occurred dialing the database, please try again later")
		return
	}

	w.WithSource(&corde.InteractionRespData{
		Embeds: []corde.Embed{{
			Title:       c.Name.Full,
			Description: fmt.Sprintf("You rolled %s.\nCongratulations!", c.Name.Full),
			URL:         c.SiteURL,
			Color:       anilist.Color,
			Footer: corde.Footer{
				IconURL: anilist.IconURL,
				Text:    "View them on anilist",
			},
			Thumbnail: corde.Image{
				URL: c.Image.Large,
			},
		}},
	},
	)
}

func ephemeral(w corde.ResponseWriter, message string) {
	w.WithSource(&corde.InteractionRespData{
		Content: message,
		Flags:   corde.RESPONSE_FLAGS_EPHEMERAL,
	})
}

func IDs(c []Character) []int {
	ids := make([]int, len(c))
	for i, v := range c {
		ids[i] = int(v.ID)
	}

	return ids
}
