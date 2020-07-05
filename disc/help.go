package disc

import "github.com/andersfylling/disgord"

func help(data *disgord.MessageCreate) {
	client.CreateMessage(
		ctx,
		data.Message.ChannelID,
		&disgord.CreateMessageParams{
			Embed: &disgord.Embed{
				Title: "Help",
				Description: `
				roll (r) : Roll a new waifu
				list (l) : List the waifus you have
				search (s) : Search for a character by name / ID
				profile (s) : Display profile information for yourself, or the user mentioned
				invite : Invite link to add the bot to your server
				help (h) : Display this help page
				`,
				Color: 0xeec400,
			},
		})
}
