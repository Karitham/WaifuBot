package disc

import (
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"time"

	"github.com/Karitham/WaifuBot/config"
	"github.com/diamondburned/arikawa/bot"
	"github.com/diamondburned/arikawa/discord"
	"github.com/diamondburned/arikawa/gateway"
	"github.com/diamondburned/arikawa/utils/wsutil"
)

// Bot represent the bot
type Bot struct {
	Ctx *bot.Context
}

var c config.ConfStruct

// Start starts the bot, registers the command and updates its status
func Start(cf config.ConfStruct) {
	c = cf
	var commands = &Bot{}

	// Start the bot
	wait, err := bot.Start(c.BotToken, commands, func(ctx *bot.Context) error {
		ctx.HasPrefix = bot.NewPrefix(c.Prefix...)

		ctx.SilentUnknown.Command = true
		ctx.SilentUnknown.Subcommand = true

		ctx.MustRegisterSubcommandCustom(&Search{}, "search")
		ctx.MustRegisterSubcommandCustom(&Trending{}, "trending")

		ctx.ChangeCommandInfo("Roll", "", "roll a random character")
		ctx.ChangeCommandInfo("Profile", "", "display user profile")
		ctx.ChangeCommandInfo("Favorite", "", "set a char as favorite")
		ctx.ChangeCommandInfo("Quote", "", "set profile quote")
		ctx.ChangeCommandInfo("List", "", "display user characters")
		ctx.ChangeCommandInfo("Help", "", "display general help")
		ctx.ChangeCommandInfo("Give", "", "give a char to a user")
		ctx.ChangeCommandInfo("Invite", "", "send invite link")
		ctx.ChangeCommandInfo("Claim", "", "claim a dropped character")

		ctx.AddAliases("List", "l", "L")
		ctx.AddAliases("Invite", "i", "I")
		ctx.AddAliases("Roll", "r", "R")
		ctx.AddAliases("Profile", "p", "P")
		ctx.AddAliases("Help", "h", "H")
		ctx.AddAliases("Favorite", "f", "F")
		ctx.AddAliases("Quote", "q", "Q")
		ctx.AddAliases("Give", "g", "G")
		ctx.AddAliases("Claim", "c", "C")

		ctx.Gateway.Identifier.IdentifyData = gateway.IdentifyData{
			Token: c.BotToken,

			Presence: &gateway.UpdateStatusData{
				Game: &discord.Activity{
					Name: c.BotStatus,
					Type: discord.GameActivity,
				},
				Status: discord.OnlineStatus,
			},
		}

		ctx.AddHandler(func(m *gateway.MessageCreateEvent) {
			// Filter bot message
			if m.Author.ID == ctx.Ready.User.ID {
				return
			}
			// Higher chances the more you interact with the bot
			r := rand.New(
				rand.NewSource(time.Now().UnixNano()),
			).Intn(c.DropsOnInteract - d.ChanInc[m.ChannelID])

			if r == 0 {
				commands.drop(m)
				d.ChanInc[m.ChannelID] = 0
			}
		})

		wsutil.WSDebug = func(v ...interface{}) {
			log.Println("WaifuBot - Gateway Debug: ", v)
		}

		return nil
	})
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("Bot started")

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
			b.Ctx.Ready.User.ID,
			1073801280,
		),
	}, nil
}
