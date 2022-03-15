package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/Karitham/WaifuBot/internal/anilist"
	"github.com/Karitham/WaifuBot/internal/db"
	"github.com/Karitham/WaifuBot/internal/discord"
	"github.com/Karitham/corde"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/urfave/cli/v2"
)

func main() {
	disc := discordCmd{}
	d := dbCmd{}
	dev := false

	app := &cli.App{
		Name:        "waifubot",
		Usage:       "Run the bot, and use utils",
		Version:     "v0.7.2",
		Description: "A discord gacha bot",
		Commands: []*cli.Command{
			{
				Name:    "register",
				Aliases: []string{"r"},
				Usage:   "Register the bot commands",
				Action:  disc.register,
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:        "BOT_TOKEN",
						EnvVars:     []string{"DISCORD_TOKEN", "BOT_TOKEN"},
						Required:    true,
						Destination: &disc.botToken,
					},
					&cliSnowflake{
						EnvVars:  []string{"DISCORD_GUILD_ID", "GUILD_ID"},
						Dest:     disc.guildID,
						Required: true,
					},
					&cliSnowflake{
						EnvVars:  []string{"DISCORD_APP_ID", "APP_ID"},
						Dest:     &disc.appID,
						Required: true,
					},
				},
			},
			{
				Name:    "update-character",
				Usage:   "Update the character in the database",
				Aliases: []string{"uc"},
				Action:  d.update,
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:        "DB_URL",
						EnvVars:     []string{"DB_STR", "DB_URL"},
						Destination: &d.dbURL,
						Required:    true,
					},
				},
			},
		},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "BOT_TOKEN",
				EnvVars:     []string{"DISCORD_TOKEN", "BOT_TOKEN"},
				Required:    true,
				Destination: &disc.botToken,
			},
			&cliSnowflake{
				EnvVars: []string{"DISCORD_GUILD_ID", "GUILD_ID"},
				Dest:    disc.guildID,
			},
			&cliSnowflake{
				EnvVars: []string{"DISCORD_APP_ID", "APP_ID"},
				Dest:    &disc.appID,
			},
			&cli.StringFlag{
				Name:        "PUBLIC_KEY",
				EnvVars:     []string{"DISCORD_PUBLIC_KEY", "PUBLIC_KEY"},
				Destination: &disc.publicKey,
			},
			&cli.DurationFlag{
				Name:        "ROLL_COOLDOWN",
				EnvVars:     []string{"ROLL_TIMEOUT", "ROLL_COOLDOWN"},
				Value:       time.Hour * 2,
				Destination: &disc.rollCooldown,
			},
			&cli.IntFlag{
				Name:        "TOKENS_NEEDED",
				EnvVars:     []string{"TOKENS_NEEDED"},
				Value:       3,
				Destination: &disc.tokensNeeded,
			},
			&cli.StringFlag{
				Name:        "DB_URL",
				EnvVars:     []string{"DB_STR", "DB_URL"},
				Destination: &disc.dbURL,
			},
			&cli.StringFlag{
				Name:        "PORT",
				EnvVars:     []string{"PORT"},
				Value:       "8080",
				Destination: &disc.port,
			},
			&cli.Int64Flag{
				Name:    "ANILIST_MAX_CHARS",
				Value:   30_000,
				EnvVars: []string{"ANILIST_MAX_CHARS"},
			},
			&cli.BoolFlag{
				Name:        "DEV",
				EnvVars:     []string{"DEV"},
				Destination: &dev,
			},
		},
		Action: disc.run,
		Before: func(*cli.Context) error {
			if dev {
				log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
			}
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal().Err(err).Msg("error running app")
	}
}

type discordCmd struct {
	botToken        string
	appID           corde.Snowflake
	guildID         *corde.Snowflake
	publicKey       string
	anilistMaxChars int64
	rollCooldown    time.Duration
	tokensNeeded    int
	dbURL           string
	port            string
}

func (d *discordCmd) register(c *cli.Context) error {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	bot := &discord.Bot{
		AppID:    d.appID,
		BotToken: d.botToken,
		GuildID:  d.guildID,
	}

	if err := bot.RegisterCommands(); err != nil {
		return fmt.Errorf("error registering commands %v", err)
	}
	return nil
}

func (r *discordCmd) run(c *cli.Context) error {
	log.Logger = log.Level(zerolog.TraceLevel)

	db, err := db.NewDB(r.dbURL)
	if err != nil {
		return fmt.Errorf("error connecting to db %v", err)
	}

	bot := &discord.Bot{
		Store:        db,
		AnimeService: anilist.New(),
		AppID:        r.appID,
		GuildID:      r.guildID,
		BotToken:     r.botToken,
		PublicKey:    r.publicKey,
		RollCooldown: r.rollCooldown,
		TokensNeeded: int32(r.tokensNeeded),
	}

	return discord.New(bot).ListenAndServe(":" + r.port)
}

type dbCmd struct {
	dbURL string
}

func (r *dbCmd) update(c *cli.Context) error {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	a := c.Args()
	if a.Len() < 1 {
		return fmt.Errorf("no character name provided")
	}

	DB, err := db.NewDB(r.dbURL)
	if err != nil {
		return fmt.Errorf("error connecting to db %v", err)
	}

	char, err := anilist.New(anilist.NoCache).Character(c.Context, c.Args().First())
	if err != nil {
		return err
	}
	if len(char) < 1 {
		return fmt.Errorf("character not found")
	}

	if _, err := DB.SetChar(c.Context, db.SetCharParams{
		Image: char[0].ImageURL,
		Name:  strings.Join(strings.Fields(char[0].Name), " "),
		ID:    char[0].ID,
	}); err != nil {
		return fmt.Errorf("error updating db %v", err)
	}

	return nil
}
