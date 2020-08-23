package discord

import (
	"fmt"
	"strings"

	"github.com/Karitham/WaifuBot/database"

	"github.com/andersfylling/disgord"
)

func quote(data *disgord.MessageCreate, args CmdArguments) {
	if len(args) > 0 {
		// Transform args into a full string
		argStr := strings.Join(args, " ")

		// Set quote in database
		database.NewQuote{
			UserID: data.Message.Author.ID,
			Quote:  argStr,
		}.SetQuote()

		avatar := getUserAvatar(data.Message.Author)

		// Send confirmation message
		_, err := client.CreateMessage(
			ctx,
			data.Message.ChannelID,
			&disgord.CreateMessageParams{
				Embed: &disgord.Embed{
					Title:       "New quote set",
					Thumbnail:   &disgord.EmbedThumbnail{URL: avatar},
					Description: fmt.Sprintf("Your new quote is %s", argStr),
					Timestamp:   data.Message.Timestamp,
					Color:       0xffe2fe,
				},
			},
		)
		if err != nil {
			fmt.Println("There was an error sending quote message: ", err)
		}
	} else {
		_, err := client.CreateMessage(
			ctx,
			data.Message.ChannelID,
			&disgord.CreateMessageParams{
				Embed: &disgord.Embed{
					Title:     "Error, quote requires at least 1 argument",
					Timestamp: data.Message.Timestamp,
					Color:     0xcc0000,
				},
			},
		)
		if err != nil {
			fmt.Println("There was an error sending error profile message: ", err)
		}
	}
}

func quoteHelp(data *disgord.MessageCreate) {
	_, err := client.CreateMessage(
		ctx,
		data.Message.ChannelID,
		&disgord.CreateMessageParams{
			Embed: &disgord.Embed{
				Title: "Quote Help || alias q",
				Description: fmt.Sprintf(
					"This is the help for the Quote functionnality\n\n"+
						"You can add a favourite quote to be displayed on your profile by using the following syntax :\n"+
						"`%squote The thing you want to quote`\n",
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
		fmt.Println("There was an error sending quote help message: ", err)
	}
}
