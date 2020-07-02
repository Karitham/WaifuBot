package main

import (
	"bot/data"
	"time"
)

func main() {
	/* const filename = "/config.json"
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	path := dir + filename
	disc.BotRun(path) */
	data.InitDB()

	Karitham := data.UserBson{
		UserID: 206794847581896705,
		Date:   time.Now(),
		Waifus: []int{4},
	}
	data.Store(Karitham)

}
