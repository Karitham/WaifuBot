package main

import (
	"github.com/Karitham/WaifuBot/config"
	"github.com/Karitham/WaifuBot/database"
	"github.com/Karitham/WaifuBot/discord"
)

func main() {
	// Get the config from config.json
	conf := config.Retrieve("config.json")

	// Run the services
	database.Init(conf)
	discord.BotRun(conf)
}
