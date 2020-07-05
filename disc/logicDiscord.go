package disc

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/andersfylling/disgord"
	"github.com/andersfylling/disgord/std"
	"github.com/sirupsen/logrus"
)

// ConfigT is used to unmarshal the config.json
type ConfigT struct {
	Prefix   string `json:"Prefix"`
	BotToken string `json:"Bot_Token"`
	MaxChar  int    `json:"Max_Character_Roll"`
}

// Global Variables to ease working with client/sesion etc
var ctx = context.Background()
var config ConfigT
var client *disgord.Client
var session disgord.Session

// BotRun the bot and handle events
func BotRun(configfile string) {
	config = configFromJSON(configfile)

	var log = &logrus.Logger{
		Out:       os.Stderr,
		Formatter: new(logrus.TextFormatter),
		Hooks:     make(logrus.LevelHooks),
		Level:     logrus.ErrorLevel,
	}

	client = disgord.New(disgord.Config{
		BotToken: config.BotToken,
		Logger:   log,
	})

	defer client.StayConnectedUntilInterrupted(ctx)

	filter, _ := std.NewMsgFilter(ctx, client)
	filter.SetPrefix(config.Prefix)

	// create a handler and bind it to new message events
	go client.On(disgord.EvtMessageCreate,
		// middleware
		filter.NotByBot,    // ignore bot messages
		filter.HasPrefix,   // read original
		std.CopyMsgEvt,     // read & copy original
		filter.StripPrefix, // write copy

		// handler
		reply, // call reply func
		// specific
	) // handles copy

	fmt.Println("The bot is currently running")
}

func reply(s disgord.Session, data *disgord.MessageCreate) {

	// Parses the message into command / args
	command := strings.Fields(data.Message.Content)[0]
	args := strings.Fields(data.Message.Content)[1:]

	switch {
	case command == "search" || command == "s":
		search(data, args)
	case command == "help" || command == "h":
		help(data)
	case command == "roll" || command == "r":
		roll(data)
	case command == "list" || command == "l":
		list(data)
	case command == "invite":
		invite(data)
	default:
		unknown(data)
	}
}
