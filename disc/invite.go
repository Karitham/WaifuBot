package disc

import "github.com/andersfylling/disgord"

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
				Title: "Invite",
				URL:   botURL,
				Color: 0x49b675,
			},
		})
}
