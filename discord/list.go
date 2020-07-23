package discord

import (
	"bot/database"
	"fmt"
	"strconv"

	"github.com/andersfylling/disgord"
)

func list(data *disgord.MessageCreate, args []string) {
	var user disgord.User
	var desc string
	var page, i int
	var err error

	// check if there is a page input
	if len(args) > 0 {
		page, err = strconv.Atoi(args[0])
		if page > 1 {
			page--
		}
		if err != nil {
			fmt.Println(err)
		}
	}

	// If there is a mention, display the person's profile instead
	if data.Message.Mentions != nil {
		user = *data.Message.Mentions[0]
	} else {
		user = *data.Message.Author
	}

	// Make the database query
	WList := database.ViewUserData(user.ID)

	// Check if the list is empty, if not, return a formatted description
	if len(WList.Waifus) < 1 {
		desc = "This user's list is empty"
	} else {
		// Display the correct page
		for i = 15 * page; i < 15+15*page && i < len(WList.Waifus); i++ {
			desc += fmt.Sprintf("`%d` : %s\n", WList.Waifus[i].ID, WList.Waifus[i].Name)
		}

		// if there's a next page, tell the user it's possible to see it
		if i < len(WList.Waifus) {
			desc += fmt.Sprintf("\nYou can use %slist %d to see the next page", conf.Prefix, page+2)
		}
	}

	// get avatar URL
	avatar, err := user.AvatarURL(128, false)
	if err != nil {
		fmt.Println(err)
	}

	// Send the message
	client.CreateMessage(
		ctx,
		data.Message.ChannelID,
		&disgord.CreateMessageParams{
			Embed: &disgord.Embed{
				Title:       fmt.Sprintf("%s's Waifu list", user.Username),
				Description: desc,
				Thumbnail:   &disgord.EmbedThumbnail{URL: avatar},
				Color:       0x88ffcc,
			}})
}

func listHelp(data *disgord.MessageCreate) {
	client.CreateMessage(
		ctx,
		data.Message.ChannelID,
		&disgord.CreateMessageParams{
			Embed: &disgord.Embed{
				Title: "List Help || alias `l`",
				Description: fmt.Sprintf(
					"This is the help for the List functionnality\n\n"+
						"List is used to display a list of your owned waifus. It is displayed from oldest to newest.\n"+
						"Use the following syntax to display your list :\n"+
						"`%slist <page> <@User>`"+
						"(fields enclosed in <> are optionals)\n"+
						"You can tag a user to see his list too",
					conf.Prefix,
				),
				Footer: &disgord.EmbedFooter{
					Text: fmt.Sprintf("Help requested by %s", data.Message.Author.Username),
				},
				Color: 0xeec400,
			},
		})
}
