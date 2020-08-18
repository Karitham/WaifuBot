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
	resp, err := query.CharSearchByPopularity(
		rand.New(
			rand.NewSource(
				time.Now().UnixNano(),
			),
		).Intn(conf.MaxCharDrop),
	)
	if err != nil {
		fmt.Println("Error getting random char : ", err)
		drop(data)
		return
	}
	enableClaim(resp)
	printDrop(data, resp.Page.Characters[0].Image.Large)
}

func printDrop(data *disgord.MessageCreate, image string) {
	// Sends the message
	_, err := client.CreateMessage(
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
					Text: fmt.Sprintf(
						"This characters initials are : %s",
						formatCharInitials(
							getCharInitials(char.Page.Characters[0].Name.Full),
						),
					),
				},
				Color: 0xF2FF2E,
			},
		},
	)
	if err != nil {
		fmt.Println("There was an error sending drop message: ", err)
	}
}

func formatCharInitials(initials []string) (formatted string) {
	for _, v := range initials {
		formatted = formatted + v + "."
	}
	return
}

// claim is used to claim a waifu and add it to your database
func claim(data *disgord.MessageCreate, args []string) {
	if len(args) > 0 && char.Page.Characters != nil {
		if strings.EqualFold(strings.Join(args, " "), char.Page.Characters[0].Name.Full) {
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
			avatar := getUserAvatar(data.Message.Author)

			// Create desc
			desc := fmt.Sprintf(
				"Well done %s, you claimed %s\n"+
					"It appears in :\n- %s",
				data.Message.Author.Username,
				char.Page.Characters[0].Name.Full,
				char.Page.Characters[0].Media.Nodes[0].Title.Romaji,
			)

			_, err := client.CreateMessage(
				ctx,
				data.Message.ChannelID,
				&disgord.CreateMessageParams{
					Embed: &disgord.Embed{
						Title:       "Claim successfull",
						URL:         char.Page.Characters[0].SiteURL,
						Description: desc,
						Thumbnail:   &disgord.EmbedThumbnail{URL: avatar},
						Image: &disgord.EmbedImage{
							URL: char.Page.Characters[0].Image.Large,
						},
						Timestamp: data.Message.Timestamp,
						Color:     0xFF924B,
					},
				},
			)
			if err != nil {
				fmt.Println("There was an error claiming character: ", err)
			}
			// Reset the char value
			char = query.CharStruct{}
		} else {
			resp, err := client.CreateMessage(
				ctx,
				data.Message.ChannelID,
				&disgord.CreateMessageParams{
					Embed: &disgord.Embed{
						Title:     "Claim unsucessfull",
						Timestamp: data.Message.Timestamp,
						Color:     0x626868,
					},
				},
			)
			if err != nil {
				fmt.Println("Create message returned error :", err)
			}
			go deleteMessage(resp, conf.DelWrongClaimAfter)
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
		go deleteMessage(resp, conf.DelWrongClaimAfter)
	}
}

func getCharInitials(name string) (initials []string) {
	for _, v := range strings.Split(strings.TrimSpace(name), " ") {
		initials = append(initials, strings.ToUpper(string(v[0])))
	}
	return
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
		fmt.Println("There was an error sending claim help message: ", err)
	}
}
