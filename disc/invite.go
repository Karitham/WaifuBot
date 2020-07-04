package disc

import "github.com/andersfylling/disgord"

func invite(data *disgord.MessageCreate) {
	botURL, err := client.InviteURL(ctx)
	if err != nil {
		data.Message.Reply(ctx, session, err)
	}
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
