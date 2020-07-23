package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"
)

// ConfJSONStruct is used to unmarshal the config.json
type ConfJSONStruct struct {
	Prefix              string        `json:"Prefix"`
	BotToken            string        `json:"Bot_Token"`
	MongoURL            string        `json:"Mongo_URL"`
	MaxChar             int           `json:"Max_Character_Roll"`
	TimeBetweenRolls    time.Duration `json:"Time_Between_Rolls"`
	DelIllegalRollAfter time.Duration `json:"Delete_Illegal_Roll_After"`
	DropsOnInteract     int           `json:"Drops_On_Interact"`
}

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
