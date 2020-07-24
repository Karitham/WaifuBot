package main

import (
	"bot/config"
	"bot/database"
	"bot/discord"
	"time"
)

func main() {
	// Get the config from config.json
	conf := config.Retrieve("./config.json")

	// Configure message delete time
	conf.DelMessageAfter *= time.Minute

	// Run the database handler in a goroutine
	go database.Init(conf)

	// Run the bot
	discord.BotRun(conf)
}
