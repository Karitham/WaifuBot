package config

import (
	"time"

	"github.com/BurntSushi/toml"
)

// ConfStruct is used to unmarshal the config.toml
type ConfStruct struct {
	Prefix            []string `toml:"Prefix"`
	BotToken          string   `toml:"Bot_Token"`
	BotStatus         string   `toml:"Bot_Status"`
	MaxCharacterRoll  uint64   `toml:"Max_Character_Roll"`
	MaxCharacterDrop  uint     `toml:"Max_Character_Drop"`
	DropsOnInteract   uint64   `toml:"Drops_On_Interact"`
	ListLen           int      `toml:"List_Len"`
	ListMaxUpdateTime duration `toml:"List_Max_Update_Time"`
	TimeBetweenRolls  duration `toml:"Time_Between_Rolls"`
	Database          Database `toml:"Database"`
}

// Database represent the needed things for the database
type Database struct {
	Dbname   string `toml:"dbname"`
	Host     string `toml:"host"`
	Password string `toml:"password"`
	User     string `toml:"user"`
}

type duration struct {
	time.Duration
}

func (d *duration) UnmarshalText(text []byte) (err error) {
	d.Duration, err = time.ParseDuration(string(text))
	return
}

// Retrieve retrieves the config from the file
func Retrieve(filename string) (config *ConfStruct, err error) {
	if _, err = toml.DecodeFile(filename, &config); err != nil {
		return nil, err
	}
	return config, nil
}
