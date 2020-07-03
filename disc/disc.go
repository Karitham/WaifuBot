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
	switch {
	case msg.Content == "help" || msg.Content == "h":
		help(s, data)
	case msg.Content == "roll" || msg.Content == "r":
		roll(s, data)
	case msg.Content == "invite":
		invite(s, data)
	case msg.Content == "list" || msg.Content == "l":
		list(s, data)
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

func list(s disgord.Session, data *disgord.MessageCreate) {
	var desc string
	waifuList := db.SeeWaifus(data.Message.Author.ID)
	waifus := func() string {
		for _, v := range waifuList {
			desc = fmt.Sprintf("%d\n%s", v, desc)
		}
		return desc
	}
	client.CreateMessage(
		ctx,
		data.Message.ChannelID,
		&disgord.CreateMessageParams{
			Embed: &disgord.Embed{
				Title:       "Waifu list",
				Description: waifus(),
				Color:       0x88ffcc,
			}})
}

func roll(s disgord.Session, data *disgord.MessageCreate) {
	resp := query.MakeRQ(maxCharQuery)
	db.AddWaifu(db.InputStruct{UserID: data.Message.Author.ID, Date: time.Now(), Waifu: resp.Page.Characters[0].ID})
	desc := fmt.Sprintf("You rolled waifu id : %d", resp.Page.Characters[0].ID)
	url := fmt.Sprintf("https://anilist.co/character/%d", resp.Page.Characters[0].ID)
	client.CreateMessage(
		ctx,
		data.Message.ChannelID,
		&disgord.CreateMessageParams{
			Embed: &disgord.Embed{
				Title:       resp.Page.Characters[0].Name.Full,
				URL:         url,
				Description: desc,
				Color:       0x225577,
				Image: &disgord.EmbedImage{
					URL: resp.Page.Characters[0].Image.Large,
				},
			}})
}

func invite(s disgord.Session, data *disgord.MessageCreate) {
	botURL, err := client.InviteURL(ctx)
	if err != nil {
		data.Message.Reply(ctx, s, err)
	}
	client.CreateMessage(
		ctx,
		data.Message.ChannelID,
		&disgord.CreateMessageParams{
			Embed: &disgord.Embed{
				Title: "Invite",
				URL:   botURL,
				Color: 0x49b675,
			},
		})
}

func help(s disgord.Session, data *disgord.MessageCreate) {
	client.CreateMessage(
		ctx,
		data.Message.ChannelID,
		&disgord.CreateMessageParams{
			Embed: &disgord.Embed{
				Title: "Help",
				Description: `
				roll (r): 	Roll a new waifu
				list (l): 	List the waifus you have
				invite : 	Invite link to add the bot to your server
				help (h) :	show the commands you can use
				`,
				Color: 0xcc0000,
			},
		})
}
