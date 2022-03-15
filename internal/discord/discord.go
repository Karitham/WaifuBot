package discord

import (
	"context"
	"time"

	"github.com/Karitham/corde"
	"github.com/Karitham/corde/components"
	"github.com/Karitham/corde/snowflake"
	"github.com/rs/zerolog/log"
)

const (
	AnilistColor   = 0x02a9ff
	AnilistIconURL = "https://anilist.co/img/icons/favicon-32x32.png"
)

// Store is the database
type Store interface {
	PutChar(context.Context, snowflake.Snowflake, Character) error
	Chars(context.Context, snowflake.Snowflake) ([]Character, error)
	VerifyChar(context.Context, snowflake.Snowflake, int64) (Character, error)
	CharsIDs(ctx context.Context, userID snowflake.Snowflake) ([]int64, error)
	DeleteChar(context.Context, snowflake.Snowflake, int64) (Character, error)
	CharsStartingWith(context.Context, snowflake.Snowflake, string) ([]Character, error)
	User(context.Context, snowflake.Snowflake) (User, error)
	Profile(context.Context, snowflake.Snowflake) (Profile, error)
	SetUserDate(context.Context, snowflake.Snowflake, time.Time) error
	SetUserFavorite(context.Context, snowflake.Snowflake, int64) error
	SetUserQuote(context.Context, snowflake.Snowflake, string) error
	GiveUserChar(ctx context.Context, dst snowflake.Snowflake, src snowflake.Snowflake, charID int64) error
	AddDropToken(context.Context, snowflake.Snowflake) error
	ConsumeDropTokens(context.Context, snowflake.Snowflake, int32) (User, error)
	Tx(fn func(s Store) error) error
}

// TrackingService is the interface for the anilist service
type TrackingService interface {
	RandomCharer
	AnimeSearcher
	CharSearcher
	MangaSearcher
	UserSearcher
}

// Bot holds the bot state
type Bot struct {
	mux          *corde.Mux
	Store        Store
	AnimeService TrackingService
	AppID        snowflake.Snowflake
	GuildID      *snowflake.Snowflake
	BotToken     string
	PublicKey    string
	RollCooldown time.Duration
	TokensNeeded int32
}

// New runs the bot
func New(b *Bot) *corde.Mux {
	b.mux = corde.NewMux(b.PublicKey, b.AppID, b.BotToken)
	b.mux.OnNotFound = b.RemoveUnknownCommands

	b.mux.Route("give", b.give)
	b.mux.Route("search", b.search)
	b.mux.Route("profile", b.profile)
	b.mux.Route("verify", b.verify)
	b.mux.Route("exchange", b.exchange)
	b.mux.SlashCommand("list", trace(b.list))
	b.mux.SlashCommand("roll", trace(b.roll))
	b.mux.SlashCommand("info", trace(b.info))

	return b.mux
}

func (b *Bot) RemoveUnknownCommands(r corde.ResponseWriter, i *corde.Request[components.JsonRaw]) {
	log.Error().Str("command", i.Route).Int("type", int(i.Type)).Msg("Unknown command")
	r.Respond(components.NewResp().Content("I don't know what that means, you shouldn't be able to do that").Ephemeral())

	var opt []func(*corde.CommandsOpt)
	if b.GuildID != nil {
		opt = append(opt, corde.GuildOpt(*b.GuildID))
	}

	b.mux.DeleteCommand(i.ID, opt...)
}
