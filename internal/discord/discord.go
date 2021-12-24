package discord

import (
	"github.com/Karitham/corde"
	"github.com/rs/zerolog/log"
)

type Store interface {
	charStore
}

type AnimeService interface {
	randomer
	animeSearcher
	charSearcher
	mangaSearcher
	userSearcher
}

type Bot struct {
	mux          *corde.Mux
	Store        Store
	AnimeService AnimeService
	AppID        corde.Snowflake
	GuildID      corde.Snowflake
	BotToken     string
	PublicKey    string
}

// Run runs the bot
func Run(b *Bot) {
	mux := corde.NewMux(b.PublicKey, b.AppID, b.BotToken)

	for _, c := range commands {
		err := mux.RegisterCommand(c, corde.Guild(b.GuildID))
		if err != nil {
			log.Err(err).Msg("error creating command")
		}
	}

	mux.SetRoute(appCommand("roll"), b.Roller)
	mux.SetRoute(appCommand("search/char"), b.SearchChar)
	mux.SetRoute(appCommand("search/user"), b.SearchUser)
	mux.SetRoute(appCommand("search/manga"), b.SearchManga)
	mux.SetRoute(appCommand("search/anime"), b.SearchAnime)
	mux.SetRoute(appCommand("list"), b.List)

	log.Info().Msg("Gateway connected")
	mux.ListenAndServe(":8070")
}

func appCommand(route string) corde.InteractionCommand {
	return corde.InteractionCommand{Type: corde.APPLICATION_COMMAND, Route: route}
}
