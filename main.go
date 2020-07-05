package main

import (
	"bot/config"
	"bot/database"
	"bot/disc"
)

func main() {
	go database.Init(config.Retrieve("./config.json"))
	disc.BotRun(config.Retrieve("./config.json"))
}
