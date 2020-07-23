package main

import (
	"bot/config"
	"bot/database"
	"bot/discord"
)

func main() {
	// Get the config from config.json
	conf := config.Retrieve("./config.json")

	// Run the database handler in a goroutine
	go database.Init(conf)

	// Run the bot
	discord.BotRun(conf)
}
