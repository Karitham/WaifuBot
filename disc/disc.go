package disc

import (
	"fmt"
	"math/rand"
	"strconv"
	"sync"
	"time"

	"github.com/Karitham/WaifuBot/anilist"
	"github.com/Karitham/WaifuBot/config"
	"github.com/Karitham/WaifuBot/db"
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
	conf    *config.ConfStruct
	conn    *db.Queries
}

// Start starts the bot, registers the command and updates its status
func Start(configuration *config.ConfStruct, connection *db.Queries) (func() error, error) {
	var b = &Bot{
		Ctx: &bot.Context{},
		dropper: &Dropper{
			Waifu:   make(map[discord.ChannelID]anilist.CharStruct),
			ChanInc: make(map[discord.ChannelID]uint64),
			Mutex:   new(sync.Mutex),
		},
		seed: rand.New(rand.NewSource(time.Now().UnixNano())),
		conf: configuration,
		conn: connection,
	}

	// Start the bot
	waitFn, err := bot.Start(b.conf.BotToken, b, func(ctx *bot.Context) error {
		ctx.HasPrefix = bot.NewPrefix(b.conf.Prefix...)

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
			Token: b.conf.BotToken,

			Presence: &gateway.UpdateStatusData{
				Activities: []discord.Activity{
					{
						Name: b.conf.BotStatus,
						Type: discord.GameActivity,
					},
				},

				Status: gateway.OnlineStatus,
			},
		}

		ctx.Session.Gateway.AddIntents(gateway.IntentGuildMessageReactions)
		ctx.AddHandler(func(m *gateway.MessageCreateEvent) {
			// Filter bot message
			if m.Author.Bot {
				return
			}

			// Higher chances the more you interact with the bot
			r := b.seed.Uint64() % (b.conf.DropsOnInteract - b.dropper.ChanInc[m.ChannelID])

			if r == 0 || (b.conf.DropsOnInteract-b.dropper.ChanInc[m.ChannelID]) == 1 {
				b.drop(m)
				b.dropper.ChanInc[m.ChannelID] = 0
			}

			b.dropper.ChanInc[m.ChannelID]++
		})

		return nil
	})
	if err != nil {
		return nil, err
	}

	b.Me, err = b.Ctx.Me()
	if err != nil {
		return nil, err
	}

	return waitFn, nil
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
			92224,
		),
	}, nil
}

func parseUser(m *gateway.MessageCreateEvent) (user discord.User) {
	if len(m.Mentions) > 0 {
		return m.Mentions[0].User
	}
	return m.Author
}
