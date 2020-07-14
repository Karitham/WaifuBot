package disc

import (
	"bot/query"
	"fmt"
	"strconv"

	"github.com/andersfylling/disgord"
)

func animelist(data *disgord.MessageCreate, args []string) {
	var desc string
	var page int

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

	res, err := query.TrendingSearch(args)
	if err != nil {
		fmt.Println(err)
	}
	// Check if the list is empty, if not, return a formatted description
	for i := 10 * page; i < 10; i++ {
		desc += fmt.Sprintf("`%d` : %s\n", res.Media.ID, res.Media.Title.UserPreferred)
	}

	// Send the message
	client.CreateMessage(
		ctx,
		data.Message.ChannelID,
		&disgord.CreateMessageParams{
			Embed: &disgord.Embed{
				Title: fmt.Sprintf("Trending Anime List"),
				Color: 0x88ffcc,
			}})
}
