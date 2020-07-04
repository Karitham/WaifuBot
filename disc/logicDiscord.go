package disc

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

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
	msg := data.Message
	switch {
	case msg.Content == "searchID" || msg.Content == "sID":
		//searchID(data, search)
	case msg.Content == "search" || msg.Content == "s":
		//searchName(data, search)
	case msg.Content == "help" || msg.Content == "h":
		help(data)
	case msg.Content == "roll" || msg.Content == "r":
		roll(data)
	case msg.Content == "list" || msg.Content == "l":
		list(data)
	case msg.Content == "invite":
		invite(data)
	default:
		unknown(data)
	}
}

// configFromJSON reads config from file
func configFromJSON(file string) ConfigT {
	var config ConfigT
	body, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println(err)
	}
	json.Unmarshal(body, &config)
	return config
}
