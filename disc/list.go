package disc

import (
	db "bot/data"
	"fmt"

	"github.com/andersfylling/disgord"
)

func list(data *disgord.MessageCreate) {
	var desc string
	waifuList := db.SeeWaifus(data.Message.Author.ID).Waifu
	waifus := func() string {
		for _, v := range waifuList {
			desc = fmt.Sprintf("%d\n%s", v, desc)
		}
		return desc
	}
	client.CreateMessage(
		ctx,
		data.Message.ChannelID,
		&disgord.CreateMessageParams{
			Embed: &disgord.Embed{
				Title:       "Waifu list",
				Description: waifus(),
				Color:       0x88ffcc,
			}})
}
