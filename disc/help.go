package disc

import (
	"fmt"
	"strings"

	"github.com/andersfylling/disgord"
)

func help(data *disgord.MessageCreate, args []string) {
	if len(args) > 0 {
		switch arg := strings.ToLower(args[0]); {
		case arg == "search" || arg == "s":
			searchHelp(data)
		case arg == "favourite" || arg == "favorite" || arg == "f":
			favouriteHelp(data)
		case arg == "trendinganimes" || arg == "ta":
			trendingAnimeHelp(data)
		case arg == "searchanime" || arg == "sa":
			searchAnimeHelp(data)
		case arg == "give" || arg == "g":
			giveCharHelp(data)
		case arg == "quote" || arg == "q":
			quoteHelp(data)
		case arg == "profile" || arg == "p":
			profileHelp(data)
		case arg == "roll" || arg == "r":
			rollHelp(data)
		case arg == "list" || arg == "l":
			listHelp(data)
		case arg == "invite":
			inviteHelp(data)
		case arg == "claim" || arg == "c":
			claimHelp(data)
		default:
			defaultHelp(data)
		}
	} else {
		defaultHelp(data)
	}
}

func defaultHelp(data *disgord.MessageCreate) {
	client.CreateMessage(
		ctx,
		data.Message.ChannelID,
		&disgord.CreateMessageParams{
			Embed: &disgord.Embed{
				Title: "Help || alias h",
				Description: fmt.Sprintf(
					"This is the help function.\n\n"+
						"Use `%shelp functionName` to find out more about each function\n"+
						"Current available functions : ```\nsearch, favourite, trendingAnime, searchAnime, give, quote, profile, roll, list, invite, claim \n```"+
						"You can also read the source code here : https://github.com/Karitham/WaifuBot",
					conf.Prefix,
				),
				Color: 0xeec400,
			},
		})
}

// 					Syntax : list <Optional page> <@Optional User>
// 					Syntax : give <ID> <@User>
// 					Syntax : search <ID / Name>
// 					Syntax : profile  <@Optional User>
// 					Syntax : favourite <ID / Name>
// 					Syntax : quote <text>
// 					Syntax : searchAnime <ID / Name>
