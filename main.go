package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"github.com/andersfylling/disgord"
	"github.com/machinebox/graphql"
)

// GraphURL : GraphQL API URL
var GraphURL = "https://graphql.anilist.co"
var tokenFile = "./token.json"

// Response schema
type Response struct {
	MEDIA struct {
		ID    int
		TITLE struct {
			ROMAJI  string
			ENGLISH string
			NATIVE  string
		}
	}
}

// Query : makes the query
func Query() Response {
	// create a client (safe to share across requests)
	client := graphql.NewClient(GraphURL)

	// make a request
	req := graphql.NewRequest(`
query ($id: Int) {
    Media (id: $id, type: ANIME) {
        id
        title {
        romaji
        english
        native
        }
    }
}
`)

	// set any variables
	req.Var("id", 99263)
	ctx := context.Background()
	// run it and capture the response
	var res Response //interface{}
	if err := client.Run(ctx, req, &res); err != nil {
		log.Fatal(err)
	}
	return res
}

// Connect : Get token from file & connect
func Connect() {
	tok := tokenFromJSON(tokenFile)
	client := disgord.New(disgord.Config{
		BotToken: tok,
	})
	defer client.StayConnectedUntilInterrupted(context.Background())
}

// tokenFromJSON : Reads token from file & returns the token
func tokenFromJSON(file string) (tok string) {
	// open file & defer its closing
	jsonFile, err := os.Open(file)
	if err != nil {
		log.Println(err)
	}
	defer jsonFile.Close()

	// read our opened jsonFile as a byte array & Unmarshal
	byteValue, _ := ioutil.ReadAll(jsonFile)
	err = json.Unmarshal(byteValue, &tok)
	if err != nil {
		log.Println(err)
	}
	return tok
}

func main() {
	Query()
}
