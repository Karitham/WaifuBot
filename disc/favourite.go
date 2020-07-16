package disc

import (
	"bot/database"
	"bot/query"
	"fmt"

	"github.com/andersfylling/disgord"
)

func favourite(data *disgord.MessageCreate, args []string) {
	if len(args) > 0 {
		// Query Char using search to Anilists graphQL api
		resp, err := query.CharSearch(ParseArgToSearch(args))
		if err != nil {
			fmt.Println(err)
		}

		// Set favourite in database
		database.SetFavourite(database.FavouriteStruct{
			UserID: data.Message.Author.ID,
			Favourite: database.CharLayout{
				ID:    resp.Character.ID,
				Name:  resp.Character.Name.Full,
				Image: resp.Character.Image.Large,
			},
		})

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
