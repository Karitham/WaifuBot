package disc

import (
	"bot/query"

	"github.com/andersfylling/disgord"
)

func search(data *disgord.MessageCreate, args []string) {
	// check if there is a search term
	if len(args) > 0 {
		resp, err := query.CharSearch(args)
		if err == nil {
			client.CreateMessage(
				ctx,
				data.Message.ChannelID,
				&disgord.CreateMessageParams{
					Embed: &disgord.Embed{
						Title: resp.Character.Name.Full,
						URL:   resp.Character.SiteURL,
						Color: 0x225577,
						Image: &disgord.EmbedImage{
							URL: resp.Character.Image.Large,
						},
					}})
		} else {
			client.SendMsg(ctx, data.Message.ChannelID, err)
		}
	} else {
		client.CreateMessage(
			ctx,
			data.Message.ChannelID,
			&disgord.CreateMessageParams{
				Embed: &disgord.Embed{Title: "Error, search requires at least 1 argument", Color: 0xcc0000}})
	}

}
