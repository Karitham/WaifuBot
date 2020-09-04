package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"time"
)

// ConfJSONStruct is used to unmarshal the config.json
type ConfJSONStruct struct {
	Prefix              string        `json:"Prefix"`
	BotToken            string        `json:"Bot_Token"`
	MongoURL            string        `json:"Mongo_URL"`
	Status              string        `json:"Bot_Status"`
	MaxCharRoll         int           `json:"Max_Character_Roll"`
	MaxCharDrop         int           `json:"Max_Character_Drop"`
	TimeBetweenRolls    time.Duration `json:"Time_Between_Rolls"`
	DelIllegalRollAfter time.Duration `json:"Delete_Illegal_Roll_After"`
	DelWrongClaimAfter  time.Duration `json:"Delete_Wrong_Claim_After"`
	DropsOnInteract     int           `json:"Drops_On_Interact"`
}

// Retrieve reads config from file
func Retrieve(file string) ConfJSONStruct {
	var config ConfJSONStruct
	body, err := ioutil.ReadFile(file)
	if err != nil {
		log.Println("error reading config file :", err)
	}

	err = json.Unmarshal(body, &config)
	if err != nil {
		log.Println("error unmarshalling config :", err)
	}
	return configTime(config)
}

// Configure message delete time
func configTime(conf ConfJSONStruct) ConfJSONStruct {
	conf.DelIllegalRollAfter *= time.Minute
	conf.DelWrongClaimAfter *= time.Minute
	return conf
}
