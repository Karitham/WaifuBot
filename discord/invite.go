package discord

import (
	"fmt"
	"log"

	"github.com/andersfylling/disgord"
)

// Send invite link in a small embed
func invite(data *disgord.MessageCreate) {

	// Get URL
	botURL, err := client.InviteURL(ctx)
	if err != nil {
		_, er := data.Message.Reply(ctx, session, err)
		log.Println("Error getting bot url: ", err)
		if er != nil {
			log.Println("Error sending error message: ", er)
		}
	}
	// Create embed
	_, er := client.CreateMessage(
		ctx,
		data.Message.ChannelID,
		&disgord.CreateMessageParams{
			Embed: &disgord.Embed{
				Title:     "Invite",
				URL:       botURL,
				Timestamp: data.Message.Timestamp,
				Color:     0x49b675,
			},
		},
	)
	if er != nil {
		fmt.Println("There was an error sending invite message")
	}
}

func inviteHelp(data *disgord.MessageCreate) {
	_, err := client.CreateMessage(
		ctx,
		data.Message.ChannelID,
		&disgord.CreateMessageParams{
			Embed: &disgord.Embed{
				Title: "Invite Help",
				Description: fmt.Sprintf(
					"This is the help for the Invite functionnality\n\n"+
						"Invite is used to get an invite link to be able to add the bot to your server, just use\n"+
						"`%sinvite`",
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
		log.Println("There was a problem sending invite help message: ", err)
	}
}
