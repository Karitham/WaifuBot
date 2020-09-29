package discord

import (
	"fmt"
	"log"
	"strings"

	"github.com/andersfylling/disgord"
)

func help(data *disgord.MessageCreate, args []string) {
	if len(args) > 0 {
		switch strings.ToLower(args[0]) {
		case "search", "s":
			searchHelp(data)
		case "favourite", "favorite", "f":
			favouriteHelp(data)
		case "trendinganime", "animetrending", "ta":
			trendingAnimeHelp(data)
		case "trendingmanga", "mangatrending", "tm":
			trendingMangaHelp(data)
		case "searchanime", "animesearch", "sa":
			searchAnimeHelp(data)
		case "searchmanga", "mangasearch", "sm":
			searchMangaHelp(data)
		case "give", "g":
			giveCharHelp(data)
		case "quote", "q":
			quoteHelp(data)
		case "profile", "p":
			profileHelp(data)
		case "roll", "r":
			rollHelp(data)
		case "list", "l":
			listHelp(data)
		case "invite", "i":
			inviteHelp(data)
		case "claim", "c":
			claimHelp(data)
		case "verify", "v":
			verifyHelp(data)
		default:
			defaultHelp(data)
		}
	} else {
		defaultHelp(data)
	}
}

func defaultHelp(data *disgord.MessageCreate) {
	_, err := client.CreateMessage(
		ctx,
		data.Message.ChannelID,
		&disgord.CreateMessageParams{
			Embed: &disgord.Embed{
				Title: "Help Function || alias h",
				Description: fmt.Sprintf(
					"This is the help function.\n\n"+
						"Use `%shelp functionName` to find out more about each function\n"+
						"Current available functions : ```\nsearch, favourite, trendingAnime, trendingManga, searchAnime, searchManga, give, quote, profile, roll, list, invite, claim \n```"+
						"You can also read the source code here : https://github.com/Karitham/WaifuBot",
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
		log.Println("There was an error sending default help message: ", err)
	}
}
