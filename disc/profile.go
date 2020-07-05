package disc

import (
	db "bot/data"
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
	db := db.SeeWaifus(user.ID)

	// send message
	client.CreateMessage(
		ctx,
		data.Message.ChannelID,
		&disgord.CreateMessageParams{
			Embed: &disgord.Embed{
				Title:       data.Message.Author.Username,
				Thumbnail:   &disgord.EmbedThumbnail{URL: avatar},
				Description: fmt.Sprintf("This user last rolled %s.\nHe owns %d waifus.\nHis favourite waifu is", db.Date, len(db.Waifu)),
				Color:       0xffe2fe,
			},
		})
}
