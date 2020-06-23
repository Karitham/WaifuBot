package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	d "github.com/andersfylling/disgord"
)

func main() {
	Connect("token.json")
}

// Connect : Get token from file & connect
//
// Input : string (filename) || Output : none
func Connect(tokFile string) {
	tok := tokenFromJSON(tokFile)
	client := d.New(d.Config{
		BotToken: tok,
	})
	defer client.StayConnectedUntilInterrupted(context.Background())
}

// tokenFromJSON : Reads token from file, handles errors & return the token
//
// Input : string (filename) || Output : token (string)
func tokenFromJSON(file string) (tok string) {
	// open file & defer its closing
	jsonFile, err := os.Open(file)
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	// read our opened jsonFile as a byte array & Unmarshal
	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &tok)
	return tok
}
