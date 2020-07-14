package disc

import (
	"bot/query"
	"fmt"

	"github.com/andersfylling/disgord"
)

func animesearch(data *disgord.MessageCreate, args []string) {
	// check if there is a search term
	if len(args) > 0 {
		resp, err := query.CharSearch(args)
		if err == nil {
			desc := fmt.Sprintf("I found the anime ID : %d\nThe name of the anime is : %s\n", resp.Anime.ID, resp.Anime.Title.Romaji)
			client.CreateMessage(
				ctx,
				data.Message.ChannelID,
				&disgord.CreateMessageParams{
					Embed: &disgord.Embed{
						Title:       resp.Anime.Title.Romaji,
						URL:         resp.Anime.SiteURL,
						Description: desc,
						Color:       0x225577,
						Image: &disgord.EmbedImage{
							URL: resp.Anime.CoverImage.Large,
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
