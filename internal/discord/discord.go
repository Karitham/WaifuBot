package discord

import (
	"context"
	"time"

	"github.com/Karitham/corde"
	"github.com/rs/zerolog/log"
)

const (
	AnilistColor   = 0x02a9ff
	AnilistIconURL = "https://anilist.co/img/icons/favicon-32x32.png"
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
	AppID        corde.Snowflake
	GuildID      corde.Snowflake
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
	b.mux.Command("list", trace(b.list))
	b.mux.Command("roll", trace(b.roll))
	b.mux.Command("info", trace(b.info))

	return b.mux
}

func (b *Bot) RemoveUnknownCommands(r corde.ResponseWriter, i *corde.InteractionRequest) {
	log.Error().Str("command", i.Data.Name).Int("type", int(i.Type)).Msg("Unknown command")
	r.Respond(corde.NewResp().Content("I don't know what that means, you shouldn't be able to do that").Ephemeral())
	b.mux.DeleteCommand(i.ID, corde.GuildOpt(b.GuildID))
}
