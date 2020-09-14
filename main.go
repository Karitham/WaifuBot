package main

import (
	"github.com/Karitham/WaifuBot/config"
	"github.com/Karitham/WaifuBot/database"
	"github.com/Karitham/WaifuBot/discord"
)

func main() {
	c := config.Retrieve("config.yml")
	database.Init(c)
	discord.BotRun(c)
}
