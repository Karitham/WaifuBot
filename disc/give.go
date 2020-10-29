package disc

import (
	"fmt"

	"github.com/Karitham/WaifuBot/database"
	"github.com/diamondburned/arikawa/bot/extras/arguments"
	"github.com/diamondburned/arikawa/gateway"
)

// Give is used to give a character to a user
func (b *Bot) Give(m *gateway.MessageCreateEvent, cID database.CharID, _ *arguments.UserMention) (string, error) {
	user := parseUser(m)

	if ok, _ := cID.VerifyWaifu(m.Author.ID); !ok {
		return "", fmt.Errorf("%s does not own character %d", m.Author.Username, cID)
	} else if ok, _ = cID.VerifyWaifu(m.Author.ID); ok {
		return "", fmt.Errorf("%s, you already own this waifu", m.Author.Username)
	}

	changed, err := cID.DelChar(m.Author.ID)
	if err != nil {
		return "", err
	}

	var char database.CharLayout
	for _, w := range changed.Waifus {
		if w.ID == uint(cID) {
			char = w
			break
		}
	}

	err = database.CharLayout(char).Add(user.ID)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("You have given %s to %s", char.Name, user.Username), nil
}
