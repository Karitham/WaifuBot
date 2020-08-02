package discord

import (
	"fmt"

	"github.com/andersfylling/disgord"
)

func unknown(data *disgord.MessageCreate) {
	resp, err := client.CreateMessage(
		ctx,
		data.Message.ChannelID,
		&disgord.CreateMessageParams{
			Embed: &disgord.Embed{
				Title:       "Unknown command",
				Description: fmt.Sprintf("Type %shelp to see the commands available", conf.Prefix),
				Timestamp:   data.Message.Timestamp,
				Color:       0xcc0000,
			},
		},
	)
	if err != nil {
		fmt.Println("error while creating message :", err)
	}
	go deleteMessage(resp, conf.DelIllegalRollAfter)
}
