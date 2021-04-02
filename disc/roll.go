package disc

import (
	"fmt"
	"strings"
	"time"

	"github.com/Karitham/WaifuBot/anilist"
	"github.com/rs/zerolog/log"

	"github.com/diamondburned/arikawa/v2/discord"
	"github.com/diamondburned/arikawa/v2/gateway"
)

// Roll drops a random character and adds it to the database
func (b *Bot) Roll(m *gateway.MessageCreateEvent) (*discord.Embed, error) {
	t, err := b.conn.GetDate(b.Ctx.Context(), int64(m.Author.ID))
	if err != nil {
		err := b.conn.CreateUser(b.Ctx.Context(), int64(m.Author.ID))
		if err != nil {
			log.Err(err).
				Str("Type", "ROLL").
				Int("user", int(m.Author.ID)).
				Msg("Could not create user")

		}
	}

	if nextRollTime := time.Until(t.UTC().Add(b.conf.TimeBetweenRolls.Duration)); nextRollTime > 0 {
		return nil, fmt.Errorf("**illegal roll**,\nroll in %s", nextRollTime.Truncate(time.Second))
	}

	list, err := b.conn.GetUserCharsIDs(b.Ctx.Context(), int64(m.Author.ID))
	if err != nil {
		log.Err(err).
			Str("Type", "ROLL").
			Msg("Could not get chars from DB")

		return nil, err
	}

	char, err := anilist.CharSearchByPopularity(b.seed.Uint64()%b.conf.MaxCharacterRoll, list)
	if err != nil {
		log.Err(err).
			Str("Type", "ROLL").
			Ints64("List", list).
			Msg("Could not search char")

		return nil, err
	}

	log.Trace().
		Str("Type", "ROLL").
		Int("User", int(m.Author.ID)).
		Uint("ID", char.Page.Characters[0].ID).
		Str("Name", char.Page.Characters[0].Name.Full).
		Msg("user rolled a waifu")

	err = b.conn.RollChar(
		b.Ctx.Context(),
		int64(m.Author.ID),
		int64(char.Page.Characters[0].ID),
		char.Page.Characters[0].Image.Large,
		strings.Join(
			strings.Fields(char.Page.Characters[0].Name.Full),
			" ",
		),
	)
	if err != nil {
		log.Err(err).
			Str("Type", "ROLL").
			Msg("Could not add char to the database")

		return nil, err
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
