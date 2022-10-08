package discord

import (
	"time"

	"github.com/Karitham/corde"
	"github.com/rs/zerolog/log"
)

func trace[T corde.InteractionDataConstraint](next func(w corde.ResponseWriter, i *corde.Request[T])) func(w corde.ResponseWriter, i *corde.Request[T]) {
	return func(w corde.ResponseWriter, i *corde.Request[T]) {
		start := time.Now()
		l := log.With().
			Str("route", i.Route).
			Stringer("guild", i.GuildID).
			Stringer("channel", i.ChannelID).
			Stringer("user", i.Member.User.ID).
			Logger()

		i.Context = l.WithContext(i.Context)
		next(w, i)

		l.Debug().Str("took", time.Since(start).String()).Send()
	}
}
