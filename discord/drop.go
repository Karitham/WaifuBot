package discord

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
}

func printDrop(data *disgord.MessageCreate, image string) {
	// Sends the message
	client.CreateMessage(
		ctx,
		data.Message.ChannelID,
		&disgord.CreateMessageParams{
			Embed: &disgord.Embed{
				Title:       "A new character appeared",
				Description: fmt.Sprintf("use %sclaim to get this character for yourself", conf.Prefix),
				Image: &disgord.EmbedImage{
					URL: image,
				},
				Footer: &disgord.EmbedFooter{
					Text: fmt.Sprintf("Help requested by %s", data.Message.Author.Username),
				},
				Color: 0xF2FF2E,
			},
		},
	)
}
