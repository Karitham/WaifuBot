package disc

import (
	"bot/query"
	"fmt"
	"strconv"

	"github.com/andersfylling/disgord"
)

func trendingAnime(data *disgord.MessageCreate, args []string) {
	var desc string

	// check if there is a page input
	if len(args) > 0 {
		page, err := strconv.Atoi(args[0])
		if page > 1 {
			page--
		}
		if err != nil {
			fmt.Println(err)
		}
	}

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
	client.CreateMessage(
		ctx,
		data.Message.ChannelID,
		&disgord.CreateMessageParams{
			Embed: &disgord.Embed{
				Title:       fmt.Sprintf("Trending Anime List"),
				Description: desc,
				Color:       0x88ffcc,
			}})
}
