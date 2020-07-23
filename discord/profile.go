package discord

import (
	"bot/database"
	"fmt"
	"time"

	"github.com/andersfylling/disgord"
)

func profile(data *disgord.MessageCreate) {
	var user disgord.User

	// If a user is mentionned, check the users profile instead
	if data.Message.Mentions != nil {
		user = *data.Message.Mentions[0]
	} else {
		user = *data.Message.Author
	}

	// Setrieve user information from database
	db := database.ViewUserData(user.ID)

	// Get avatar URL
	avatar, err := user.AvatarURL(64, false)
	if err != nil {
		fmt.Println(err)
	}

	// Send message
	client.CreateMessage(
		ctx,
		data.Message.ChannelID,
		&disgord.CreateMessageParams{
			Embed: &disgord.Embed{
				Title: data.Message.Author.Username,
				Thumbnail: &disgord.EmbedThumbnail{
					URL: avatar,
				},
				Description: desc(db),
				Image: &disgord.EmbedImage{
					URL: db.Favourite.Image,
				},
				Timestamp: data.Message.Timestamp,
				Color:     0xffe2fe,
			},
		},
	)
}

// Format Description
func desc(db database.OutputStruct) string {
	return fmt.Sprintf(
		`
		%s
		This user last rolled %s ago.
		Has rolled %d Waifus, and has claimed %d.
		%s`,
		quoteDesc(db.Quote),
		time.Now().Sub(db.Date).Truncate(time.Second),
		len(db.Waifus), db.ClaimedWaifus,
		favDesc(db.Favourite.Name),
	)
}

// Format favourite char
func favDesc(favChar string) string {
	if favChar == "" {
		return "This user has not set a favourite waifu yet"
	}
	return fmt.Sprintf("This user's favourite waifu is %s", favChar)
}

// Format quote
func quoteDesc(quote string) string {
	if quote == "" {
		return "This user has not set a custom quote yet"
	}
	return fmt.Sprintf(
		`Favourite quote is:
		"%s"`,
		quote,
	)
}

func profileHelp(data *disgord.MessageCreate) {
	client.CreateMessage(
		ctx,
		data.Message.ChannelID,
		&disgord.CreateMessageParams{
			Embed: &disgord.Embed{
				Title: "Profile Help || alias p",
				Description: fmt.Sprintf(
					"This is the help for the Profile functionnality\n\n"+
						"Profile displays the profile of the concerned user.\n"+
						"Use it like so :\n"+
						"`%sprofile <@User>`\n"+
						"(fields enclosed in <> are optionals)\n"+
						"You can tag a user to see his list too",
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
