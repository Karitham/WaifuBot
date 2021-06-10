package disc

import (
	"context"
	"errors"
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
	err := b.DB.GiveChar(context.Background(), db.GiveCharParams{
		ID:    ID,
		Giver: int64(m.Author.ID),
		Given: int64(user.ID),
	})
	if err, ok := err.(*pq.Error); ok && err.Code == "23503" {
		err := b.DB.CreateUser(b.Ctx.Context(), int64(user.ID))
		if err != nil {
			log.Err(err).
				Str("Type", "LIST").
				Int("ID", int(ID)).
				Int("Giver", int(m.Author.ID)).
				Int("Receiver", int(user.ID)).
				Msg("Could not create user")

			return "", errors.New("there was an error giving the character. Please retry later or raise an issue on <https://github.com/Karitham/WaifuBot>")
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

		return "", errors.New("you already own this character. Give failed")
	}

	return fmt.Sprintf("You have given %d to %s", ID, user.Username), nil
}
