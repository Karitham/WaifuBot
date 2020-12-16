package main

import (
	"flag"
	"log"

	"github.com/Karitham/WaifuBot/config"
	"github.com/Karitham/WaifuBot/database"
	"github.com/Karitham/WaifuBot/disc"
)

var configFile string

func init() {
	flag.StringVar(&configFile, "config", "config.toml", "used to set the config file on start")
	flag.Parse()
}

func main() {
	// Retrieve config and start the bot
	conf := config.Retrieve(configFile)

	log.SetPrefix("[WaifuBot] ")

	database.Init(&conf)
	disc.Start(&conf)
}
