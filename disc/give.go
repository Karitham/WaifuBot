package disc

import (
	"bot/database"
	"fmt"
	"strconv"

	"github.com/andersfylling/disgord"
)

func giveChar(data *disgord.MessageCreate, arg []string) {
	charID, desc, valid := validGive(data, arg)

	if valid == true {
		database.DelWaifu(database.DelWaifuStruct{UserID: data.Message.Author.ID, WaifuID: charID})
	} else {
		// Get avatar
		avatar, err := data.Message.Author.AvatarURL(64, false)
		if err != nil {
			fmt.Println(err)
		}

		// Send message
		client.CreateMessage(
			ctx,
			data.Message.ChannelID,
			&disgord.CreateMessageParams{
				Embed: &disgord.Embed{
					Title:       "Give Waifu Failed",
					Thumbnail:   &disgord.EmbedThumbnail{URL: avatar},
					Description: desc,
					Color:       0xcc0000,
				},
			})
	}
}

// Verify if the user is able to give this character
func validGive(data *disgord.MessageCreate, arg []string) (int, string, bool) {
	if len(arg) > 0 {
		CharID, err := strconv.Atoi(arg[0])
		switch {
		case CharID == 0 || err != nil:
			return CharID, fmt.Sprintf("Error, %d is not a valid WaifuID,\nRefer to help to see this command's syntax", CharID), false
		case data.Message.Mentions == nil:
			return CharID, fmt.Sprintf("Error, please tag a discord user,\nRefer to help to see this command's syntax"), false
		case database.OwnsCharacter(data.Message.Mentions[0].ID, CharID) == false:
			return CharID, fmt.Sprintf("You do not own the character ID %d,\nPlease verify if the ID you entered is correct", CharID), false
		default:
			return CharID, fmt.Sprintf("You gave %d to %s", CharID, data.Message.Mentions[0].Username), true
		}
	}
	return 0, "", false
}
