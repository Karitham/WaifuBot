package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"time"

	"github.com/andersfylling/disgord"
	"github.com/andersfylling/disgord/std"
	"github.com/machinebox/graphql"
	"github.com/sirupsen/logrus"
)

// RespCharT : Struct for handling character requests from anilists API
type RespCharT struct {
	Page struct {
		PageInfo struct {
			LastPage int
		}
		Characters []struct {
			ID    int
			Image struct {
				Large  string
				Medium string
			}
			Name struct {
				First  string
				Full   string
				Last   string
				Native string
			}
		}
	}
}

// ConfigT is used to unmarshal the config.json
type ConfigT struct {
	Prefix   string `json:"Prefix"`
	BotToken string `json:"Bot_Token"`
}

const tokenFile = "config.json"

var config ConfigT
var pageTotal int

var log = &logrus.Logger{
	Out:       os.Stderr,
	Formatter: new(logrus.TextFormatter),
	Hooks:     make(logrus.LevelHooks),
	Level:     logrus.ErrorLevel,
}

func main() {
	pageTotal = 85000
	configFromJSON(tokenFile)
	bot()
}

func bot() {
	client := disgord.New(disgord.Config{
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
	fmt.Println("the bot is currently running")
}

func reply(s disgord.Session, data *disgord.MessageCreate) {
	msg := data.Message
	var resp RespCharT
	// test the message content and respond accordingly
	if msg.Content == "roll" {
		resp = makeRQ()
		response := fmt.Sprintf("https://anilist.co/character/%d", resp.Page.Characters[0].ID)
		msg.Reply(context.Background(), s, response)
		fmt.Println("I just sent a message")
	}
}

// makeRQ make the request and setups correctly the page total
func makeRQ() RespCharT {
	res, err := Char(random())
	if err != nil {
		fmt.Println(err)
	}
	pageTotal = res.Page.PageInfo.LastPage
	return res
}

// random : search the char by ID entered in discord
func random() int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	random := r.Intn(pageTotal)
	return random
}

// configFromJSON : Reads token from file & returns the token
func configFromJSON(file string) ConfigT {
	/* Read file config.json return Type Config */
	body, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println(err)
	}
	json.Unmarshal(body, &config)
	return config
}

// Char : makes a character query
func Char(id int) (res RespCharT, err error) {
	graphURL := "https://graphql.anilist.co"
	client := graphql.NewClient(graphURL)
	req := graphql.NewRequest(`
	query ($pageNumber : Int) {
		Page(perPage: 1, page: $pageNumber) {
			pageInfo {
				lastPage
			}
			characters {
					id
					image {
						large
						medium
					}
					name {
						first
						last
						full
						native
						}
					}
				
		}
	}
	`)

	req.Var("pageNumber", id)

	ctx := context.Background()
	err = client.Run(ctx, req, &res)
	return
}
