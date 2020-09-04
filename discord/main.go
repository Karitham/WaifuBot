package discord

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/Karitham/WaifuBot/config"
	"github.com/Karitham/WaifuBot/query"

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

// ListCache is used as a cache for updating list
var ListCache = make(map[disgord.Snowflake]*disgord.Message)

type embedUpdate struct {
	ChannelID disgord.Snowflake
	MessageID disgord.Snowflake
	Embed     disgord.Embed
}

// BotRun the bot and handle events
func BotRun(cf config.ConfJSONStruct) {
	// sets the config for the whole disc package
	conf = cf

	// init the client
	client = disgord.New(
		disgord.Config{
			BotToken: cf.BotToken,
			Presence: &disgord.UpdateStatusPayload{
				Since: nil,
				Game: &disgord.Activity{
					Name: conf.Status,
					Type: disgord.ActivityTypeGame,
				},
				Status: disgord.StatusOnline,
				AFK:    false,
			},
		},
	)

	defer func() {
		err := client.StayConnectedUntilInterrupted(ctx)
		if err != nil {
			log.Println("The bot is no longer working, ", err)
		}
	}()

	// Filter incomming messages & set the prefix
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

	log.Println("The bot is currently running")
}

func reply(s disgord.Session, data *disgord.MessageCreate) {
	cmd, args := ParseMessage(data)

	// Check if it recognises the command, if not, send back an error message
	switch cmd {
	case "search", "s":
		search(data, args)
	case "favourite", "favorite", "f":
		favourite(data, args)
		incDropper(data)
	case "trendinganime", "animetrending","ta":
		trendingMedia(data, "ANIME", args)
		incDropper(data)
	case "trendingmanga", "mangatrending", "tm":
		trendingMedia(data, "MANGA", args)
		incDropper(data)
	case "searchanime", "animesearch", "sa":
		searchMedia(data, "ANIME", args)
		incDropper(data)
	case "searchmanga", "mangasearch", "sm":
		searchMedia(data, "MANGA", args)
		incDropper(data)
	case "give", "g":
		giveChar(data, args)
		incDropper(data)
	case "quote", "q":
		quote(data, args)
		incDropper(data)
	case "profile", "p":
		profile(data)
		incDropper(data)
	case "roll", "r":
		roll(data)
		incDropper(data)
	case "list", "l":
		list(data, args)
	case "invite", "i":
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
	if len(args) > 0 {
		id, err := strconv.Atoi(args[0])
		arg := strings.Join(args, " ")
		if err != nil && id != 0 {
			log.Println(err)
		}
		return query.CharSearchInput{ID: id, Name: arg}
	}
	return query.CharSearchInput{ID: 0, Name: ""}
}

func unknown(data *disgord.MessageCreate) {
	resp, err := client.CreateMessage(
		ctx,
		data.Message.ChannelID,
		&disgord.CreateMessageParams{
			Embed: &disgord.Embed{
				Title:       "Unknown command",
				Description: fmt.Sprintf("Type %shelp to see the commands available", conf.Prefix),
				Timestamp:   data.Message.Timestamp,
				Color:       0xcc0000,
			},
		},
	)
	if err != nil {
		log.Println("Error while creating message :", err)
	}
	go deleteMessage(resp, time.Minute)
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
		log.Println("Error while deleting message :", err)
	}
}

func getUserAvatar(user *disgord.User) (avatar string) {
	avatar, err := user.AvatarURL(128, false)
	if err != nil {
		log.Println("There was an error while getting this user's avatar", err)
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

func setEmbedTo(update embedUpdate) (msg *disgord.Message) {
	msg, err := client.SetMsgEmbed(ctx, update.ChannelID, update.MessageID, &update.Embed)
	if err != nil {
		log.Println("Couldn't update embed :", err)
	}
	return
}
