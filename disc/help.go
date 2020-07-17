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
				roll (r) : Roll a new waifu
				list (l) : List the waifus you, or the user mentionned owns
					Syntax : list <Optional page> <@Optional User>
				give (g) : Give a waifu to the user mentioned,
					Syntax : give <ID> <@User>
				search (s) : Search for a character by name / ID
					Syntax : search <ID / Name>
				profile (p) : Display profile information for yourself, or the user mentioned
					Syntax : profile  <@Optional User>
				favourite (f) : Set a favourite waifu to appear on your profile, you may choose any character you want
					Syntax : favourite <ID / Name>
				quote (q) : Set a custom quote on your profile
					Syntax : quote <text>
				searchAnime (sa) : Search for a anime in the Anilist database
					Syntax : searchAnime <ID / Name>
				trendingAnimes (ta) : Displays the top 10 trending animes from Anilist
				invite : Invite link to add the bot to your server
				help (h) : Display this help page
				`,
				Color: 0xeec400,
			},
		})
}
