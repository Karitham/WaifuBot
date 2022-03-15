package main

import (
	"os"
	"strconv"
	"time"

	"github.com/Karitham/WaifuBot/internal/anilist"
	"github.com/Karitham/WaifuBot/internal/db"
	"github.com/Karitham/WaifuBot/internal/discord"
	"github.com/Karitham/corde/snowflake"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	//nolint:errcheck
	godotenv.Load()

	log.Logger = log.Level(zerolog.TraceLevel)
	if os.Getenv("ENV") == "dev" {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}

	token := os.Getenv("BOT_TOKEN")
	publicKey := os.Getenv("PUBLIC_KEY")
	appID := snowflake.SnowflakeFromString(os.Getenv("APP_ID"))
	guildID := snowflake.SnowflakeFromString(os.Getenv("GUILD_ID"))

	rollCD, _ := time.ParseDuration(os.Getenv("ROLL_TIMEOUT"))
	if rollCD == 0 {
		rollCD = time.Minute * 5
	}

	tokensNeeded, _ := strconv.Atoi(os.Getenv("TOKENS_NEEDED"))
	if tokensNeeded == 0 {
		tokensNeeded = 3
	}

	bot := &discord.Bot{
		AppID:        appID,
		BotToken:     token,
		PublicKey:    publicKey,
		GuildID:      guildID,
		RollCooldown: rollCD,
		TokensNeeded: int32(tokensNeeded),
	}

	// register commands
	if len(os.Args) > 1 && os.Args[1] == "register" {
		err := bot.RegisterCommands()
		if err != nil {
			log.Err(err).Msg("failed to register commands")
		}

		log.Info().Msg("registered commands")
		return
	}

	store, err := db.NewDB(os.Getenv("DB_STR"))
	if err != nil {
		log.Fatal().Err(err).Msg("Error connecting to db")
	}

	bot.Store = store
	bot.AnimeService = anilist.New()

	port := os.Getenv("PORT")
	if err := discord.New(bot).ListenAndServe(":" + port); err != nil {
		log.Fatal().Err(err).Msg("error running bot")
	}
}
