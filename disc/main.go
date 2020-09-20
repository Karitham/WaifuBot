package disc

import (
	"log"

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
