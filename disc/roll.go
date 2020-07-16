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
	// checkTimings verify if your query is legal
	ableToRoll := database.ViewUserData(data.Message.Author.ID).Date.Add(conf.TimeBetweenRolls * time.Hour)

	// verify if the roll is legal
	if ableToRoll.Sub(time.Now()) < 0 {
		// Makes the querry and adds the character to the database
		resp := queryRandom(data)

		// Create a descrption adapated to the character retrieved
		desc := fmt.Sprintf("You rolled character %d\nIt appears in :\n%s", resp.Page.Characters[0].ID, resp.Page.Characters[0].Media.Nodes[0].Title.Romaji)

		// Sends the message
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
	} else {
		client.CreateMessage(
			ctx,
			data.Message.ChannelID,
			&disgord.CreateMessageParams{
				Content: "cheh",
				Embed: &disgord.Embed{
					Title:       "Illegal roll",
					Description: fmt.Sprintf("You can roll in %s", ableToRoll.Sub(time.Now())),
					Color:       0xcc0000,
				}})
	}
}

// queryRandom makes a character query and adds it to the database
func queryRandom(data *disgord.MessageCreate) query.CharStruct {
	resp := query.RandomCharQuery(conf.MaxChar)
	database.AddChar(database.InputChar{
		UserID: data.Message.Author.ID,
		Date:   time.Now(),
		WaifuList: database.CharLayout{
			ID:    resp.Page.Characters[0].ID,
			Name:  resp.Page.Characters[0].Name.Full,
			Image: resp.Page.Characters[0].Image.Large,
		},
	})
	return resp
}
