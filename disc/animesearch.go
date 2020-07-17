package disc

import (
	"bot/query"
	"fmt"
	"strings"

	"github.com/andersfylling/disgord"
)

func searchAnime(data *disgord.MessageCreate, args CmdArguments) {
	// check if there is a search term
	if len(args) > 0 {
		resp, err := query.SearchAnime(strings.Join(args, " "))
		if err == nil {
			desc := fmt.Sprintf(
				"I found the anime ID %d.\n "+
					"Description : %s...\n "+
					"This anime is %s. \n"+
					"Number of episodes : %d.\n"+
					"Adult anime (hentai/ecchi) : %t",
				resp.Media.ID,
				strings.SplitN(resp.Media.Description, ".", 4)[0],
				strings.ToLower(resp.Media.Status),
				resp.Media.Episodes,
				resp.Media.IsAdult)
			client.CreateMessage(
				ctx,
				data.Message.ChannelID,
				&disgord.CreateMessageParams{
					Embed: &disgord.Embed{
						Title:       resp.Media.Title.Romaji,
						URL:         resp.Media.SiteURL,
						Description: desc,
						Color:       0x225577,
						Image: &disgord.EmbedImage{
							URL: resp.Media.CoverImage.Large,
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
