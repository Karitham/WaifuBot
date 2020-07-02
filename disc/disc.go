package disc

import (
	db "bot/data"
	"bot/query"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"

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
var ctx = context.Background()
var config ConfigT
var client *disgord.Client

var log = &logrus.Logger{
	Out:       os.Stderr,
	Formatter: new(logrus.TextFormatter),
	Hooks:     make(logrus.LevelHooks),
	Level:     logrus.ErrorLevel,
}

// BotRun the bot and handle events
func BotRun(configfile string) {
	config = configFromJSON(configfile)
	maxCharQuery = config.MaxChar

	client = disgord.New(disgord.Config{
		BotToken: config.BotToken,
		Logger:   log,
	})
	defer client.StayConnectedUntilInterrupted(ctx)

	log, _ := std.NewLogFilter(client)
	filter, _ := std.NewMsgFilter(ctx, client)
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
		reply, // call reply func
		// specific
	) // handles copy

	fmt.Println("The bot is currently running")
}

func reply(s disgord.Session, data *disgord.MessageCreate) {
	msg := data.Message
	var resp query.RespCharType

	// helps the user and
	if msg.Content == "help" || msg.Content == "h" {
		help := fmt.Sprintf("The commands available to you right now are :\n\t%sroll\n\t%sinvite", config.Prefix, config.Prefix)
		msg.Reply(ctx, s, help)
	}

	// send back the URL to a waifu
	if msg.Content == "roll" || msg.Content == "r" {
		resp = query.MakeRQ(maxCharQuery)
		db.AddWaifu(db.UserBson{UserID: msg.Author.ID, Date: time.Now(), Waifu: resp.Page.Characters[0].ID})
		response := fmt.Sprintf("https://anilist.co/character/%d", resp.Page.Characters[0].ID)
		msg.Reply(ctx, s, response)
	}

	// send back bot invite url
	if msg.Content == "invite" {
		botURL, err := client.InviteURL(ctx)
		if err != nil {
			msg.Reply(ctx, s, err)
		}
		msg.Reply(ctx, s, botURL)
	}

}

// Read file config.json return Type ConfigType
//
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
