package disc

import (
	"bot/query"
	"fmt"
	"math/rand"
	"time"

	"github.com/andersfylling/disgord"
)

type toDrop query.CharStruct

func drop(data *disgord.MessageCreate) {
	query := query.CharSearchByPopularity(rand.New(rand.NewSource(time.Now().UnixNano())).Intn(conf.MaxChar))
	enableClaim(query)
	printDrop(data, query.Page.Characters[0].Image.Large)
	fmt.Println(query.Page.Characters[0].Name.Full)
}

func printDrop(data *disgord.MessageCreate, image string) {
	// Sends the message
	client.CreateMessage(
		ctx,
		data.Message.ChannelID,
		&disgord.CreateMessageParams{
			Embed: &disgord.Embed{
				Title:       "A new character appeared",
				Description: fmt.Sprintf("user %sclaim to claim this character for yourself", conf.Prefix),
				Color:       0xF2FF2E,
				Image: &disgord.EmbedImage{
					URL: image,
				},
			}})
}
