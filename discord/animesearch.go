package discord

import (
	"bot/query"
	"fmt"
	"strings"

	"github.com/andersfylling/disgord"
)

func searchAnime(data *disgord.MessageCreate, args CmdArguments) {
	// check if there is a search term
	if len(args) > 0 {
		resp, err := query.SearchAnime(args.ParseArgToSearch().Name)
		if err == nil {
			desc := fmt.Sprintf(
				"I found the anime ID %d.\n "+
					"This anime is %s."+
					"Description : %s...\n ",
				resp.Media.ID,
				strings.ToLower(resp.Media.Status),
				formatDesc(resp.Media.Description),
			)
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
					},
				},
			)
		} else {
			client.SendMsg(ctx, data.Message.ChannelID, err)
		}
	} else {
		client.CreateMessage(
			ctx,
			data.Message.ChannelID,
			&disgord.CreateMessageParams{
				Embed: &disgord.Embed{
					Title:     "Error, search requires at least 1 argument",
					Timestamp: data.Message.Timestamp,
					Color:     0xcc0000,
				},
			},
		)
	}

}

func searchAnimeHelp(data *disgord.MessageCreate) {
	client.CreateMessage(
		ctx,
		data.Message.ChannelID,
		&disgord.CreateMessageParams{
			Embed: &disgord.Embed{
				Title: "Anime Search Help || alias sa",
				Description: fmt.Sprintf(
					"This is the help for search anime functionnality\n\n"+
						"You can search an anime by its name using the following syntax\n"+
						"`%ssearchAnime Name`\n",
					conf.Prefix,
				),
				Footer: &disgord.EmbedFooter{
					Text: fmt.Sprintf("Help requested by %s", data.Message.Author.Username),
				},
				Timestamp: data.Message.Timestamp,
				Color:     0xeec400,
			},
		},
	)
}

func formatDesc(inputDesc string) (desc string) {
	splitInput := strings.Split(inputDesc, " ")
	if len(splitInput) <= 40 {
		desc = strings.Join(splitInput, " ")
	} else {
		desc = strings.Join(splitInput[:40], " ")
	}
	return
}
