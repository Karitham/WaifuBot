package disc

import (
	"bot/config"
	"bot/database"
	"bot/query"
	"fmt"
	"time"

	"github.com/andersfylling/disgord"
)

func roll(data *disgord.MessageCreate, config config.ConfJSONStruct) {
	resp := query.RandomCharQuery(config.MaxChar)
	database.AddWaifu(database.InputStruct{UserID: data.Message.Author.ID, Date: time.Now(), Waifu: resp.Page.Characters[0].ID})
	desc := fmt.Sprintf("You rolled waifu id : %d", resp.Page.Characters[0].ID)
	client.CreateMessage(
		ctx,
		data.Message.ChannelID,
		&disgord.CreateMessageParams{
			Embed: &disgord.Embed{
				Title:       resp.Page.Characters[0].Name.Full,
				URL:         resp.Page.Characters[0].SiteURL,
				Description: desc,
				Color:       0x225577,
				Image: &disgord.EmbedImage{
					URL: resp.Page.Characters[0].Image.Medium,
				},
			}})
}
