package disc

import (
	"bot/database"
  "bot/query"
	"fmt"
	"strconv"

	"github.com/andersfylling/disgord"
)

func list(data *disgord.MessageCreate, args []string) {
	var desc string
	var page, i int
	var err error

	// check if there is a page input
	if len(args) > 0 {
		page, err = strconv.Atoi(args[0])
		if page > 1 {
			page--
		}
		if err != nil {
			fmt.Println(err)
		}
	}

	// Check if the list is empty, if not, return a formatted description
	for i = 10 * page; i < 10; i++ {
			desc += fmt.Sprintf("`%d` : %s\n", Page.Media[i].ID, Page.Media[i].UserPreferred)
		}

	// Send the message
	client.CreateMessage(
		ctx,
		data.Message.ChannelID,
		&disgord.CreateMessageParams{
			Embed: &disgord.Embed{
				Title:       fmt.Sprintf("Trending Anime List"),
				Color:       0x88ffcc,
			}})
}
