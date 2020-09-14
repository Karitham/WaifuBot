package discord

import (
	"fmt"
	"log"

	"github.com/Karitham/WaifuBot/database"
	"github.com/Karitham/WaifuBot/query"
	"github.com/andersfylling/disgord"
)

func giveChar(data *disgord.MessageCreate, args CmdArguments) {
	// Verify if user possesses the Waifu he wants to give, also deletes the character from his database if valid
	desc, valid := validGive(data, args)

	// Get the author of the database
	avatar := getUserAvatar(data.Message.Author)

	if valid {
		// Get char
		resp, err := query.CharSearch(args.ParseArgToSearch())
		if err != nil {
			log.Println(err)
		}

		// Add the char to the mentioned user's database
		database.InputClaimChar{
			UserID: data.Message.Mentions[0].ID,
			CharList: database.CharLayout{
				ID:    resp.Character.ID,
				Name:  resp.Character.Name.Full,
				Image: resp.Character.Image.Large,
			}}.AddChar()

		// Send confirmation Message
		_, err = client.CreateMessage(
			ctx,
			data.Message.ChannelID,
			&disgord.CreateMessageParams{
				Embed: &disgord.Embed{
					Title:       "Waifu Given",
					Description: fmt.Sprintf("You gave %s to %s successfully.", resp.Character.Name.Full, data.Message.Mentions[0].Username),
					Thumbnail:   &disgord.EmbedThumbnail{URL: resp.Character.Image.Large},
					Color:       0x43e99a,
				},
			},
		)
		if err != nil {
			log.Println("There was an error giving the character: ", err)
		}
	} else {
		// Send message
		resp, err := client.CreateMessage(
			ctx,
			data.Message.ChannelID,
			&disgord.CreateMessageParams{
				Embed: &disgord.Embed{
					Title:       "Give Waifu Failed",
					Thumbnail:   &disgord.EmbedThumbnail{URL: avatar},
					Description: desc,
					Timestamp:   data.Message.Timestamp,
					Color:       0xcc0000,
				},
			},
		)
		if err != nil {
			log.Println("Create message returned error :", err)
		}
		go deleteMessage(resp, conf.DeleteIllegalRollAfter)
	}
}

// Verify if user possesses the Waifu he wants to give, also deletes the character from his database if valid
func validGive(data *disgord.MessageCreate, arg CmdArguments) (desc string, isValid bool) {
	if len(arg) > 0 {
		resp := arg.ParseArgToSearch()
		switch {
		case resp.ID == 0:
			return fmt.Sprintf("Error, %d is not a valid WaifuID,\nRefer to %shelp to see this command's syntax", resp.ID, conf.Prefix), false
		case data.Message.Mentions == nil:
			return fmt.Sprintf("Error, please tag a discord user,\nRefer to %shelp to see this command's syntax", conf.Prefix), false
		case !database.DelWaifuStruct{UserID: data.Message.Author.ID, CharID: resp.ID}.DelChar():
			return fmt.Sprintf("You do not own the character ID %d,\nVerify if the ID you entered is correct", resp.ID), false
		default:
			return "", true
		}
	}
	return "Please enter arguments,\nRefer to help to see how to use this command", false
}

// Help function for Give
func giveCharHelp(data *disgord.MessageCreate) {
	_, err := client.CreateMessage(
		ctx,
		data.Message.ChannelID,
		&disgord.CreateMessageParams{
			Embed: &disgord.Embed{
				Title: "Give Help || alias g",
				Description: fmt.Sprintf(
					"This is the help for the give functionality\n\n"+
						"This permits you to give one of your beloved waifus to one of your friends.\n"+
						"You can give a waifu to another user using the following syntax :\n"+
						"`%sgive ID @User`",
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
		log.Println("There was an error sending help for give char: ", err)
	}
}
