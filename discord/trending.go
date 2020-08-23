package discord

import (
	"fmt"

	"github.com/Karitham/WaifuBot/query"

	"github.com/andersfylling/disgord"
)

func trendingAnime(data *disgord.MessageCreate, args []string) {
	var desc string

	// query the trending anime
	res, err := query.TrendingSearch()
	if err != nil {
		fmt.Println(err)
	}

	// return a formatted description
	for i := range res.Page.Media {
		desc += fmt.Sprintf("`%d` : %s\n", res.Page.Media[i].ID, res.Page.Media[i].Title.UserPreferred)
	}

	// Send the message
	_, er := client.CreateMessage(
		ctx,
		data.Message.ChannelID,
		&disgord.CreateMessageParams{
			Embed: &disgord.Embed{
				Title:       "Trending Anime List",
				Description: desc,
				Timestamp:   data.Message.Timestamp,
				Color:       0x88ffcc,
			},
		},
	)
	if er != nil {
		fmt.Println("There was an error sending trending anime message: ", er)
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
						"`%strendingAnimes",
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
		fmt.Println("There was an error sending trending anime help message: ", err)
	}
}
