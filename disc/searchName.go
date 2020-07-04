package disc

import (
	"bot/query"
	"fmt"

	"github.com/andersfylling/disgord"
)

func searchName(data *disgord.MessageCreate, search string) {
	resp, err := query.CharByName(search)
	if err != nil {
		fmt.Println(err)
	}

	client.CreateMessage(
		ctx,
		data.Message.ChannelID,
		&disgord.CreateMessageParams{
			Embed: &disgord.Embed{
				Title:       resp.Character.Name.Full,
				URL:         resp.Character.SiteURL,
				Description: resp.Character.Description,
				Color:       0x225577,
				Image: &disgord.EmbedImage{
					URL: resp.Character.Image.Large,
				},
			}})
}
