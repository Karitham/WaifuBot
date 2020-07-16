package main

import (
	"bot/config"
	"bot/database"
	"fmt"
)

func main() {
	// Get the config from config.json
	conf := config.Retrieve("./config.json")

	// Run the database handler in a goroutine
	database.Init(conf)

	// Run the bot
	// disc.BotRun(conf)

	fmt.Println(database.OwnsCharacter(
		206794847581896705,
		2010,
	))
}
