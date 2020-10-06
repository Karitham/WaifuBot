package disc

import (
	"fmt"

	"github.com/Karitham/WaifuBot/database"
	"github.com/diamondburned/arikawa/discord"
	"github.com/diamondburned/arikawa/gateway"
	"github.com/diamondburned/dgwidgets"
)

// Page represent a page
type Page = int

// List shows the user's list
func (b *Bot) List(m *gateway.MessageCreateEvent) error {
	uData, err := database.ViewUserData(m.Author.ID)
	if err != nil {
		return err
	}

	// Create widget
	p := dgwidgets.NewPaginator(b.Ctx.Session, m.ChannelID)

	// Make pages
	p.Add(
		func(charlist []database.CharLayout) (embeds []discord.Embed) {
			for j := 0; j <= len(charlist)/c.ListLen; j++ {
				embeds = append(embeds, discord.Embed{
					Title: fmt.Sprintf("%s's list", m.Author.Username),
					Description: func(l []database.CharLayout) (d string) {
						if len(l) >= 0 {
							for i := c.ListLen * j; i < c.ListLen+c.ListLen*j && i < len(l); i++ {
								d += fmt.Sprintf("`%d`\f - %s\n", l[i].ID, l[i].Name)
							}
							return d
						}
						return ""
					}(charlist),
					Footer: &discord.EmbedFooter{
						Text: fmt.Sprintf("Page %d out of %d", j+1, len(charlist)/c.ListLen+1),
					},
					Color: 3447003,
				})
			}
			return embeds
		}(uData.Waifus)...,
	)

	// What to do when timeout
	p.ColourWhenDone = 0xFFFF00
	p.DeleteReactionsWhenDone = true
	p.Widget.Timeout = c.ListMaxUpdateTime.Duration

	return p.Spawn()
}
