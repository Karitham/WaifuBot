package disc

import (
	"bot/database"
	"bot/query"
	"fmt"
	"strings"

	"github.com/andersfylling/disgord"
)

// char currently stored to be claimed
var char query.CharStruct

// enableClaim is used to make the new character claimable by the user
func enableClaim(in query.CharStruct) {
	char = in
}

// claim is used to claim a waifu and add it to your database
func claim(data *disgord.MessageCreate, args []string) {
	fmt.Println(char)
	if len(args) > 0 && char.Page.Characters[0].Name.Full != "" {
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
						Color:       0xFF924B,
						Image: &disgord.EmbedImage{
							URL: char.Page.Characters[0].Image.Large,
						},
					}})
			// Reset the char value
			char = query.CharStruct{}
			fmt.Println(char)
		} else {
			client.CreateMessage(
				ctx,
				data.Message.ChannelID,
				&disgord.CreateMessageParams{
					Embed: &disgord.Embed{
						Title:       "Claim unsucessfull",
						Description: fmt.Sprintf("Hint : this character's initial are %s", getCharInitials()),
						Color:       0x626868,
					}})
		}
	} else {
		client.CreateMessage(
			ctx,
			data.Message.ChannelID,
			&disgord.CreateMessageParams{
				Embed: &disgord.Embed{Title: "Error, enter the name to claim the character", Color: 0xcc0000}})
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
				Color: 0xeec400,
			},
		})
}
