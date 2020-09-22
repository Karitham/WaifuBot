package disc

import (
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
		if err != nil {
			return nil, err
		}
	}

	uData, err := database.ViewUserData(m.Author.ID)
	if err != nil {
		return nil, err
	}

	desc := func(l []database.CharLayout) (d string) {
		if len(l) > p*15+15 {
			l = l[p*15 : p*15+15]
		} else if len(l) > p*15 {
			l = l[p*15:]
		}
		for _, waifu := range l {
			d += fmt.Sprintf("%d - %s\n", waifu.ID, waifu.Name)
		}
		return
	}(uData.Waifus)

	return &discord.Embed{
		Title:       fmt.Sprintf("%s's list page %d", m.Author.Username, p),
		Description: desc,
	}, nil
}
