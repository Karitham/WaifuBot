package discord

import (
	"fmt"
	"log"
	"time"

	"github.com/Karitham/WaifuBot/database"

	"github.com/andersfylling/disgord"
)

func list(data *disgord.MessageCreate, args []string) {
	page := CmdArguments(args).ParseArgToSearch().ID
	if page < 0 {
		page = 0
	}
	user := getUser(data)
	charList := database.ViewUserData(user.ID)

	// If a list has been sent not too long ago, replace said list
	val, ok := ListCache[user.ID]
	if ok && time.Since(val.Timestamp.Time.Add(conf.ListMaxUpdateTime)) <= 0 {
		deleteMessage(data.Message, 0)
		msg := setEmbedTo(
			embedUpdate{
				ChannelID: val.ChannelID,
				MessageID: val.ID,
				Embed: *formatListEmbed(
					getUserAvatar(&user),
					(len(charList.Waifus)-1)/15,
					formatDescList(page, charList),
					&user,
				),
			},
		)
		ListCache[user.ID] = msg
		return
	}
	// Send List
	msg := sendList(
		data,
		formatListEmbed(
			getUserAvatar(&user),
			(len(charList.Waifus)-1)/15,
			formatDescList(page, charList),
			&user,
		),
	)
	ListCache[user.ID] = msg
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
		log.Println("There was an error sending list message : ", err)
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
		log.Println("There was an error sending list help message: ", err)
	}
}
