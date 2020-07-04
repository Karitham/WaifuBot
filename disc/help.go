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
				roll (r): 	Roll a new waifu
				list (l): 	List the waifus you have
				invite : 	Invite link to add the bot to your server
				help (h) :	show the commands you can use
				`,
				Color: 0xeec400,
			},
		})
}
