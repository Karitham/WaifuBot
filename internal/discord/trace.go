package discord

import (
	"time"

	"github.com/Karitham/corde"
	"github.com/rs/zerolog/log"
)

func trace(next corde.Handler) corde.Handler {
	return func(w corde.ResponseWriter, i *corde.InteractionRequest) {
		start := time.Now()
		defer log.Trace().
			Str("route", i.Data.Name).
			Stringer("user", i.Member.User.ID).
			Int("type", int(i.Type)).
			Str("took", time.Since(start).String()).
			Send()

		next(w, i)
	}
}
