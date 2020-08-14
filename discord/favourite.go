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
		_, er := client.CreateMessage(
			ctx,
			data.Message.ChannelID,
			&disgord.CreateMessageParams{
				Embed: &disgord.Embed{
					Title:       "New favourite waifu set",
					Description: fmt.Sprintf("Your new favourite waifu is %s", resp.Character.Name.Full),
					Image: &disgord.EmbedImage{
						URL: resp.Character.Image.Large,
					},
					Timestamp: data.Message.Timestamp,
					Color:     0xffe2fe,
				},
			},
		)
		if er != nil {
			fmt.Println("There was an error sending favourite message: ", er)
		}
	} else {
		favouriteHelp(data)
	}
}

func favouriteHelp(data *disgord.MessageCreate) {
	_, err := client.CreateMessage(
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
				Footer: &disgord.EmbedFooter{
					Text: fmt.Sprintf("Help requested by %s", data.Message.Author.Username),
				},
				Timestamp: data.Message.Timestamp,
				Color:     0xeec400,
			},
		},
	)
	if err != nil {
		fmt.Println("There was an error sending favourite help message: ", err)
	}
}
