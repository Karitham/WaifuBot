package discord

import (
	"fmt"

	"github.com/Karitham/WaifuBot/query"

	"github.com/andersfylling/disgord"
)

func trendingMedia(data *disgord.MessageCreate, format string, args []string) {
	var desc string
	var formattedType string = "Anime "
	if format != "ANIME" {
		formattedType = "Manga "
	}

	// query the trending anime
	res, err := query.TrendingMediaQuery(format)
	if err != nil {
		fmt.Println(err)
	}

	// return a formatted description
	for i := range res.Page.Media {
		desc += fmt.Sprintf("%d - %s\n", i+1, res.Page.Media[i].Title.Romaji)
	}

	// Send the message
	_, er := client.CreateMessage(
		ctx,
		data.Message.ChannelID,
		&disgord.CreateMessageParams{
			Embed: &disgord.Embed{
				Title:       formattedType + "Currently Trending",
				Description: desc,
				Color:       0x0e6b0e,
				Footer: &disgord.EmbedFooter{
					IconURL: "https://anilist.co/img/icons/favicon-32x32.png",
					Text:    "Trending " + formattedType + "list created from anilist",
				},
			},
		},
	)
	if er != nil {
		fmt.Println("There was an error sending trending media message: ", er)
	}
}

func trendingAnimeHelp(data *disgord.MessageCreate) {
	_, err := client.CreateMessage(
		ctx,
		data.Message.ChannelID,
		&disgord.CreateMessageParams{
			Embed: &disgord.Embed{
				Title: "Trending Anime Help || alias ta",
				Description: fmt.Sprintf(
					"This is the help for the Trending Anime functionnality\n\n"+
						"Trending anime is used to display the top 10 trending anime from anilist. Use it like so :\n"+
						"`%strendingAnime",
					conf.Prefix,
				),
				Footer: &disgord.EmbedFooter{
					Text: fmt.Sprintf("Help requested by %s", data.Message.Author.Username),
				},
				Color: 0xeec400,
			},
		},
	)
	if err != nil {
		fmt.Println("There was an error sending trending anime help message: ", err)
	}
}

func trendingMangaHelp(data *disgord.MessageCreate) {
	_, err := client.CreateMessage(
		ctx,
		data.Message.ChannelID,
		&disgord.CreateMessageParams{
			Embed: &disgord.Embed{
				Title: "Trending Manga Help || alias tm",
				Description: fmt.Sprintf(
					"This is the help for the Trending Manga functionnality\n\n"+
						"Trending manga is used to display the top 10 trending manga from anilist. Use it like so :\n"+
						"`%strendingManga",
					conf.Prefix,
				),
				Footer: &disgord.EmbedFooter{
					Text: fmt.Sprintf("Help requested by %s", data.Message.Author.Username),
				},
				Color: 0xeec400,
			},
		},
	)
	if err != nil {
		fmt.Println("There was an error sending trending manga help message: ", err)
	}
}
