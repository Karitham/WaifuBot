package disc

import (
	"fmt"
	"log"

	"github.com/Karitham/WaifuBot/database"
	"github.com/diamondburned/arikawa/bot/extras/arguments"
	"github.com/diamondburned/arikawa/gateway"
)

// CharacterID represent a character CharacterID
type CharacterID int

// Give is used to give a character to a user
func (b *Bot) Give(m *gateway.MessageCreateEvent, cID CharacterID, user *arguments.UserMention) (string, error) {
	changed, err := database.CharDelStruct{
		UserID: user.ID(),
		CharID: int(cID),
	}.DelChar()
	if err != nil {
		return "", err
	}

	var char database.CharLayout
	for _, v := range changed.Waifus {
		if v.ID == int64(cID) {
			char = v
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
