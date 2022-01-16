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
	User(context.Context, corde.Snowflake) (User, error)
	Profile(context.Context, corde.Snowflake) (Profile, error)
	SetUserDate(context.Context, corde.Snowflake, time.Time) error
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

	b.mux.Route("search", b.search)
	b.mux.Command("list", trace(b.list))
	b.mux.Command("roll", trace(b.roll))
	b.mux.Command("profile", trace(b.profile))

	return b.mux
}

func (b *Bot) registerCommands() error {
	actual := []corde.CreateCommand{searchCmd, rollCmd, listCmd, profileCmd}

	commands, err := b.mux.GetCommands(corde.GuildOpt(b.GuildID))
	if err != nil {
		log.Err(err).Msg("Failed to get commands")
	}

	if b.ForceRegisterCMD {
		var toRegister []corde.CreateCommander
		for _, c := range actual {
			toRegister = append(toRegister, c)
		}

		log.Info().Msg("Forcing register of CMD")
		return b.mux.BulkRegisterCommand(toRegister, corde.GuildOpt(b.GuildID))
	}

	for _, c := range commands {
		for i, r := range actual {
			if c.Name == r.Name {
				actual = remove(actual, i)
				break
			}
		}
	}

	if len(actual) != 0 {
		var toRegister []corde.CreateCommander
		for _, c := range actual {
			toRegister = append(toRegister, c)
		}

		return b.mux.BulkRegisterCommand(toRegister, corde.GuildOpt(b.GuildID))
	}

	return nil
}

func (b *Bot) RemoveUnknownCommands(r corde.ResponseWriter, i *corde.InteractionRequest) {
	r.Respond(corde.NewResp().Content("I don't know what that means, you shouldn't be able to do that").Ephemeral())
	b.mux.DeleteCommand(i.ID, corde.GuildOpt(b.GuildID))
}

func remove[T any](s []T, i int) []T {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

func trace(next corde.Handler) corde.Handler {
	return func(w corde.ResponseWriter, i *corde.InteractionRequest) {
		start := time.Now()
		defer log.Trace().Stringer("user", i.Member.User.ID).TimeDiff("took", time.Now(), start).Msg(i.Data.Name)

		next(w, i)
	}
}
