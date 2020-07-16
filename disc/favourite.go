package disc

import (
	"bot/database"
	"fmt"

	"github.com/andersfylling/disgord"
)

func favourite(data *disgord.MessageCreate, args CmdArguments) {
	if len(args) > 0 {
		// Query Char using search to Anilists graphQL api
		resp, err := args.ParseArgToSearch().CharSearch()
		if err != nil {
			fmt.Println(err)
		}

		// Set favourite in database
		database.FavouriteStruct{
			UserID: data.Message.Author.ID,
			Favourite: database.CharLayout{
				ID:    resp.Character.ID,
				Name:  resp.Character.Name.Full,
				Image: resp.Character.Image.Large,
			},
		}.SetFavourite()

		// Send confirmation message
		client.CreateMessage(
			ctx,
			data.Message.ChannelID,
			&disgord.CreateMessageParams{
				Embed: &disgord.Embed{
					Title:       "New favourite waifu set",
					Description: fmt.Sprintf("Your new favourite waifu is %s", resp.Character.Name.Full),
					Color:       0xffe2fe,
					Image: &disgord.EmbedImage{
						URL: resp.Character.Image.Large,
					},
				}})
	} else {
		client.CreateMessage(
			ctx,
			data.Message.ChannelID,
			&disgord.CreateMessageParams{
				Embed: &disgord.Embed{Title: "Error, favourite requires at least 1 argument", Color: 0xcc0000}})
	}
}
