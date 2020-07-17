package main

import (
	"bot/config"
	"bot/database"
	"bot/disc"
)

func main() {
	// Conf Get the config from config.json
	conf := config.Retrieve("./config.json")

	// Run the database handler in a goroutine
	go database.Init(conf)

	// Run the bot
	disc.BotRun(conf)
}
