package config

import (
	"log"
	"time"

	"github.com/jinzhu/configor"
)

// ConfStruct is used to unmarshal the config.json
type ConfStruct struct {
	Prefix                 string        `yaml:"Prefix"`
	BotToken               string        `yaml:"Bot_Token"`
	MongoURL               string        `yaml:"Mongo_URL"`
	BotStatus              string        `yaml:"Bot_Status"`
	MaxCharacterRoll       int           `yaml:"Max_Character_Roll"`
	MaxCharacterDrop       int           `yaml:"Max_Character_Drop"`
	DeleteIllegalRollAfter time.Duration `yaml:"Delete_Illegal_Roll_After"`
	DeleteWrongClaimAfter  time.Duration `yaml:"Delete_Wrong_Claim_After"`
	TimeBetweenRolls       time.Duration `yaml:"Time_Between_Rolls"`
	ListMaxUpdateTime      time.Duration `yaml:"List_Max_Update_Time"`
	DropsOnInteract        int           `yaml:"Drops_On_Interact"`
}

// Retrieve retrieves the config from the file
func Retrieve(filename string) (config ConfStruct) {
	err := configor.Load(&config, filename)
	if err != nil {
		log.Println(err)
	}
	config.DeleteIllegalRollAfter *= time.Minute
	config.DeleteWrongClaimAfter *= time.Minute
	config.ListMaxUpdateTime *= time.Minute
	config.TimeBetweenRolls *= time.Hour
	return config
}
