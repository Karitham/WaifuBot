package disc

import (
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"sync"
	"time"

	"github.com/Karitham/WaifuBot/config"
	"github.com/Karitham/WaifuBot/query"
	"github.com/diamondburned/arikawa/v2/bot"
	"github.com/diamondburned/arikawa/v2/discord"
	"github.com/diamondburned/arikawa/v2/gateway"
)

// Bot represent the bot
type Bot struct {
	Ctx     *bot.Context
	dropper *Dropper
	seed    rand.Source64
	Me      *discord.User
}

var c *config.ConfStruct

// Start starts the bot, registers the command and updates its status
func Start(cf *config.ConfStruct) {
	c = cf
	var b = &Bot{
		Ctx: &bot.Context{},
		dropper: &Dropper{
			Waifu:   make(map[discord.ChannelID]query.CharStruct),
			ChanInc: make(map[discord.ChannelID]uint64),
			Mux:     new(sync.Mutex),
		},
		seed: rand.New(rand.NewSource(time.Now().UnixNano())),
	}

	// Start the bot
	wait, err := bot.Start(c.BotToken, b, func(ctx *bot.Context) error {
		ctx.HasPrefix = bot.NewPrefix(c.Prefix...)

		ctx.SilentUnknown.Command = true
		ctx.SilentUnknown.Subcommand = true

		ctx.MustRegisterSubcommand(&Search{}, "search", "s")
		ctx.MustRegisterSubcommand(&Trending{}, "trending", "t")

		ctx.ChangeCommandInfo("Roll", "", "roll a random character")
		ctx.ChangeCommandInfo("Profile", "", "display user profile")
		ctx.ChangeCommandInfo("Favorite", "", "set a char as favorite")
		ctx.ChangeCommandInfo("Quote", "", "set profile quote")
		ctx.ChangeCommandInfo("List", "", "display user characters")
		ctx.ChangeCommandInfo("Help", "", "display general help")
		ctx.ChangeCommandInfo("Give", "", "give a char to a user")
		ctx.ChangeCommandInfo("Invite", "", "send invite link")
		ctx.ChangeCommandInfo("Verify", "", "check if a user owns the waifu")
		ctx.ChangeCommandInfo("Claim", "", "claim a dropped character")

		ctx.AddAliases("List", "l", "L")
		ctx.AddAliases("Invite", "i", "I")
		ctx.AddAliases("Roll", "r", "R")
		ctx.AddAliases("Profile", "p", "P")
		ctx.AddAliases("Help", "h", "H")
		ctx.AddAliases("Favorite", "f", "F")
		ctx.AddAliases("Quote", "q", "Q")
		ctx.AddAliases("Give", "g", "G")
		ctx.AddAliases("Verify", "v", "V")
		ctx.AddAliases("Claim", "c", "C")

		ctx.Gateway.Identifier.IdentifyData = gateway.IdentifyData{
			Token: c.BotToken,

			Presence: &gateway.UpdateStatusData{
				Activities: &[]discord.Activity{
					{
						Name: c.BotStatus,
						Type: discord.GameActivity,
					},
				},

				Status: discord.OnlineStatus,
			},
		}

		ctx.AddHandler(func(m *gateway.MessageCreateEvent) {
			// Filter bot message
			if m.Author.Bot {
				return
			}

			// Higher chances the more you interact with the bot
			b.dropper.ChanInc[m.ChannelID]++
			r := b.seed.Uint64() % (c.DropsOnInteract - b.dropper.ChanInc[m.ChannelID])

			if r == 0 {
				b.drop(m)
				b.dropper.ChanInc[m.ChannelID] = 0
			}
		})

		return nil
	})
	if err != nil {
		log.Fatalln(err)
	}

	b.Me, err = b.Ctx.Me()
	if err != nil {
		log.Println(err)
	}

	guilds, err := b.Ctx.Guilds()
	if err != nil {
		log.Println(err)
	}
	log.Printf("Bot started in %d servers", len(guilds))

	// Wait for closing
	if err := wait(); err != nil {
		log.Fatalln("Gateway fatal error:", err)
	}
}

func parseArgs(b string) (ID int) {
	if id, err := strconv.Atoi(string(b)); id != 0 && err == nil {
		ID = id
	}
	return
}

// Invite sends invite link
func (b *Bot) Invite(_ *gateway.MessageCreateEvent) (*discord.Embed, error) {
	return &discord.Embed{
		Title: "Invite",

		URL: fmt.Sprintf(
			"https://discord.com/oauth2/authorize?scope=bot&client_id=%d&permissions=%d",
			b.Me.ID,
			1073801280,
		),
	}, nil
}

func parseUser(m *gateway.MessageCreateEvent) (user discord.User) {
	if len(m.Mentions) > 0 {
		return m.Mentions[0].User
	}
	return m.Author
}
