package discord

import (
	"time"

	"github.com/Karitham/corde"
	"github.com/Karitham/corde/components"
	"github.com/rs/zerolog/log"
)

func trace[T components.InteractionDataConstraint](
	next func(w corde.ResponseWriter, i *corde.Request[T]),
) func(w corde.ResponseWriter, i *corde.Request[T]) {
	return func(w corde.ResponseWriter, i *corde.Request[T]) {
		start := time.Now()
		l := log.With().Str("route", i.Route).
			Stringer("user", i.Member.User.ID).
			Int("type", int(i.Type)).Logger()

		i.Context = l.WithContext(i.Context)
		next(w, i)

		l.Trace().Str("took", time.Since(start).String()).Send()
	}
}
