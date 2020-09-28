package discord

import (
	"fmt"

	"github.com/Karitham/WaifuBot/database"
	"github.com/Karitham/WaifuBot/query"
	"github.com/andersfylling/disgord"
)

func check(data *disgord.MessageCreate, args CmdArguments) {
	// Verify if user possesses the Waifu.
	desc, valid := checkWaifuValid(data, args)

	// Get the avatar of the author
	avatar := getUserAvatar(data.Message.Author)

	if valid {
		// Get char
		_, err := query.CharSearch(args.ParseArgToSearch())
		if err != nil {
			fmt.Println(err)
		}

		// Send confirmation Message
		_, err = client.CreateMessage(
			ctx,
			data.Message.ChannelID,
			&disgord.CreateMessageParams{
				Embed: &disgord.Embed{
					Title:       "Waifu Verification",
					Thumbnail:   &disgord.EmbedThumbnail{URL: avatar},
					Description: fmt.Sprintf("%s possesses the Waifu you want.", data.Message.Mentions[0].Username),
					Timestamp:   data.Message.Timestamp,
					Color:       0x43e99a,
				},
			},
		)
		if err != nil {
			fmt.Println("There was an error verifying: ", err)
		}
	} else {
		// Send message
		_, err := client.CreateMessage(
			ctx,
			data.Message.ChannelID,
			&disgord.CreateMessageParams{
				Embed: &disgord.Embed{
					Title:       "Waifu Verification",
					Thumbnail:   &disgord.EmbedThumbnail{URL: avatar},
					Description: desc,
					Timestamp:   data.Message.Timestamp,
					Color:       0xcc0000,
				},
			},
		)
		if err != nil {
			fmt.Println("Create message returned error :", err)
		}
	}
}

// Verify if user possesses the Waifu he wants to give, also deletes the character from his database if valid
func checkWaifuValid(data *disgord.MessageCreate, arg CmdArguments) (desc string, isValid bool) {
	if len(arg) > 0 {
		resp := arg.ParseArgToSearch()
		switch {
		case resp.ID == 0:
			return fmt.Sprintf("Error, %d is not a valid WaifuID,\nRefer to %shelp to see this command's syntax", resp.ID, conf.Prefix), false
		case data.Message.Mentions == nil:
			return fmt.Sprintf("Error, please tag a discord user,\nRefer to %shelp to see this command's syntax", conf.Prefix), false
		case !database.CheckWaifuStruct{UserID: data.Message.Mentions[0].ID, CharID: resp.ID}.CheckWaifu():
			return fmt.Sprintf("%s does not possess this waifu.", data.Message.Mentions[0].Username), false
		default:
			return "", true
		}
	}
	return "Please enter arguments,\nRefer to help to see how to use this command", false
}

// Help function for Give
func checkHelp(data *disgord.MessageCreate) {
	_, err := client.CreateMessage(
		ctx,
		data.Message.ChannelID,
		&disgord.CreateMessageParams{
			Embed: &disgord.Embed{
				Title: "Check Help || alias chk",
				Description: fmt.Sprintf(
					"This is the help for the check functionality\n\n"+
						"This permits you to check if one of your friends has got the waifu you want.\n"+
						"You can give a waifu to another user using the following syntax :\n"+
						"`%scheck ID @User`",
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
		fmt.Println("There was an error sending help for check waifu: ", err)
	}
}
