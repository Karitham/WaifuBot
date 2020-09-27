package disc

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/Karitham/WaifuBot/database"
	"github.com/diamondburned/arikawa/bot"
	"github.com/diamondburned/arikawa/discord"
	"github.com/diamondburned/arikawa/gateway"
)

// List shows the user's list
func (b *Bot) List(m *gateway.MessageCreateEvent, page bot.RawArguments) (*discord.Embed, error) {
	var p int
	var err error

	if page != "" {
		p, err = strconv.Atoi(string(page))
		if err != nil && p < 0 {
			return nil, errors.New("err : invalid page number entered")
		}
	}

	uData, err := database.ViewUserData(m.Author.ID)
	if err != nil {
		return nil, err
	}

	embed, err := createListEmbed(m.Author, p, uData.Waifus)
	if err != nil {
		return nil, err
	}

	return embed, nil
}

func createListEmbed(user discord.User, page int, list []database.CharLayout) (embed *discord.Embed, err error) {
	return &discord.Embed{
		Title: fmt.Sprintf("%s's list page %d", user.Username, page),
		Description: func(l []database.CharLayout) (d string) {
			if len(l) >= 0 {
				for i := c.ListLen * page; i < c.ListLen+c.ListLen*page && i < len(l); i++ {
					d += fmt.Sprintf("%d - %s\n", l[i].ID, l[i].Name)
				}
				return d
			}
			return "This user's list is empty"
		}(list),
	}, nil
}
