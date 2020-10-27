package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path"
	"time"

	"github.com/Karitham/WaifuBot/config"
	"github.com/Karitham/WaifuBot/database"
	"github.com/Karitham/WaifuBot/disc"
)

func main() {
	var configFile, logDir string

	// Set flags
	flag.StringVar(&configFile, "config", "config.toml", "used to set the config file on start")
	flag.StringVar(&logDir, "log", "logs", "used to set the logging folder")
	flag.Parse()

	// Retrieve config and start the bot
	c := config.Retrieve(configFile)

	// Set up logging
	err := os.Mkdir(logDir, 0666)
	if err != nil {
		log.Println(err)
	}

	logFile := path.Join(logDir, fmt.Sprintf("%s.txt", time.Now().Format("2006-01-02 15h04m")))

	lf, err := os.Create(logFile)
	if err != nil {
		fmt.Println(err)
	}
	defer lf.Close()

	log.SetOutput(lf)

	database.Init(c)
	disc.Start(c)
}
