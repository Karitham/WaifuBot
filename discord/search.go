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
						Color:       0x225577,
						Image: &disgord.EmbedImage{
							URL: resp.Character.Image.Large,
						},
					}})
		} else {
			client.SendMsg(ctx, data.Message.ChannelID, err)
		}
	} else {
		client.CreateMessage(
			ctx,
			data.Message.ChannelID,
			&disgord.CreateMessageParams{
				Embed: &disgord.Embed{Title: "Error, search requires at least 1 argument", Color: 0xcc0000}})
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
				Color: 0xeec400,
			},
		})
}
