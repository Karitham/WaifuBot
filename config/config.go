package config

import (
	"strings"
	"time"

	"github.com/Netflix/go-env"
)

// ConfStruct is used to unmarshal the environ
type ConfStruct struct {
	DatabaseURL       string        `env:"DB_URL,required=true,default=postgres://db:5432/waifudb?sslmode=disable"`
	BotToken          string        `env:"TOKEN,required=true"`
	BotStatus         string        `env:"STATUS,defaul=use w.help for help"`
	Prefix            Prefixes      `env:"PREFIX,default=w."`
	MaxCharacterRoll  int           `env:"MAX_CHAR_ROLL,default=10000"`
	DropsOnInteract   int           `env:"INTERACT_DROPS,default=25"`
	ListLen           int           `env:"LIST_LEN,default=10"`
	ListMaxUpdateTime time.Duration `env:"LIST_UPDATE_TIME,default=5m"`
	TimeBetweenRolls  time.Duration `env:"ROLL_COOLDOWN,default=2h"`
	MaxCharacterDrop  int           `env:"MAX_CHAR_DROP,default=5000"`
	LoggingLevel      int           `env:"LOG_LEVEL,default=3"`
}

// Retrieve retrieves the config from the file
func Retrieve() (*ConfStruct, error) {
	var conf ConfStruct

	_, err := env.UnmarshalFromEnviron(&conf)
	if err != nil {
		return nil, err
	}

	return &conf, nil
}

type Prefixes []string

func (p *Prefixes) UnmarshalEnvironmentValue(data string) error {
	for _, s := range strings.Split(data, ",") {
		*p = append(*p, strings.TrimSpace(s))
	}
	return nil
}
