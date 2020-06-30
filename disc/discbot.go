package disc

import (
	"bot/query"
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

var maxCharQuery int
var botURL string

var log = &logrus.Logger{
	Out:       os.Stderr,
	Formatter: new(logrus.TextFormatter),
	Hooks:     make(logrus.LevelHooks),
	Level:     logrus.ErrorLevel,
}

// BotRun the bot and handle events
func BotRun(configfile string) {
	config := configFromJSON(configfile)
	maxCharQuery = config.MaxChar
	var client = disgord.New(disgord.Config{
		BotToken: config.BotToken,
		Logger:   log,
	})
	defer client.StayConnectedUntilInterrupted(context.Background())

	log, _ := std.NewLogFilter(client)
	filter, _ := std.NewMsgFilter(context.Background(), client)
	filter.SetPrefix(config.Prefix)

	// create a handler and bind it to new message events
	// tip: read the documentation for std.CopyMsgEvt and understand why it is used here.
	go client.On(disgord.EvtMessageCreate,
		// middleware
		filter.NotByBot,    // ignore bot messages
		filter.HasPrefix,   // read original
		log.LogMsg,         // log command message
		std.CopyMsgEvt,     // read & copy original
		filter.StripPrefix, // write copy
		// handler
		reply) // handles copy
	fmt.Println("The bot is currently running")
	botURL, _ = client.InviteURL(context.Background())
}

func reply(s disgord.Session, data *disgord.MessageCreate) {
	msg := data.Message
	var resp query.RespCharType
	// test the message content and respond accordingly
	if msg.Content == "roll" {
		resp = query.MakeRQ(maxCharQuery)
		response := fmt.Sprintf("https://anilist.co/character/%d", resp.Page.Characters[0].ID)
		msg.Reply(context.Background(), s, response)
	}
	if msg.Content == "invite" {
		msg.Reply(context.Background(), s, botURL)
	}
}

// Read file config.json return Type Config
// configFromJSON : Reads token from file & returns the token
func configFromJSON(file string) ConfigT {
	var config ConfigT
	body, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println(err)
	}
	json.Unmarshal(body, &config)
	return config
}
