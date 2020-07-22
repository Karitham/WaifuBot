package discord

import (
	"bot/database"
	"bot/query"
	"fmt"

	"github.com/andersfylling/disgord"
)

func favourite(data *disgord.MessageCreate, args CmdArguments) {
	if len(args) > 0 {
		// Query Char using search to Anilists graphQL api
		resp, err := query.CharSearch(args.ParseArgToSearch())
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

func favouriteHelp(data *disgord.MessageCreate) {
	client.CreateMessage(
		ctx,
		data.Message.ChannelID,
		&disgord.CreateMessageParams{
			Embed: &disgord.Embed{
				Title: "Favourite Help || alias f ",
				Description: fmt.Sprintf(
					"This is the help for the Favourite functionnality\n\n"+
						"You can add a character as favourite by using the following syntax :\n"+
						"`%sfavourite Name/ID`\n"+
						"You can use Name OR ID\n",
					conf.Prefix,
				),
				Color: 0xeec400,
			},
		})
}
