package disc

import (
	"fmt"

	"github.com/andersfylling/disgord"
)

func help(data *disgord.MessageCreate) {
	client.CreateMessage(
		ctx,
		data.Message.ChannelID,
		&disgord.CreateMessageParams{
			Embed: &disgord.Embed{
				Title: "Help",
				Description: fmt.Sprint(
					"`roll` (r) : Roll a new waifu\n",
					"`claim`(c) : Claims the dropped waifu for yourself\n",
					"`list` (l) : List the waifus you, or the user mentionned owns\n",
					"`give` (g) : Give a waifu to the user mentioned\n",
					"`search` (s) : Search for a character by name / ID\n",
					"`profile` (p) : Display profile information for yourself, or the user mentioned\n",
					"`favourite` (f) : Set a favourite waifu to appear on your profile, you may choose any character you want\n",
					"`quote` (q) : Set a custom quote on your profile\n",
					"`searchAnime` (sa) : Search for a anime in the Anilist database\n",
					"`trendingAnimes` (ta) : Displays the top 10 trending animes from Anilist\n",
					"`invite` : Invite link to add the bot to your server\n",
					"`help` (h) : Display this help page\n",
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
