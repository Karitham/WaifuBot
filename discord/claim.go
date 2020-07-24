package discord

import (
	"bot/database"
	"bot/query"
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/andersfylling/disgord"
)

// char currently stored to be claimed
var char query.CharStruct

// enableClaim is used to make the new character claimable by the user
func enableClaim(in query.CharStruct) {
	char = in
}

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
				Timestamp: data.Message.Timestamp,
				Color:     0xF2FF2E,
			},
		},
	)
}

// claim is used to claim a waifu and add it to your database
func claim(data *disgord.MessageCreate, args []string) {
	if len(args) > 0 && char.Page.Characters != nil {
		if strings.ToLower(strings.Join(args, " ")) == strings.ToLower(char.Page.Characters[0].Name.Full) {

			// Add to db
			database.InputChar{
				UserID: data.Message.Author.ID,
				CharList: database.CharLayout{
					ID:    char.Page.Characters[0].ID,
					Name:  char.Page.Characters[0].Name.Full,
					Image: char.Page.Characters[0].Image.Large,
				},
			}.AddChar()

			// Increment claimed waifu
			database.ClaimIncrementStruct{UserID: data.Message.Author.ID, Increment: 1}.ClaimIncrement()
			// Send confirmation message
			avatar, err := data.Message.Author.AvatarURL(128, false)
			if err != nil {
				fmt.Println(err)
			}

			client.CreateMessage(
				ctx,
				data.Message.ChannelID,
				&disgord.CreateMessageParams{
					Embed: &disgord.Embed{
						Title:       "Claim successfull",
						URL:         char.Page.Characters[0].SiteURL,
						Description: fmt.Sprintf("Well done %s, you claimed %s", data.Message.Author.Username, char.Page.Characters[0].Name.Full),
						Thumbnail:   &disgord.EmbedThumbnail{URL: avatar},
						Image: &disgord.EmbedImage{
							URL: char.Page.Characters[0].Image.Large,
						},
						Timestamp: data.Message.Timestamp,
						Color:     0xFF924B,
					}})
			// Reset the char value
			char = query.CharStruct{}
		} else {
			resp, err := client.CreateMessage(
				ctx,
				data.Message.ChannelID,
				&disgord.CreateMessageParams{
					Embed: &disgord.Embed{
						Title:       "Claim unsucessfull",
						Description: fmt.Sprintf("Hint : this character's initial are %s", getCharInitials()),
						Timestamp:   data.Message.Timestamp,
						Color:       0x626868,
					},
				},
			)
			if err != nil {
				fmt.Println("Create message returned error :", err)
			}
			go deleteMessage(resp, conf.DelMessageAfter)
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
			fmt.Println("Create message returned error :", err)
		}
		go deleteMessage(resp, conf.DelMessageAfter)
	}
}

func getCharInitials() (initials string) {
	for _, v := range strings.Split(char.Page.Characters[0].Name.Full, " ") {
		initials = fmt.Sprintf("%s%s.", initials, strings.ToUpper(string(v[0])))
	}
	return
}

func claimHelp(data *disgord.MessageCreate) {
	client.CreateMessage(
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
}
