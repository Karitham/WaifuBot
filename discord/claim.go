package discord

import (
	"fmt"
	"log"
	"math/rand"
	"strings"
	"time"

	"github.com/Karitham/WaifuBot/database"
	"github.com/Karitham/WaifuBot/query"

	"github.com/andersfylling/disgord"
)

// CharStruct handles data for RandomChar function
type CharStruct struct {
	ID         int64
	SiteURL    string
	LargeImage string
	Name       string
	MediaTitle string
}

// char currently stored to be claimed
var char = make(map[disgord.Snowflake]CharStruct)

// enableClaim is used to make the new character claimable by the user
func enableClaim(data *disgord.MessageCreate, in query.CharStruct) {
	char[data.Message.ChannelID] = CharStruct{
		ID:         in.Page.Characters[0].ID,
		SiteURL:    in.Page.Characters[0].SiteURL,
		LargeImage: in.Page.Characters[0].Image.Large,
		Name:       strings.Join(strings.Fields(in.Page.Characters[0].Name.Full), " "),
		MediaTitle: in.Page.Characters[0].Media.Nodes[0].Title.Romaji,
	}
}

func drop(data *disgord.MessageCreate) {
	resp, err := query.CharSearchByPopularity(
		rand.New(
			rand.NewSource(
				time.Now().UnixNano(),
			),
		).Intn(conf.MaxCharacterDrop),
	)
	if err != nil {
		log.Println("Error getting random char : ", err)
		drop(data)
		return
	}
	enableClaim(data, resp)
	printDrop(data)
}

// Sends the message
func printDrop(data *disgord.MessageCreate) {
	_, err := client.CreateMessage(
		ctx,
		data.Message.ChannelID,
		&disgord.CreateMessageParams{
			Embed: &disgord.Embed{
				Title:       "A new character appeared !",
				Description: fmt.Sprintf("Can you guess who it is ?\nUse %sclaim to get this character for yourself", conf.Prefix),
				Thumbnail: &disgord.EmbedThumbnail{
					URL: char[data.Message.ChannelID].LargeImage,
				},
				Footer: &disgord.EmbedFooter{
					Text: fmt.Sprintf(
						"This characters initials are : %s",
						getCharInitials(char[data.Message.ChannelID].Name),
					),
				},
				Color: 0xF2FF2E,
			},
		},
	)
	if err != nil {
		log.Println("There was an error sending drop message: ", err)
	}
}

func getCharInitials(name string) (initials string) {
	for _, v := range strings.Fields(name) {
		initials = initials + strings.ToUpper(string(v[0])) + "."
	}
	return
}

// Claim is used to claim a waifu and add it to your database
func claim(data *disgord.MessageCreate, args []string) {
	if len(args) > 0 && char[data.Message.ChannelID].Name != "" {
		if strings.EqualFold(
			strings.Join(args, " "),
			char[data.Message.ChannelID].Name,
		) {
			// Add to db
			database.InputClaimChar{
				UserID: data.Message.Author.ID,
				CharList: database.CharLayout{
					ID:    char[data.Message.ChannelID].ID,
					Name:  char[data.Message.ChannelID].Name,
					Image: char[data.Message.ChannelID].LargeImage,
				},
			}.AddChar()

			// Increment the number of claimed waifus
			database.ClaimIncrementStruct{UserID: data.Message.Author.ID, Increment: 1}.ClaimIncrement()

			// Create the description
			desc := fmt.Sprintf(
				"Well done %s, you claimed %s\n"+
					"It appears in :\n- %s",
				data.Message.Author.Username,
				char[data.Message.ChannelID].Name,
				char[data.Message.ChannelID].MediaTitle,
			)

			_, err := client.CreateMessage(
				ctx,
				data.Message.ChannelID,
				&disgord.CreateMessageParams{
					Embed: &disgord.Embed{
						Title:       "Claim successful",
						URL:         char[data.Message.ChannelID].SiteURL,
						Description: desc,
						Thumbnail:   &disgord.EmbedThumbnail{URL: char[data.Message.ChannelID].LargeImage},
						Timestamp:   data.Message.Timestamp,
						Color:       0xFF924B,
					},
				},
			)
			if err != nil {
				log.Println("There was an error claiming character: ", err)
			}
			// Reset the char value
			char[data.Message.ChannelID] = CharStruct{}
		} else {
			resp, err := client.CreateMessage(
				ctx,
				data.Message.ChannelID,
				&disgord.CreateMessageParams{
					Embed: &disgord.Embed{
						Title:       "Claim unsuccessful",
						Description: "It isn't the good person ! Try again !",
						Timestamp:   data.Message.Timestamp,
						Color:       0x626868,
						Footer: &disgord.EmbedFooter{
							Text: fmt.Sprintf(
								"This characters initials are : %s",
								getCharInitials(char[data.Message.ChannelID].Name),
							),
						},
					},
				},
			)
			if err != nil {
				log.Println("Create message returned error :", err)
			}
			go deleteMessage(resp, conf.DeleteWrongClaimAfter.Duration)
		}
	} else {
		resp, err := client.CreateMessage(
			ctx,
			data.Message.ChannelID,
			&disgord.CreateMessageParams{
				Embed: &disgord.Embed{
					Title:       "Error",
					Description: "Please see\n`help claim`\nfor more information on the syntax",
					Timestamp:   data.Message.Timestamp,
					Color:       0xcc0000,
				},
			},
		)
		if err != nil {
			log.Println("Create message returned error :", err)
		}
		go deleteMessage(resp, conf.DeleteWrongClaimAfter.Duration)
	}
}

func claimHelp(data *disgord.MessageCreate) {
	_, err := client.CreateMessage(
		ctx,
		data.Message.ChannelID,
		&disgord.CreateMessageParams{
			Embed: &disgord.Embed{
				Title: "Claim Help || alias c",
				Description: fmt.Sprintf(
					"This is the help for the claim functionnality\n\n"+
						"You can claim a character for yourself after it has been dropped by using the following syntax :\n"+
						"`%sclaim Name`\n",
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
		log.Println("There was an error sending claim help message: ", err)
	}
}
