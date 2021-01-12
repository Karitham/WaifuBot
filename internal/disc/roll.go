package disc

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/rs/zerolog/log"

	"github.com/Karitham/WaifuBot/internal/anilist"
	"github.com/Karitham/WaifuBot/internal/db"
	"github.com/diamondburned/arikawa/v2/discord"
	"github.com/diamondburned/arikawa/v2/gateway"
)

// Roll drops a random character and adds it to the database
func (b *Bot) Roll(m *gateway.MessageCreateEvent) (*discord.Embed, error) {
	t, err := b.conn.GetDate(b.Ctx.Context(), int64(m.Author.ID))
	if err == sql.ErrNoRows {
		err := b.conn.CreateUser(b.Ctx.Context(), int64(m.Author.ID))
		if err != nil {
			log.Err(err).
				Str("Type", "ROLL").
				Int("user", int(m.Author.ID)).
				Msg("Could not create user")

		}
	} else if err != nil {
		log.Err(err).
			Msg("Could not get date")

		return nil, err
	}

	if nextRollTime := time.Until(t.Add(b.conf.TimeBetweenRolls.Duration)); nextRollTime > 0 {
		return nil, fmt.Errorf("**illegal roll**,\nroll in %s", nextRollTime.Truncate(time.Second))
	}

	char, err := anilist.CharSearchByPopularity(b.seed.Uint64() % b.conf.MaxCharacterRoll)
	if err != nil {
		log.Err(err).
			Str("Type", "ROLL").
			Msg("Could not search char")

		return nil, err
	}

	err = b.conn.InsertChar(context.Background(), db.InsertCharParams{
		ID:     int64(char.Page.Characters[0].ID),
		UserID: int64(m.Author.ID),
		Image: sql.NullString{
			String: char.Page.Characters[0].Image.Large,
			Valid:  true,
		},
		Name: sql.NullString{
			String: char.Page.Characters[0].Name.Full,
			Valid:  true,
		},
	})
	if err != nil {
		log.Err(err).
			Str("Type", "ROLL").
			Msg("Could not insert char")

		return nil, err
	}

	err = b.conn.UpdateUserDate(context.Background(), db.UpdateUserDateParams{
		ID:   int64(m.Author.ID),
		Date: time.Now().UTC(),
	})
	if err != nil {
		log.Err(err).
			Str("Type", "ROLL").
			Msg("Could not update date")

		return nil, err
	}

	log.Trace().
		Str("Type", "ROLL").
		Int("User", int(m.Author.ID)).
		Uint("ID", char.Page.Characters[0].ID).
		Str("Name", char.Page.Characters[0].Name.Full).
		Msg("user rolled a waifu")

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
