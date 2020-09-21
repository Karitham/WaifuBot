package main

import (
	"github.com/Karitham/WaifuBot/config"
	"github.com/Karitham/WaifuBot/database"
	"github.com/Karitham/WaifuBot/disc"
)

func main() {
	c := config.Retrieve("config.toml")
	database.Init(c)
	disc.Start(c)
}
