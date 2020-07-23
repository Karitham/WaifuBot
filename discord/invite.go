package discord

import (
	"fmt"

	"github.com/andersfylling/disgord"
)

// Send invite link in a small embed
func invite(data *disgord.MessageCreate) {

	// Get URL
	botURL, err := client.InviteURL(ctx)
	if err != nil {
		data.Message.Reply(ctx, session, err)
	}

	// Create embed
	client.CreateMessage(
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
}

func inviteHelp(data *disgord.MessageCreate) {
	client.CreateMessage(
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
}
