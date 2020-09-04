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
		_, err := data.Message.Reply(ctx, session, err)
		log.Println("Error getting bot url: ", err)
		if err != nil {
			log.Println("Error sending error message: ", err)
		}
	}
	// Create embed
	_, err = client.CreateMessage(
		ctx,
		data.Message.ChannelID,
		&disgord.CreateMessageParams{
			Content :  "Invite link : <"+botURL+">",
		},
	)
	if err != nil {
		log.Println("There was an error sending invite message", err)
	}
}

func inviteHelp(data *disgord.MessageCreate) {
	_, err := client.CreateMessage(
		ctx,
		data.Message.ChannelID,
		&disgord.CreateMessageParams{
			Embed: &disgord.Embed{
				Title: "Invite Help || alias i",
				Description: fmt.Sprintf(
					"This is the help for the Invite functionality\n\n"+
						"This function can be used to get an invite link to be able to add the bot to your server\n"+
						"For this, simply use this command :" +
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
