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
	disc.BotRun(path)
}
