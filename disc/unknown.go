package disc

import (
	"bot/config"
	"fmt"

	"github.com/andersfylling/disgord"
)

func unknown(data *disgord.MessageCreate, config config.ConfJSONStruct) {
	client.CreateMessage(
		ctx,
		data.Message.ChannelID,
		&disgord.CreateMessageParams{
			Embed: &disgord.Embed{
				Title:       "Unknown command",
				Description: fmt.Sprintf("Type %shelp to see the commands available", config.Prefix),
				Color:       0xcc0000,
			},
		})
}
