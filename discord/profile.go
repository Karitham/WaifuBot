package discord

import (
	"fmt"
	"time"

	"github.com/Karitham/WaifuBot/database"

	"github.com/andersfylling/disgord"
)

func profile(data *disgord.MessageCreate) {
	var user disgord.User
	var avatar string
	var name string

	// If a user is mentioned, check the users profile instead
	if data.Message.Mentions != nil {
		user = *data.Message.Mentions[0]
		avatar = getUserAvatar(data.Message.Mentions[0])
		name = data.Message.Mentions[0].Username
	} else {
		user = *data.Message.Author
		avatar = getUserAvatar(data.Message.Author)
		name = user.Username
	}

	// Retrieve user information from database
	db := database.ViewUserData(user.ID)

	// Send message
	_, err := client.CreateMessage(
		ctx,
		data.Message.ChannelID,
		&disgord.CreateMessageParams{
			Embed: &disgord.Embed{
				Title: name,
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
	if err != nil {
		fmt.Println("There was an error sending profile message: ", err)
	}
}

// Format description
func desc(db database.UserDataStruct) string {
	return fmt.Sprintf(
		`
		%s
		This user last rolled %s ago.
		Has rolled %d Waifus, and has claimed %d.
		%s`,
		quoteDesc(db.Quote),
		time.Since(db.Date).Truncate(time.Second),
		(len(db.Waifus))-db.ClaimedWaifus, db.ClaimedWaifus,
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

// Help function for Profile
func profileHelp(data *disgord.MessageCreate) {
	_, err := client.CreateMessage(
		ctx,
		data.Message.ChannelID,
		&disgord.CreateMessageParams{
			Embed: &disgord.Embed{
				Title: "Profile Help || alias p",
				Description: fmt.Sprintf(
					"This is the help section for the Profile function\n\n"+
						"The Profile option displays the profile of the concerned user.\n"+
						"Use it like so :\n"+
						"`%sprofile <@User>`\n"+
						"(fields enclosed in <> are optionals)\n"+
						"You can tag a user to see his profile.",
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
		fmt.Println("There was an error sending profile help message: ", err)
	}
}
