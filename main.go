package main

import (
	"os"

	"github.com/Karitham/WaifuBot/internal/anilist"
	"github.com/Karitham/WaifuBot/internal/discord"
	"github.com/Karitham/WaifuBot/internal/filestore"
	"github.com/Karitham/corde"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	log.Logger = log.Output(zerolog.NewConsoleWriter())
	log.Logger = log.Output(os.Stderr)
	log.Logger = log.Level(zerolog.TraceLevel)

	//nolint:errcheck
	godotenv.Load()
	token := os.Getenv("BOT_TOKEN")
	pk := os.Getenv("PUBLIC_KEY")
	port := os.Getenv("PORT")
	appID := corde.SnowflakeFromString(os.Getenv("APP_ID"))
	guildID := corde.SnowflakeFromString(os.Getenv("GUILD_ID"))
	_, ok := os.LookupEnv("FORCE_REGISTER_CMD")

	var store discord.Store
	if s := os.Getenv("STORE"); s != "FALSE" {
		fs := filestore.New("waifus.db")
		defer fs.Close()
		store = fs
	}

	bot := &discord.Bot{
		Store:            store,
		AnimeService:     anilist.New(),
		AppID:            appID,
		BotToken:         token,
		PublicKey:        pk,
		ForceRegisterCMD: ok,
		GuildID:          guildID,
	}

	log.Info().Str("PORT", port).Msg("starting bot")
	if err := discord.New(bot).ListenAndServe(":" + port); err != nil {
		log.Fatal().Err(err).Msg("error running bot")
	}
}
