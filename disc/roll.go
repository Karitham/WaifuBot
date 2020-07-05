package disc

import (
	"bot/database"
	"bot/query"
	"fmt"
	"time"

	"github.com/andersfylling/disgord"
)

// WaifuRolled is used to input the waifu rolled into the database
type WaifuRolled struct {
	ID    int64  `bson:"ID"`
	Name  string `bson:"Name"`
	Image string `bson:"Image"`
}

func roll(data *disgord.MessageCreate) {

	// Query the character and add it to the database
	resp := query.RandomCharQuery(conf.MaxChar)
	database.AddWaifu(database.InputWaifu{
		UserID: data.Message.Author.ID,
		Date:   time.Now(),
		WaifuList: struct {
			ID    int64  "bson:\"ID\""
			Name  string "bson:\"Name\""
			Image string "bson:\"Image\""
		}{
			ID:    resp.Page.Characters[0].ID,
			Name:  resp.Page.Characters[0].Name.Full,
			Image: resp.Page.Characters[0].Image.Large,
		},
	})

	// Create a descrption adapated to the character retrieved
	desc := fmt.Sprintf("You rolled waifu %d", resp.Page.Characters[0].ID)

	// Send a message
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
					URL: resp.Page.Characters[0].Image.Large,
				},
			}})
}
