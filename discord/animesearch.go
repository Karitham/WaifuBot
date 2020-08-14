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
		resp, queryErr := query.SearchAnime(args.ParseArgToSearch().Name)
		if queryErr == nil {
			desc := fmt.Sprintf(
				"I found the anime ID %d.\n "+
					"This anime is %s."+
					"Description : %s...\n ",
				resp.Media.ID,
				strings.ToLower(resp.Media.Status),
				formatDescAnimeSearch(resp.Media.Description),
			)
			_, err := client.CreateMessage(
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
			if err != nil {
				fmt.Println("There was an error when searching an anime: ", err)
			}
		} else {
			_, err := client.SendMsg(ctx, data.Message.ChannelID, queryErr)
			if err != nil {
				fmt.Println("there was an error sending error message on anime search: ", err)
			}
		}
	} else {
		_, err := client.CreateMessage(
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
		if err != nil {
			fmt.Println("there was an error sending error message on anime search: ", err)
		}
	}

}

func searchAnimeHelp(data *disgord.MessageCreate) {
	_, err := client.CreateMessage(
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
	if err != nil {
		fmt.Println("Error sending search anime help: ", err)
	}
}

func formatDescAnimeSearch(inputDesc string) (desc string) {
	splitInput := strings.Split(inputDesc, " ")
	if len(splitInput) <= 40 {
		desc = strings.Join(splitInput, " ")
	} else {
		desc = strings.Join(splitInput[:40], " ")
	}
	return
}
