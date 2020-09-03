package discord

import (
	"fmt"
	"strings"

	"github.com/Karitham/WaifuBot/query"

	"github.com/andersfylling/disgord"
)

func search(data *disgord.MessageCreate, args CmdArguments) {
	// check if there is a search term
	if len(args) > 0 {
		resp, err := query.CharSearch(args.ParseArgToSearch())
		if err == nil {
			_, err := client.CreateMessage(
				ctx,
				data.Message.ChannelID,
				&disgord.CreateMessageParams{
					Content: resp.Character.SiteURL,
					Embed: &disgord.Embed{
						Title:       resp.Character.Name.Full,
						URL:         resp.Character.SiteURL,
						Description: fmt.Sprintf("I found character ID: `%d`\nThis character appears in :\n- %s", resp.Character.ID, resp.Character.Media.Nodes[0].Title.Romaji),
						Thumbnail: &disgord.EmbedThumbnail{
							URL: resp.Character.Image.Large,
						},
						Color: 0x1663be,
					},
				},
			)
			if err != nil {
				fmt.Println("There was an error sending search message: ", err)
			}
		} else {
			resp, err := client.SendMsg(ctx, data.Message.ChannelID, err)
			if err != nil {
				fmt.Println("Create message returned error :", err)
			}
			go deleteMessage(resp, conf.DelIllegalRollAfter)
		}
	} else {
		searchHelp(data)
	}
}

func searchHelp(data *disgord.MessageCreate) {
	_, err := client.CreateMessage(
		ctx,
		data.Message.ChannelID,
		&disgord.CreateMessageParams{
			Embed: &disgord.Embed{
				Title: "Character Search Help || alias s",
				Description: fmt.Sprintf(
					"This is the help for search character functionality\n\n"+
						"You can search a character using :\n"+
						"`%ssearch Name/ID`\n"+
						"You can search by either its name or its Anilist ID",
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
		fmt.Println("There was an error sending search help message: ", err)
	}
}

func searchMedia(data *disgord.MessageCreate, format string, args CmdArguments) {
	// check if there is a search term
	if len(args) > 0 {
		resp, queryErr := query.SearchMedia(args.ParseArgToSearch().Name, format)
		var formattedAdultString string = "❌"
		if resp.Media.IsAdult != false {
			formattedAdultString = "✔️"
		}
		if queryErr == nil {
			desc := fmt.Sprintf("\n%s...\n ", formatDescMediaSearch(resp.Media.Description))
			_, err := client.CreateMessage(
				ctx,
				data.Message.ChannelID,
				&disgord.CreateMessageParams{
					Content: resp.Media.SiteURL,
					Embed: &disgord.Embed{
						Title:       resp.Media.Title.Romaji,
						URL:         resp.Media.SiteURL,
						Description: desc,
						Color:       0x1663be,
						Thumbnail: &disgord.EmbedThumbnail{
							URL: resp.Media.CoverImage.Medium,
						},
						Footer: &disgord.EmbedFooter{
							IconURL: "https://anilist.co/img/icons/favicon-32x32.png",
							Text: fmt.Sprintf(
								"Score : %d%% | Status : %s | Adults only : %s",
								resp.Media.MeanScore,
								resp.Media.Status,
								formattedAdultString,
							),
						},
					},
				},
			)
			if err != nil {
				fmt.Println("There was an error when sending search anime message: ", err)
			}
		} else {
			_, err := client.SendMsg(ctx, data.Message.ChannelID, queryErr)
			if err != nil {
				fmt.Println("There was an error sending error message on anime search: ", err)
			}
		}
	} else {
		_, err := client.CreateMessage(
			ctx,
			data.Message.ChannelID,
			&disgord.CreateMessageParams{
				Embed: &disgord.Embed{
					Title:     "Error, search requires at least 1 argument",
					Timestamp: data.Message.Timestamp,
					Color:     0xcc0000,
				},
			},
		)
		if err != nil {
			fmt.Println("There was an error sending error message on anime search: ", err)
		}
	}

}

func searchAnimeHelp(data *disgord.MessageCreate) {
	_, err := client.CreateMessage(
		ctx,
		data.Message.ChannelID,
		&disgord.CreateMessageParams{
			Embed: &disgord.Embed{
				Title: "Anime Search Help || alias sa",
				Description: fmt.Sprintf(
					"This is the help for the search anime functionnality\n\n"+
						"You can search an anime by its name using the following syntax\n"+
						"`%ssearchAnime Name`\n",
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
		fmt.Println("Error sending search anime help: ", err)
	}
}

func searchMangaHelp(data *disgord.MessageCreate) {
	_, err := client.CreateMessage(
		ctx,
		data.Message.ChannelID,
		&disgord.CreateMessageParams{
			Embed: &disgord.Embed{
				Title: "Manga Search Help || alias sm",
				Description: fmt.Sprintf(
					"This is the help for the search manga functionnality\n\n"+
						"You can search a manga by its name using the following syntax\n"+
						"`%ssearchManga Name`\n",
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
		fmt.Println("Error sending search manga help: ", err)
	}
}

func formatDescMediaSearch(inputDesc string) (desc string) {
	splitInput := strings.Split(inputDesc, " ")
	if len(splitInput) <= 40 {
		desc = strings.Join(splitInput, " ")
	} else {
		desc = strings.Join(splitInput[:40], " ")
	}
	return
}
