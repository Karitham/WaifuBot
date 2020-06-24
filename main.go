package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"time"

	q "querries"

	"github.com/andersfylling/disgord"
)

var tokenFile = "./token.json"

func main() {
	res := q.Char(random())
	for res.Character.Name.Full == "" {
		res = q.Char(random())
	}
	fmt.Println(res.Character.ID)
}

// random : search the char by ID entered in discord
func random() int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	random := r.Int() % 63000
	return random
}

// connect : Get token from file & connect
func connect() {
	tok := tokenFromJSON(tokenFile)
	client := disgord.New(disgord.Config{BotToken: tok})
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
