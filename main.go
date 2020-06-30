package main

import (
	"bot/disc"
	"os"
)

func main() {
	const filename = "/config.json"
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	path := dir + filename
	const maxCharQuery = 5000
	disc.BotRun(path, maxCharQuery)
}
