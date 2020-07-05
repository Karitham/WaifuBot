package disc

import (
	"github.com/andersfylling/disgord"
)

func list(data *disgord.MessageCreate) {
	client.CreateMessage(
		ctx,
		data.Message.ChannelID,
		&disgord.CreateMessageParams{
			Embed: &disgord.Embed{
				Title:       "Waifu list",
				Description: "waifus()",
				Color:       0x88ffcc,
			}})
}
