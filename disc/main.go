package disc

import (
	"bot/config"
	"bot/query"
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/andersfylling/disgord"
	"github.com/andersfylling/disgord/std"
)

// CmdArguments represents the arguments entered by the user after a command
type CmdArguments []string

// Global Variables to ease working with client/sesion etc
var ctx = context.Background()
var client *disgord.Client
var session disgord.Session
var conf config.ConfJSONStruct

// DropIncrement controls the dropping
var DropIncrement = make(map[disgord.Snowflake]int)

// BotRun the bot and handle events
func BotRun(cf config.ConfJSONStruct) {
	// sets the config for the whole disc package
	conf = cf

	// init the client
	client = disgord.New(disgord.Config{BotToken: cf.BotToken})

	// stay connected to discord
	defer client.StayConnectedUntilInterrupted(ctx)

	// filter incomming messages & set the prefix
	filter, _ := std.NewMsgFilter(ctx, client)
	filter.SetPrefix(cf.Prefix)

	// create a handler and bind it to new message events
	go client.On(disgord.EvtMessageCreate,
		// middleware
		filter.NotByBot,    // ignore bot messages
		filter.HasPrefix,   // read original
		std.CopyMsgEvt,     // read & copy original
		filter.StripPrefix, // write copy
		increment,          // used to drop waifus

		// handler
		reply, // call reply func
	) // handles copy

	fmt.Println("The bot is currently running")
}

func reply(s disgord.Session, data *disgord.MessageCreate) {
	cmd, args := ParseMessage(data)

	// Check if it recognises the command, if not, send back an error message
	switch {
	case cmd == "search" || cmd == "s":
		search(data, args)
	case cmd == "favourite" || cmd == "favorite" || cmd == "f":
		favourite(data, args)
	case cmd == "trendinganimes" || cmd == "ta":
		trendingAnime(data, args)
	case cmd == "searchanime" || cmd == "sa":
		searchAnime(data, args)
	case cmd == "give" || cmd == "g":
		giveChar(data, args)
	case cmd == "quote" || cmd == "q":
		quote(data, args)
	case cmd == "profile" || cmd == "p":
		profile(data)
	case cmd == "help" || cmd == "h":
		help(data)
	case cmd == "roll" || cmd == "r":
		roll(data)
	case cmd == "list" || cmd == "l":
		list(data, args)
	case cmd == "invite":
		invite(data)
	case cmd == "claim" || cmd == "c":
		claim(data, args)
	default:
		unknown(data)
	}

}

// ParseMessage parses the message into command / args
func ParseMessage(data *disgord.MessageCreate) (string, []string) {
	var command string
	var args []string

	if len(data.Message.Content) > 0 {
		command = strings.ToLower(strings.Fields(data.Message.Content)[0])
		if len(data.Message.Content) > 1 {
			args = strings.Fields(data.Message.Content)[1:]
		}
	}
	return command, args
}

// ParseArgToSearch parses any arg to an int, if no int is entered, returns 0 as the result
func (args CmdArguments) ParseArgToSearch() query.CharSearchInput {
	id, err := strconv.Atoi(args[0])
	arg := strings.Join(args, " ")
	if err != nil && id != 0 {
		fmt.Println(err)
	}
	return query.CharSearchInput{ID: id, Name: arg}
}

func increment(s disgord.Session, data *disgord.MessageCreate) {
	// Increment
	DropIncrement[data.Message.ChannelID]++

	// Drop
	if DropIncrement[data.Message.ChannelID] >= conf.DropsOnInteract {
		drop(data)
		DropIncrement[data.Message.ChannelID] = 0
	}
}
