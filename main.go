package main

import (
	"os"
	"strconv"
	"time"

	"github.com/Karitham/WaifuBot/internal/anilist"
	"github.com/Karitham/WaifuBot/internal/db"
	"github.com/Karitham/WaifuBot/internal/discord"
	"github.com/Karitham/corde"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

//go:generate sqlc generate -f ./internal/sqlc.yaml

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	log.Logger = log.Level(zerolog.TraceLevel)

	//nolint:errcheck
	godotenv.Load()
	token := os.Getenv("BOT_TOKEN")
	publicKey := os.Getenv("PUBLIC_KEY")
	port := os.Getenv("PORT")
	appID := corde.SnowflakeFromString(os.Getenv("APP_ID"))
	guildID := corde.SnowflakeFromString(os.Getenv("GUILD_ID"))
	_, force := os.LookupEnv("FORCE_REGISTER_CMD")

	store, err := db.NewDB(os.Getenv("DB_STR"))
	if err != nil {
		log.Fatal().Err(err).Msg("Error connecting to db")
	}

	rollTimeout, _ := time.ParseDuration(os.Getenv("ROLL_TIMEOUT"))
	if rollTimeout == 0 {
		rollTimeout = time.Minute * 5
	}
	tokensNeeded, _ := strconv.Atoi(os.Getenv("TOKENS_NEEDED"))
	if tokensNeeded == 0 {
		tokensNeeded = 3
	}

	bot := &discord.Bot{
		Store:            store,
		AnimeService:     anilist.New(),
		AppID:            appID,
		BotToken:         token,
		PublicKey:        publicKey,
		ForceRegisterCMD: force,
		GuildID:          guildID,
		RollTimeout:      rollTimeout,
		TokensNeeded:     int32(tokensNeeded),
	}

	log.Info().Str("PORT", port).Msg("starting bot")
	if err := discord.New(bot).ListenAndServe(":" + port); err != nil {
		log.Fatal().Err(err).Msg("error running bot")
	}
}
