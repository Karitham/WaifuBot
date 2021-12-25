package discord

import (
	"database/sql"
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
		w.Respond(corde.NewResp().Content("An error occurred dialing the database, please try again later").Ephemeral().B())
		return
	}

	c, err := b.AnimeService.Random(IDs(chars))
	if err != nil {
		log.Err(err).Msg("error with anime service")
		w.Respond(corde.NewResp().Content("An error getting a random character occurred, please try again later").Ephemeral().B())
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
		w.Respond(corde.NewResp().Content("An error occurred dialing the database, please try again later").Ephemeral().B())
		return
	}

	w.Respond(corde.NewResp().Embeds(corde.NewEmbed().
		Title(c.Name.Full).
		URL(c.SiteURL).
		Color(anilist.Color).
		Footer(corde.Footer{IconURL: anilist.IconURL, Text: "View them on anilist"}).
		Thumbnail(corde.Image{URL: c.Image.Large}).
		Descriptionf("You rolled %s.\nCongratulations!", c.Name.Full).
		B(),
	).B())
}

func IDs(c []Character) []int {
	ids := make([]int, len(c))
	for i, v := range c {
		ids[i] = int(v.ID)
	}

	return ids
}
