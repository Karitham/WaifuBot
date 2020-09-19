package discord

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/Karitham/WaifuBot/database"
	"github.com/Karitham/WaifuBot/query"

	"github.com/andersfylling/disgord"
)

func roll(data *disgord.MessageCreate) {
	// checkTimings verifies if your query is legal
	ableToRoll := database.ViewUserData(data.Message.Author.ID).Date.Add(conf.TimeBetweenRolls.Duration)

	// Verify if the roll is legal
	if time.Until(ableToRoll) < 0 {
		// Makes the query and adds the character to the database
		resp := RandomToDB(data)

		// Create a description adapted to the retrieved character
		desc := fmt.Sprintf("You rolled character `%d`\nIt appears in :\n- %s", resp.Page.Characters[0].ID, resp.Page.Characters[0].Media.Nodes[0].Title.Romaji)

		// Sends the message
		_, err := client.CreateMessage(
			ctx,
			data.Message.ChannelID,
			&disgord.CreateMessageParams{
				Embed: &disgord.Embed{
					Title:       resp.Page.Characters[0].Name.Full,
					URL:         resp.Page.Characters[0].SiteURL,
					Description: desc,
					Thumbnail: &disgord.EmbedThumbnail{
						URL: resp.Page.Characters[0].Image.Large,
					},
					Timestamp: data.Message.Timestamp,
					Color:     0x57D4FF,
				},
			},
		)
		if err != nil {
			log.Println("There was an error sending roll message: ", err)
		}
	} else {
		illegalRoll(data, ableToRoll)
	}
}

// RandomToDB makes a character query and adds it to the database
func RandomToDB(data *disgord.MessageCreate) (resp query.CharStruct) {
	// Get response
	resp, err := query.CharSearchByPopularity(
		rand.New(
			rand.NewSource(
				time.Now().UnixNano(),
			),
		).Intn(conf.MaxCharacterRoll))
	if err != nil {
		log.Println("Error getting random char : ", err)
		return
	}

	database.InputChar{
		UserID: data.Message.Author.ID,
		Date:   time.Now(),
		CharList: database.CharLayout{
			ID:    resp.Page.Characters[0].ID,
			Name:  resp.Page.Characters[0].Name.Full,
			Image: resp.Page.Characters[0].Image.Large,
		},
	}.AddChar()
	return
}

func illegalRoll(data *disgord.MessageCreate, ableToRoll time.Time) {
	resp, err := client.CreateMessage(
		ctx,
		data.Message.ChannelID,
		&disgord.CreateMessageParams{
			Embed: &disgord.Embed{
				Title:       "Illegal roll",
				Description: fmt.Sprintf("You can roll in %s", time.Until(ableToRoll).Truncate(time.Second)),
				Timestamp:   data.Message.Timestamp,
				Color:       0xcc0000,
			},
		},
	)
	if err != nil {
		log.Println("Create message returned error :", err)
	}
	go deleteMessage(resp, conf.DeleteIllegalRollAfter.Duration)
}

func rollHelp(data *disgord.MessageCreate) {
	_, err := client.CreateMessage(
		ctx,
		data.Message.ChannelID,
		&disgord.CreateMessageParams{
			Embed: &disgord.Embed{
				Title: "Quote Help || alias r",
				Description: fmt.Sprintf(
					"This is the help for the Roll functionnality\n\n"+
						"You can roll random waifus using :\n"+
						"`%sroll`\n",
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
		log.Println("There was an error sending roll help message: ", err)
	}
}
