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
	b.giveMu.Lock()
	defer b.giveMu.Unlock()

	if c, err := b.DB.GetChar(context.Background(), db.GetCharParams{
		ID:     ID,
		UserID: int64(m.Author.ID),
	}); err != nil || c.ID == 0 {
		return "", errors.New("you do not own this character")
	}

	for {
		char, err := b.DB.GiveChar(context.Background(), db.GiveCharParams{
			ID:    ID,
			Giver: int64(m.Author.ID),
			Given: int64(user.ID),
		})

		if err == nil {
			return fmt.Sprintf("You have given %s to %s", char.Name, user.Username), nil
		} else if err != nil {
			err, ok := err.(*pq.Error)
			if !ok {
				log.Err(err).
					Str("Type", "GIVE").
					Int("giver", int(m.Author.ID)).
					Int("given", int(user.ID)).
					Int("char", int(char.ID)).
					Msg("unknown error")
				return "", errors.New("unknown error")
			}

			switch err.Code {
			case "23503":
				err := b.DB.CreateUser(context.Background(), int64(user.ID))
				if err != nil {
					log.Err(err).
						Str("Type", "GIVE").
						Int("giver", int(m.Author.ID)).
						Int("given", int(user.ID)).
						Msg("unknown error")
					return "", fmt.Errorf("unknown error")
				}
			case "23505":
				return "", fmt.Errorf("%s already owns this character", user.Username)
			}
		}
	}
}
