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

	d "github.com/andersfylling/disgord"
	"github.com/machinebox/graphql"
)

// RespChar : Struct for handling character requests from anilists API
type RespChar struct {
	Page struct {
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
		PageInfo struct {
			LastPage int
		}
	}
}

const tokenFile = "./token.json"

var pageTotal int

func main() {
	pageTotal = 80000
	fmt.Println(makeRQ().Page.Characters[0].Image.Large)
	go connect()
}

// makeRQ make the request and setups correctly the page total
func makeRQ() RespChar {
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
	random := r.Int() % pageTotal
	return random
}

// connect : Get token from file & connect
func connect() {
	tok := tokenFromJSON(tokenFile)
	client := d.New(d.Config{BotToken: tok})
	defer client.StayConnectedUntilInterrupted(context.Background())
}

// tokenFromJSON : Reads token from file & returns the token
func tokenFromJSON(file string) (tok string) {
	// open file & defer its closing
	jsonFile, err := os.Open(file)
	defer jsonFile.Close()
	// read our opened jsonFile as a byte array & Unmarshal
	byteValue, err := ioutil.ReadAll(jsonFile)
	if json.Unmarshal(byteValue, &tok) != nil {
		log.Fatal(err)
	}
	return tok
}

// Char : makes a character query
func Char(id int) (RespChar, error) {
	graphURL := "https://graphql.anilist.co"
	client := graphql.NewClient(graphURL)
	var res RespChar

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
	err := client.Run(ctx, req, &res)
	return res, err
}
