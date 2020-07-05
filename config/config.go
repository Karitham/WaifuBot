package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// ConfJSONStruct is used to unmarshal the config.json
type ConfJSONStruct struct {
	Prefix   string `json:"Prefix"`
	BotToken string `json:"Bot_Token"`
	MongoURL string `json:"Mongo_URL"`
	MaxChar  int    `json:"Max_Character_Roll"`
}

var config ConfJSONStruct

// Retrieve reads config from file
func Retrieve(file string) ConfJSONStruct {
	var config ConfJSONStruct
	body, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println(err)
	}
	json.Unmarshal(body, &config)
	return config
}
