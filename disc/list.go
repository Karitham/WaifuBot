package disc

import (
	"fmt"
	"strings"

	"github.com/Karitham/WaifuBot/database"
	"github.com/diamondburned/arikawa/v2/bot/extras/arguments"
	"github.com/diamondburned/arikawa/v2/discord"
	"github.com/diamondburned/arikawa/v2/gateway"
	"github.com/diamondburned/dgwidgets"
	"go.mongodb.org/mongo-driver/mongo"
)

// List shows the user's list
func (b *Bot) List(m *gateway.MessageCreateEvent, _ ...*arguments.UserMention) error {
	user := parseUser(m)

	uData, err := database.ViewUserData(user.ID)
	if err == mongo.ErrNoDocuments {
		return fmt.Errorf("%s has no waifu", user.Username)
	} else if err != nil {
		return err
	}

	// Create widget
	p := dgwidgets.NewPaginator(b.Ctx.State, m.ChannelID)

	// What to do when timeout
	p.SetTimeout(c.ListMaxUpdateTime.Duration)
	p.ColourWhenDone = 0xFFFF00

	// Make pages
	p.Add(
		func(charlist []database.CharLayout) (embeds []discord.Embed) {
			for j := 0; j <= len(charlist)/c.ListLen; j++ {
				embeds = append(embeds, discord.Embed{

					Title: fmt.Sprintf("%s's list", user.Username),

					Description: func(l []database.CharLayout) string {
						var s strings.Builder

						if len(l) >= 0 {
							for i := c.ListLen * j; i < c.ListLen+c.ListLen*j && i < len(l); i++ {
								s.WriteString(fmt.Sprintf("`%d`\f - %s\n", l[i].ID, l[i].Name))
							}
							return s.String()
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

	return p.Spawn()
}

// Verify verify if someone has a waifu
func (b *Bot) Verify(m *gateway.MessageCreateEvent, id database.CharID, _ ...*arguments.UserMention) (string, error) {
	var user discord.User
	if len(m.Mentions) > 0 {
		user = m.Mentions[0].User
	} else {
		user = m.Author
	}

	ok, _ := id.VerifyWaifu(m.Author.ID)
	if ok {
		return fmt.Sprintf("%s owns the character", user.Username), nil
	}
	return fmt.Sprintf("%s doesn't own the character", user.Username), nil
}
