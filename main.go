package main

import (
	"bot/data"
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
	go data.InitDB()
	disc.BotRun(path)

	// data.SeeWaifus(Karitham)
}
