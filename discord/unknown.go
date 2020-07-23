package discord

import (
	"fmt"

	"github.com/andersfylling/disgord"
)

func unknown(data *disgord.MessageCreate) {
	client.CreateMessage(
		ctx,
		data.Message.ChannelID,
		&disgord.CreateMessageParams{
			Embed: &disgord.Embed{
				Title:       "Unknown command",
				Description: fmt.Sprintf("Type %shelp to see the commands available", conf.Prefix),
				Footer: &disgord.EmbedFooter{
					Text: fmt.Sprintf("You something wrong !!!"),
				},
				Color: 0xcc0000,
			},
		})
}
