package disc

import (
	"bot/database"
	"fmt"

	"github.com/andersfylling/disgord"
)

func profile(data *disgord.MessageCreate) {
	var user disgord.User

	// if a user is mentionned, check the users profile instead
	if data.Message.Mentions != nil {
		user = *data.Message.Mentions[0]
	} else {
		user = *data.Message.Author
	}

	// get avatar URL
	avatar, err := user.AvatarURL(128, false)
	if err != nil {
		client.CreateMessage(
			ctx,
			data.Message.ChannelID,
			&disgord.CreateMessageParams{Embed: &disgord.Embed{Title: fmt.Sprintf("Error, %s", err), Color: 0xcc0000}})
	}

	// retrieve user information from database
	db := database.SeeWaifus(user.ID)

	// send message
	client.CreateMessage(
		ctx,
		data.Message.ChannelID,
		&disgord.CreateMessageParams{
			Embed: &disgord.Embed{
				Title:     data.Message.Author.Username,
				Thumbnail: &disgord.EmbedThumbnail{URL: avatar},
				Description: fmt.Sprintf(`
				This user last rolled %s.
				He owns %d waifus.
				His favourite waifu is %d`, db.Date, len(db.Waifus), db.Waifus[0]),
				Color: 0xffe2fe,
			},
		})
}
