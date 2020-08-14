package discord

import (
	"bot/config"
	"bot/query"
	"context"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"

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
	defer func() {
		err := client.StayConnectedUntilInterrupted(ctx)
		if err != nil {
			fmt.Println("The bot is no longer working, ", err)
		}
	}()

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
	cmd, args := ParseMessage(data)

	// Check if it recognises the command, if not, send back an error message
	switch cmd {
	case "search", "s":
		search(data, args)
	case "favourite", "favorite", "f":
		favourite(data, args)
		incDropper(data) // used to drop waifus
	case "trendinganimes", "ta":
		trendingAnime(data, args)
		incDropper(data) // used to drop waifus
	case "searchanime", "sa":
		searchAnime(data, args)
		incDropper(data) // used to drop waifus
	case "give", "g":
		giveChar(data, args)
		incDropper(data) // used to drop waifus
	case "quote", "q":
		quote(data, args)
		incDropper(data) // used to drop waifus
	case "profile", "p":
		profile(data)
		incDropper(data) // used to drop waifus
	case "roll", "r":
		roll(data)
		incDropper(data) // used to drop waifus
	case "list", "l":
		list(data, args)
	case "invite":
		invite(data)
	case "claim", "c":
		claim(data, args)
	case "help", "h":
		help(data, args)
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

func incDropper(data *disgord.MessageCreate) {
	// Increment
	DropIncrement[data.Message.ChannelID]++

	// Higher chances the more you interact with the bot
	r := rand.New(
		rand.NewSource(time.Now().UnixNano()),
	).Intn(conf.DropsOnInteract - DropIncrement[data.Message.ChannelID])

	// Drop
	if r == 0 {
		drop(data)
		DropIncrement[data.Message.ChannelID] = 0
	}
}

func deleteMessage(resp *disgord.Message, sleep time.Duration) {
	time.Sleep(sleep)

	err := client.DeleteMessage(
		ctx,
		resp.ChannelID,
		resp.ID,
	)
	if err != nil {
		fmt.Println("error deleting message :", err)
	}
}

func getUserAvatar(user *disgord.User) (avatar string) {
	avatar, err := user.AvatarURL(128, false)
	if err != nil {
		fmt.Println("There was an error getting this user's avatar", err)
	}
	return
}

// If there is a mention, display the person's profile instead
func getUser(data *disgord.MessageCreate) (user disgord.User) {
	if data.Message.Mentions != nil {
		user = *data.Message.Mentions[0]
	} else {
		user = *data.Message.Author
	}
	return
}
