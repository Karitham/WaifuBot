package main

import "bot/data"

func main() {
	/* 	const filename = "/config.json"
	   	dir, err := os.Getwd()
	   	if err != nil {
	   		panic(err)
	   	}
	   	path := dir + filename
		   disc.BotRun(path) */

	data.InitDB()
	data.ResetDB()
}
