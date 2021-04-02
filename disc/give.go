package disc

import (
	"context"
	"fmt"

	"github.com/Karitham/WaifuBot/db"
	"github.com/diamondburned/arikawa/v2/bot/extras/arguments"
	"github.com/diamondburned/arikawa/v2/gateway"
	"github.com/lib/pq"
	"github.com/rs/zerolog/log"
)

// Give is used to give a character to a user
func (b *Bot) Give(m *gateway.MessageCreateEvent, ID int64, _ *arguments.UserMention) (string, error) {
	user := parseUser(m)

label:
	err := b.conn.GiveChar(context.Background(), db.GiveCharParams{
		ID:       ID,
		UserID:   int64(m.Author.ID),
		UserID_2: int64(user.ID),
	})
	if err, ok := err.(*pq.Error); ok && err.Code == "23503" {
		err := b.conn.CreateUser(b.Ctx.Context(), int64(user.ID))
		if err != nil {
			log.Err(err).
				Str("Type", "LIST").
				Int("ID", int(ID)).
				Int("Giver", int(m.Author.ID)).
				Int("Receiver", int(user.ID)).
				Msg("Could not create user")

			return "", err
		}
		goto label
	} else if err != nil {
		log.Debug().
			Err(err).
			Str("Type", "GIVE").
			Int("ID", int(ID)).
			Int("Giver", int(m.Author.ID)).
			Int("Receiver", int(user.ID)).
			Msg("Could not give")

		return "", err
	}

	log.Trace().
		Str("Type", "GIVE").
		Int("ID", int(ID)).
		Int("Giver", int(m.Author.ID)).
		Int("Receiver", int(user.ID)).
		Msg("Gave char")

	return fmt.Sprintf("You have given %d to %s", ID, user.Username), nil
}
