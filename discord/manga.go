package discord

import (
	"fmt"
	"github.com/Karitham/WaifuBot/query"
	"github.com/andersfylling/disgord"
	"strings"
)

func searchManga(data *disgord.MessageCreate, args CmdArguments) {
	// check if there is a search term
	if len(args) > 0 {
		resp, queryErr := query.SearchManga(args.ParseArgToSearch().Name)
		if queryErr == nil {
			desc := fmt.Sprintf("\n%s...\n ", formatDescMangaSearch(resp.Media.Description))
			_, err := client.CreateMessage(
				ctx,
				data.Message.ChannelID,
				&disgord.CreateMessageParams{
					Content: resp.Media.SiteURL,
					Embed: &disgord.Embed{
						Title:       resp.Media.Title.Romaji,
						URL:         resp.Media.SiteURL,
						Description: desc,
						Color:       0x1663be,
						Thumbnail: &disgord.EmbedThumbnail{
							URL: resp.Media.CoverImage.Medium,
						},
						Footer: &disgord.EmbedFooter{
							IconURL: "https://anilist.co/img/icons/favicon-32x32.png",
							Text: fmt.Sprintf(
								"Score : %d%% | Status : %s",
								resp.Media.MeanScore,
								resp.Media.Status,
							),
						},
					},
				},
			)
			if err != nil {
				fmt.Println("There was an error when searching a manga: ", err)
			}
		} else {
			_, err := client.SendMsg(ctx, data.Message.ChannelID, queryErr)
			if err != nil {
				fmt.Println("There was an error trying to send an error message on manga search: ", err)
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

func searchMangaHelp(data *disgord.MessageCreate) {
	_, err := client.CreateMessage(
		ctx,
		data.Message.ChannelID,
		&disgord.CreateMessageParams{
			Embed: &disgord.Embed{
				Title: "Manga Search Help || alias sm",
				Description: fmt.Sprintf(
					"This is the help for search manga functionnality\n\n"+
						"You can search an manga by its name using the following syntax\n"+
						"`%ssearchManga Name`\n",
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
		fmt.Println("Error sending search manga help: ", err)
	}
}

func formatDescMangaSearch(inputDesc string) (desc string) {
	splitInput := strings.Split(inputDesc, " ")
	if len(splitInput) <= 40 {
		desc = strings.Join(splitInput, " ")
	} else {
		desc = strings.Join(splitInput[:40], " ")
	}
	return
}

