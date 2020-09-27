package disc

import (
	"fmt"

	"github.com/Karitham/WaifuBot/database"
	"github.com/diamondburned/arikawa/discord"
	"github.com/diamondburned/arikawa/gateway"
)

// Page represent a page
type Page int

// List shows the user's list
func (b *Bot) List(m *gateway.MessageCreateEvent, page ...Page) (*discord.Embed, error) {
	var p = 1
	if len(page) > 0 {
		p = int(page[0])
	}

	uData, err := database.ViewUserData(m.Author.ID)
	if err != nil {
		return nil, err
	}

	embed, err := createListEmbed(m.Author, p-1, uData.Waifus)
	if err != nil {
		return nil, err
	}

	return embed, nil
}

func createListEmbed(user discord.User, page int, list []database.CharLayout) (embed *discord.Embed, err error) {
	return &discord.Embed{
		Title: fmt.Sprintf("%s's list", user.Username),
		Description: func(l []database.CharLayout) (d string) {
			if len(l) >= 0 {
				for i := c.ListLen * page; i < c.ListLen+c.ListLen*page && i < len(l); i++ {
					d += fmt.Sprintf("%d - %s\n", l[i].ID, l[i].Name)
				}
				return d
			}
			return "This user's list is empty"
		}(list),
		Thumbnail: &discord.EmbedThumbnail{URL: user.AvatarURL()},
		Footer:    &discord.EmbedFooter{Text: fmt.Sprintf("Page %02d/%02d", page+1, ((len(list)-1)/c.ListLen)+1)},
	}, nil
}
