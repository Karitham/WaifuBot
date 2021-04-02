package main

import (
	"flag"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/Karitham/WaifuBot/config"
	"github.com/Karitham/WaifuBot/db"
	"github.com/Karitham/WaifuBot/disc"
)

var configFile string

func init() {
	flag.StringVar(&configFile, "config", "config.toml", "used to set the config file on start")
	flag.Parse()
}

func main() {
	// Retrieve config
	conf, err := config.Retrieve(configFile)
	if err != nil {
		log.Fatal().Err(err).Msg("Error getting config")
	}

	// Setup logging
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	log.Logger = log.Level(conf.LoggingLevel)

	// Setup db
	conn, err := db.Init(conf.Database)
	if err != nil {
		log.Fatal().Err(err).Msg("Couldn't connect to db")
	}

	// Run the bot
	waitFn, err := disc.Start(conf, conn)
	if err != nil {
		log.Fatal().Err(err).Msg("Error starting the bot")
	}

	log.Info().Msg("Bot started")

	if err := waitFn(); err != nil {
		log.Fatal().
			Err(err).
			Msg("Error on keeping the bot alive")
	}
}
