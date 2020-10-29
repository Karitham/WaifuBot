package disc

import (
	"fmt"
	"log"
	"math/rand"
	"strings"
	"sync"
	"time"

	"github.com/Karitham/WaifuBot/database"
	"github.com/Karitham/WaifuBot/query"
	"github.com/diamondburned/arikawa/discord"
	"github.com/diamondburned/arikawa/gateway"
)

// Dropper is used to handle the dropping mechanism
type Dropper struct {
	Waifu   map[discord.ChannelID]query.CharStruct
	ChanInc map[discord.ChannelID]int
	Mux     *sync.Mutex
}

func (bot *Bot) drop(m *gateway.MessageCreateEvent) {
	var err error

	bot.dropper.Mux.Lock()
	defer bot.dropper.Mux.Unlock()

	bot.dropper.Waifu[m.ChannelID], err = query.CharSearchByPopularity(
		rand.New(
			rand.NewSource(
				time.Now().UnixNano(),
			),
		).Intn(c.MaxCharacterRoll),
	)
	if err != nil {
		log.Println(err)
		return
	}

	// Sanitize the name so it's claimable through discord (some characters have double spaces in their name)
	bot.dropper.Waifu[m.ChannelID].Page.Characters[0].Name.Full =
		strings.Join(strings.Fields(bot.dropper.Waifu[m.ChannelID].Page.Characters[0].Name.Full), " ")

	_, err = bot.Ctx.SendMessage(m.ChannelID, "", &discord.Embed{
		Title:       "CHARACTER DROP !",
		Description: "Can you guess who it is ?\nUse w.claim to get this character for yourself",
		Thumbnail: &discord.EmbedThumbnail{
			URL: bot.dropper.Waifu[m.ChannelID].Page.Characters[0].Image.Large,
		},
		Footer: &discord.EmbedFooter{
			Text: "This character's initials are " +
				func(name string) (initials string) {
					for _, v := range strings.Fields(name) {
						initials = initials + strings.ToUpper(string(v[0])) + "."
					}
					return
				}(bot.dropper.Waifu[m.ChannelID].Page.Characters[0].Name.Full),
		},
	})
	if err != nil {
		log.Println(err)
	}
}

// Claim a waifu and adds it to the user's database
func (bot *Bot) Claim(m *gateway.MessageCreateEvent, name ...Name) (*discord.Embed, error) {
	if len(name) == 0 {
		return nil, fmt.Errorf("if you want to claim a character, use `claim <name>`")
	}

	// Lock because we are reading from the map
	bot.dropper.Mux.Lock()
	defer bot.dropper.Mux.Unlock()
	c, ok := bot.dropper.Waifu[m.ChannelID]

	if !ok {
		return nil, fmt.Errorf("there is no character to claim")
	}

	if !strings.EqualFold(
		strings.Join(name, " "),
		c.Page.Characters[0].Name.Full,
	) {
		return nil, fmt.Errorf("wrong name entered")
	}

	// Add to db
	err := database.CharStruct(bot.dropper.Waifu[m.ChannelID]).AddClaimed(m.Author.ID)
	if err != nil {
		return nil, err
	}

	delete(bot.dropper.Waifu, m.ChannelID)

	return &discord.Embed{
		Title: "Claim successful",
		URL:   c.Page.Characters[0].SiteURL,
		Description: fmt.Sprintf(
			"Well done %s you claimed %d\nIt appears in :\n- %s",
			m.Author.Username, c.Page.Characters[0].ID, c.Page.Characters[0].Media.Nodes[0].Title.Romaji,
		),
		Thumbnail: &discord.EmbedThumbnail{
			URL: c.Page.Characters[0].Image.Large,
		},
	}, nil
}
