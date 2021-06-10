package disc

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/Karitham/WaifuBot/anilist"
	"github.com/Karitham/WaifuBot/db"
	"github.com/rs/zerolog/log"

	"github.com/diamondburned/arikawa/v2/discord"
	"github.com/diamondburned/arikawa/v2/gateway"
)

// Roll drops a random character and adds it to the database
func (b *Bot) Roll(m *gateway.MessageCreateEvent) (*discord.Embed, error) {
	p, err := b.DB.GetProfile(b.Ctx.Context(), int64(m.Author.ID))
	if err != nil {
		CreateErr := b.DB.CreateUser(b.Ctx.Context(), int64(m.Author.ID))
		if CreateErr != nil {
			log.Err(err).
				Str("Type", "ROLL").
				Int("user", int(m.Author.ID)).
				AnErr("GetProfileErr", err).
				Msg("Could not create user")
			return nil, errors.New("error creating your user profile, please raise an issue on <https://github.com/Karitham/WaifuBot>")
		}
	}

	if nextRollTime := time.Until(p.Date.UTC().Add(b.conf.TimeBetweenRolls)); nextRollTime > 0 {
		return nil, fmt.Errorf("**illegal roll**,\nroll in %s", nextRollTime.Truncate(time.Second))
	}

	list, err := b.DB.GetChars(b.Ctx.Context(), int64(m.Author.ID))
	if err != nil {
		log.Err(err).
			Str("Type", "ROLL").
			Msg("Could not get chars from DB")

		return nil, errors.New("error getting your current characters, please raise an issue on <https://github.com/Karitham/WaifuBot>")
	}

	notIn := func() []int64 {
		n := make([]int64, len(list))
		for _, i := range list {
			n = append(n, i.ID)
		}
		return n
	}()

	char, err := anilist.CharSearchByPopularity(b.seed.Uint64()%uint64(b.conf.MaxCharacterRoll), notIn)
	if err != nil {
		log.Err(err).
			Str("Type", "ROLL").
			Interface("List", list).
			Msg("Could not search char")

		return nil, errors.New("could not get a random character. Anilist is most likely down, please retry later")
	}

	err = b.DB.Tx(func(q db.Querier) error {
		err = q.InsertChar(
			b.Ctx.Context(),
			db.InsertCharParams{
				Image:  char.Page.Characters[0].Image.Large,
				Name:   strings.Join(strings.Fields(char.Page.Characters[0].Name.Full), " "),
				Type:   "ROLL",
				ID:     char.Page.Characters[0].ID,
				UserID: int64(m.Author.ID),
			},
		)
		if err != nil {
			return err
		}

		return q.UpdateUser(b.Ctx.Context(), db.User{Date: time.Now().UTC(), UserID: int64(m.Author.ID)})
	})

	if err != nil {
		log.Err(err).
			Str("Type", "ROLL").
			Msg("Could not add char to the database")

		return nil, errors.New("could not add this character to your list, you can retry. If this error occurs multiple times, please raise an issue on <https://github.com/Karitham/WaifuBot>")
	}

	return &discord.Embed{
		Title: char.Page.Characters[0].Name.Full,
		URL:   char.Page.Characters[0].SiteURL,
		Description: fmt.Sprintf(
			"You rolled character %d\nIt appears in :\n- %s",
			char.Page.Characters[0].ID, char.Page.Characters[0].Media.Nodes[0].Title.Romaji,
		),
		Thumbnail: &discord.EmbedThumbnail{
			URL: char.Page.Characters[0].Image.Large,
		},
		Footer: &discord.EmbedFooter{
			Text: fmt.Sprintf(
				"You can roll again in %s",
				b.conf.TimeBetweenRolls.Truncate(time.Second),
			),
		},
	}, nil
}
