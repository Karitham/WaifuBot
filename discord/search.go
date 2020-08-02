package discord

import (
	"bot/query"
	"fmt"

	"github.com/andersfylling/disgord"
)

func search(data *disgord.MessageCreate, args CmdArguments) {
	// check if there is a search term
	if len(args) > 0 {
		resp, err := query.CharSearch(args.ParseArgToSearch())
		if err == nil {
			desc := fmt.Sprintf("I found character %d\nThis character appears in :\n%s", resp.Character.ID, resp.Character.Media.Nodes[0].Title.Romaji)
			client.CreateMessage(
				ctx,
				data.Message.ChannelID,
				&disgord.CreateMessageParams{
					Embed: &disgord.Embed{
						Title:       resp.Character.Name.Full,
						URL:         resp.Character.SiteURL,
						Description: desc,
						Image: &disgord.EmbedImage{
							URL: resp.Character.Image.Large,
						},
						Timestamp: data.Message.Timestamp,
						Color:     0x225577,
					},
				},
			)
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
	client.CreateMessage(
		ctx,
		data.Message.ChannelID,
		&disgord.CreateMessageParams{
			Embed: &disgord.Embed{
				Title: "Character Search Help || alias s",
				Description: fmt.Sprintf(
					"This is the help for search character functionnality\n\n"+
						"You can search a character using :\n"+
						"`%ssearch Name/ID`\n"+
						"You can search by either Name OR ID",
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
