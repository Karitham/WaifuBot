package disc

import (
	"fmt"
	"log"

	"github.com/Karitham/WaifuBot/database"
	"github.com/diamondburned/arikawa/bot/extras/arguments"
	"github.com/diamondburned/arikawa/gateway"
)

// CharacterID represent a character CharacterID
type CharacterID uint

// Give is used to give a character to a user
func (b *Bot) Give(m *gateway.MessageCreateEvent, cID CharacterID, user *arguments.UserMention) (string, error) {
	changed, err := database.CharDelStruct{
		UserID: m.Author.ID,
		CharID: uint(cID),
	}.DelChar()
	if err != nil {
		return "", err
	}

	var char database.CharLayout
	for _, w := range changed.Waifus {
		if w.ID == uint(cID) {
			char = w
		}
	}

	err = database.InputChar{
		UserID:   user.ID(),
		CharList: char,
		Date:     changed.Date,
	}.AddChar()
	if err != nil {
		log.Println(err)
	}

	return fmt.Sprintf("You have given %s to %s", char.Name, user.Mention()), nil
}
