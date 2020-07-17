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

		// handler
		reply, // call reply func
	) // handles copy

	fmt.Println("The bot is currently running")
}

func reply(s disgord.Session, data *disgord.MessageCreate) {
	command, args := ParseMessage(data)

	// Check if it recognises the command, if not, send back an error message
	switch {
	case command == "search" || command == "s":
		search(data, args)
	case command == "favourite" || command == "favorite" || command == "f":
		favourite(data, args)
	case command == "trendinganimes" || command == "ta":
		trendingAnime(data, args)
	case command == "searchanime" || command == "sa":
		searchAnime(data, args)
	case command == "give" || command == "g":
		giveChar(data, args)
	case command == "quote" || command == "q":
		quote(data, args)
	case command == "profile" || command == "p":
		profile(data)
	case command == "help" || command == "h":
		help(data)
	case command == "roll" || command == "r":
		roll(data)
	case command == "list" || command == "l":
		list(data, args)
	case command == "invite":
		invite(data)
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
