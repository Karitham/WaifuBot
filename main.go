package main

import (
	"bot/data"
	"time"
)

func main() {
	/* 	const filename = "/config.json"
	   	dir, err := os.Getwd()
	   	if err != nil {
	   		panic(err)
	   	}
	   	path := dir + filename
		   disc.BotRun(path) */
	data.InitDB()

	Karitham := data.UserBson{
		UserID: 20679484758189670,
		Date:   time.Now(),
		Waifus: []int{4},
	}
	data.Store(Karitham)

}
