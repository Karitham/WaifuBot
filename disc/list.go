package disc

import (
	"fmt"

	"github.com/Karitham/WaifuBot/database"
	"github.com/diamondburned/arikawa/bot/extras/arguments"
	"github.com/diamondburned/arikawa/discord"
	"github.com/diamondburned/arikawa/gateway"
	"github.com/diamondburned/dgwidgets"
	"go.mongodb.org/mongo-driver/mongo"
)

// List shows the user's list
func (b *Bot) List(m *gateway.MessageCreateEvent, _ ...*arguments.UserMention) error {
	var user discord.User
	if len(m.Mentions) > 0 {
		user = m.Mentions[0].User
	} else {
		user = m.Author
	}

	uData, err := database.ViewUserData(user.ID)
	if err == mongo.ErrNoDocuments {
		return fmt.Errorf("%s has no waifu", user.Username)
	} else if err != nil {
		return err
	}

	// Create widget
	p := dgwidgets.NewPaginator(b.Ctx.Session, m.ChannelID)

	// Make pages
	p.Add(
		func(charlist []database.CharLayout) (embeds []discord.Embed) {
			for j := 0; j <= len(charlist)/c.ListLen; j++ {
				embeds = append(embeds, discord.Embed{
					Title: fmt.Sprintf("%s's list", user.Username),
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

// Verify verify if someone has a waifu
func (b *Bot) Verify(m *gateway.MessageCreateEvent, id CharacterID, _ ...*arguments.UserMention) (string, error) {
	var user discord.User
	if len(m.Mentions) > 0 {
		user = m.Mentions[0].User
	} else {
		user = m.Author
	}

	ok, _ := database.VerifyWaifu(uint(id), uint(user.ID))
	if ok {
		return fmt.Sprintf("%s owns the character", user.Username), nil
	}
	return fmt.Sprintf("%s doesn't own the character", user.Username), nil
}
