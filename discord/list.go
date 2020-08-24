package discord

import (
	"fmt"
	"strconv"

	"github.com/Karitham/WaifuBot/database"

	"github.com/andersfylling/disgord"
)

func list(data *disgord.MessageCreate, args []string) {
	var page int
	var err error

	// check if there is a page input
	if len(args) > 0 {
		page, err = strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("There was an error parsing list", err)
		}
		if page < 0 {
			page = 0
		}
	}

	user := getUser(data)
	// Make the database query
	charList := database.ViewUserData(user.ID)

	// Send the first list
	_ = sendList(
		data,
		formatListEmbed(
			getUserAvatar(&user),
			(len(charList.Waifus)-1)/15,
			formatDescList(page, charList),
			&user,
		),
	)
}

func sendList(data *disgord.MessageCreate, embed *disgord.Embed) (msg *disgord.Message) {
	msg, err := client.CreateMessage(
		ctx,
		data.Message.ChannelID,
		&disgord.CreateMessageParams{
			Embed: embed,
		},
	)
	if err != nil {
		fmt.Println("There was an error sending list message : ", err)
	}

	return
}

// Format Embed
func formatListEmbed(avatar string, totalPages int, desc string, user *disgord.User) *disgord.Embed {
	return &disgord.Embed{
		Title:       fmt.Sprintf("%s's Waifu list", user.Username),
		Description: desc,
		Thumbnail: &disgord.EmbedThumbnail{
			URL: avatar,
		},
		Footer: &disgord.EmbedFooter{
			Text: fmt.Sprintf("Use list <page> to see a page. There are %d total pages.", totalPages),
		},
		Color: 0x88ffcc,
	}
}

func formatDescList(page int, charList database.UserDataStruct) (desc string) {
	// Check if the list is empty, if not, return a formatted description
	if len(charList.Waifus) >= 0 {
		for i := 15 * page; i < 15+15*page && i < len(charList.Waifus); i++ {
			desc += fmt.Sprintf(
				"`%d` : %s\n",
				charList.Waifus[i].ID,
				charList.Waifus[i].Name,
			)
		}
	} else {
		desc = "This user's list is empty"
	}
	return
}

func listHelp(data *disgord.MessageCreate) {
	_, err := client.CreateMessage(
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
		},
	)
	if err != nil {
		fmt.Println("There was an error sending list help message: ", err)
	}
}
