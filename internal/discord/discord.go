package discord

import (
	"context"
	"time"

	"github.com/Karitham/WaifuBot/internal/anilist"
	"github.com/Karitham/corde"
	"github.com/rs/zerolog/log"
)

// Store is the database
type Store interface {
	PutChar(context.Context, corde.Snowflake, Character) error
	Chars(context.Context, corde.Snowflake) ([]Character, error)
	CharsIDs(ctx context.Context, userID corde.Snowflake) ([]int64, error)
	CharsStartingWith(context.Context, corde.Snowflake, string) ([]Character, error)
	User(context.Context, corde.Snowflake) (User, error)
	Profile(context.Context, corde.Snowflake) (Profile, error)
	SetUserDate(context.Context, corde.Snowflake, time.Time) error
	SetUserFavorite(context.Context, corde.Snowflake, int64) error
	SetUserQuote(context.Context, corde.Snowflake, string) error
	GiveUserChar(ctx context.Context, dst corde.Snowflake, src corde.Snowflake, charID int64) error
	Tx(fn func(s Store) error) error
}

// Check that anilist actually implements the interface
var _ AnimeService = (*anilist.Anilist)(nil)

// AnimeService is the interface for the anilist service
type AnimeService interface {
	randomCharGetter
	animeSearcher
	charSearcher
	mangaSearcher
	userSearcher
}

// Bot holds the bot state
type Bot struct {
	ForceRegisterCMD bool
	mux              *corde.Mux
	Store            Store
	AnimeService     AnimeService
	AppID            corde.Snowflake
	GuildID          corde.Snowflake
	BotToken         string
	PublicKey        string
	RollTimeout      time.Duration
	TokensNeeded     int32
}

// New runs the bot
func New(b *Bot) *corde.Mux {
	b.mux = corde.NewMux(b.PublicKey, b.AppID, b.BotToken)
	b.mux.OnNotFound = b.RemoveUnknownCommands

	if err := b.registerCommands(); err != nil {
		log.Err(err).Msg("failed to register commands")
	}

	b.mux.Route("give", b.give)
	b.mux.Route("search", b.search)
	b.mux.Route("profile", b.profile)
	b.mux.Command("list", trace(b.list))
	b.mux.Command("roll", trace(b.roll))
	b.mux.Command("info", trace(b.info))

	return b.mux
}
