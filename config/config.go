package config

import (
	"os"
	"time"

	"github.com/pelletier/go-toml"
	"github.com/rs/zerolog"
)

// ConfStruct is used to unmarshal the config.toml
type ConfStruct struct {
	Database          Database      `toml:"Database"`
	BotToken          string        `toml:"Bot_Token"`
	BotStatus         string        `toml:"Bot_Status"`
	Prefix            []string      `toml:"Prefix"`
	MaxCharacterRoll  uint64        `toml:"Max_Character_Roll"`
	DropsOnInteract   uint64        `toml:"Drops_On_Interact"`
	ListLen           int           `toml:"List_Len"`
	ListMaxUpdateTime duration      `toml:"List_Max_Update_Time"`
	TimeBetweenRolls  duration      `toml:"Time_Between_Rolls"`
	MaxCharacterDrop  uint          `toml:"Max_Character_Drop"`
	LoggingLevel      zerolog.Level `toml:"Logging_Level"`
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
func Retrieve(filename string) (*ConfStruct, error) {
	var conf ConfStruct

	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	err = toml.NewDecoder(f).Decode(&conf)
	if err != nil {
		return nil, err
	}

	return &conf, nil
}
