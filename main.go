package main

import (
	"flag"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/Karitham/WaifuBot/internal/config"
	"github.com/Karitham/WaifuBot/internal/db"
	"github.com/Karitham/WaifuBot/internal/disc"
)

var configFile string

func init() {
	flag.StringVar(&configFile, "config", "config.toml", "used to set the config file on start")
	flag.Parse()
}

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	log.Logger = log.Level(zerolog.TraceLevel)

	// Retrieve config and start the bot
	conf, err := config.Retrieve(configFile)
	if err != nil {
		log.Fatal().Err(err).Msg("Error getting config")
	}

	conn, err := db.Init(conf.Database)
	if err != nil {
		log.Fatal().Err(err).Msg("couldn't connect to db")
	}

	disc.Start(conf, conn)
}
