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
	VerifyChar(context.Context, corde.Snowflake, int64) (Character, error)
	CharsIDs(ctx context.Context, userID corde.Snowflake) ([]int64, error)
	DeleteChar(context.Context, corde.Snowflake, int64) (Character, error)
	CharsStartingWith(context.Context, corde.Snowflake, string) ([]Character, error)
	User(context.Context, corde.Snowflake) (User, error)
	Profile(context.Context, corde.Snowflake) (Profile, error)
	SetUserDate(context.Context, corde.Snowflake, time.Time) error
	SetUserFavorite(context.Context, corde.Snowflake, int64) error
	SetUserQuote(context.Context, corde.Snowflake, string) error
	GiveUserChar(ctx context.Context, dst corde.Snowflake, src corde.Snowflake, charID int64) error
	AddDropToken(context.Context, corde.Snowflake) error
	ConsumeDropTokens(context.Context, corde.Snowflake, int32) (User, error)
	Tx(fn func(s Store) error) error
}

// Interacter
type Interacter interface {
	GetInteractionCount(ctx context.Context, channelID corde.Snowflake) (int64, error)
	ResetInteractionCount(ctx context.Context, channelID corde.Snowflake) error
	IncrementInteractionCount(ctx context.Context, channelID corde.Snowflake) error

	SetChannelChar(ctx context.Context, channelID corde.Snowflake, char MediaCharacter) error
	GetChannelChar(ctx context.Context, channelID corde.Snowflake) (MediaCharacter, error)
	RemoveChannelChar(ctx context.Context, channelID corde.Snowflake) error
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
	mux               *corde.Mux
	Store             Store
	AnimeService      TrackingService
	Inter             Interacter
	AppID             corde.Snowflake
	GuildID           *corde.Snowflake
	BotToken          string
	PublicKey         string
	RollCooldown      time.Duration
	InteractionNeeded int64
	TokensNeeded      int32
}

// New runs the bot
func New(b *Bot) *corde.Mux {
	b.mux = corde.NewMux(b.PublicKey, b.AppID, b.BotToken)
	b.mux.OnNotFound = b.RemoveUnknownCommands

	traceSlash := trace[corde.SlashCommandInteractionData]
	interacterSlash := interact(b.Inter, onInteraction[corde.SlashCommandInteractionData](b))

	b.mux.Route("give", b.give)
	b.mux.Route("search", b.search)
	b.mux.Route("profile", b.profile)
	b.mux.Route("verify", b.verify)
	b.mux.Route("exchange", b.exchange)
	b.mux.SlashCommand("list", wrap(b.list, traceSlash, interacterSlash))
	b.mux.SlashCommand("roll", wrap(b.roll, traceSlash, interacterSlash))
	b.mux.SlashCommand("info", wrap(b.info, traceSlash))
	b.mux.SlashCommand("claim", wrap(b.claim, traceSlash))

	return b.mux
}

func onInteraction[T corde.InteractionDataConstraint](b *Bot) func(count int64, i *corde.Request[T]) bool {
	return func(count int64, i *corde.Request[T]) bool {
		if count < b.InteractionNeeded {
			return false
		}

		if b.GuildID != nil && *b.GuildID != i.GuildID {
			return true
		}

		b.drop(i.Context, i.GuildID, i.ChannelID)
		return true
	}
}

// interaction middleware
func interact[T corde.InteractionDataConstraint](inter Interacter, resetCount func(count int64, i *corde.Request[T]) bool) func(func(w corde.ResponseWriter, i *corde.Request[T])) func(w corde.ResponseWriter, i *corde.Request[T]) {
	return func(next func(w corde.ResponseWriter, i *corde.Request[T])) func(w corde.ResponseWriter, i *corde.Request[T]) {
		return func(w corde.ResponseWriter, i *corde.Request[T]) {
			defer next(w, i)

			err := inter.IncrementInteractionCount(i.Context, i.ChannelID)
			if err != nil {
				log.Debug().Err(err).Msg("failed to increment interaction count")
			}

			count, err := inter.GetInteractionCount(i.Context, i.ChannelID)
			if err != nil {
				log.Err(err).Msg("failed to get interaction count")
				return
			}

			if resetCount(count, i) {
				err = inter.ResetInteractionCount(i.Context, i.ChannelID)
				if err != nil {
					log.Err(err).Msg("failed to reset interaction count")
				}
				return
			}
		}
	}
}

func wrap[T corde.InteractionDataConstraint](
	next func(w corde.ResponseWriter, i *corde.Request[T]),
	fns ...func(func(w corde.ResponseWriter, i *corde.Request[T])) func(w corde.ResponseWriter, i *corde.Request[T]),
) func(w corde.ResponseWriter, i *corde.Request[T]) {
	// apply middleware in reverse order
	for i := len(fns) - 1; i >= 0; i-- {
		next = fns[i](next)
	}
	return next
}

func (b *Bot) RemoveUnknownCommands(r corde.ResponseWriter, i *corde.Request[corde.JsonRaw]) {
	log.Error().Str("command", i.Route).Int("type", int(i.Type)).Msg("Unknown command")
	r.Respond(corde.NewResp().Content("I don't know what that means, you shouldn't be able to do that").Ephemeral())

	var opt []func(*corde.CommandsOpt)
	if b.GuildID != nil {
		opt = append(opt, corde.GuildOpt(*b.GuildID))
	}

	b.mux.DeleteCommand(i.ID, opt...)
}
