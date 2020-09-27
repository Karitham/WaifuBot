package disc

import (
	"fmt"
	"log"
	"strconv"

	"github.com/Karitham/WaifuBot/config"
	"github.com/diamondburned/arikawa/bot"
	"github.com/diamondburned/arikawa/discord"
	"github.com/diamondburned/arikawa/gateway"
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
		ctx.MustRegisterSubcommand(&Search{})
		ctx.MustRegisterSubcommand(&Trending{})

		ctx.AddAliases("List", "l", "L")
		ctx.AddAliases("Roll", "r", "R")
		ctx.AddAliases("Profile", "p", "P")
		ctx.AddAliases("Help", "h", "H")
		ctx.AddAliases("Favorite", "f", "F")
		ctx.AddAliases("Quote", "q", "Q")
		ctx.AddAliases("Give", "g", "G")
		return nil
	})
	if err != nil {
		log.Fatalln(err)
	}

	// Set status
	err = commands.Ctx.Gateway.UpdateStatus(
		gateway.UpdateStatusData{
			Game: &discord.Activity{
				Name: c.BotStatus,
				Type: discord.GameActivity,
			},
			Status: discord.OnlineStatus,
		},
	)
	if err != nil {
		log.Println("couldn't set status : ", err)
	}

	log.Println("Bot started")

	// Wait for closing
	if err := wait(); err != nil {
		log.Fatalln("Gateway fatal error:", err)
	}
}

func parseArgs(b string) (name string, ID int) {
	if id, err := strconv.Atoi(string(name)); id != 0 && err == nil {
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
			67497024,
		),
	}, nil
}
