package disc

import "github.com/andersfylling/disgord"

func help(data *disgord.MessageCreate) {
	client.CreateMessage(
		ctx,
		data.Message.ChannelID,
		&disgord.CreateMessageParams{
			Embed: &disgord.Embed{
				Title: "Help",
				Description: `
				animeSearch (sa) : Search for a anime in the Anilist database 
				roll (r) : Roll a new waifu
				list (l) : List the waifus you have
				search (s) : Search for a character by name / ID
				profile (p) : Display profile information for yourself, or the user mentioned
				favourite (f) : Set a favourite waifu to appear on your profile, you may choose any character you want
				invite : Invite link to add the bot to your server
				trendingAnimes (ta) : Displays the top 10 trending animes from Anilist
				help (h) : Display this help page
				`,
				Color: 0xeec400,
			},
		})
}
