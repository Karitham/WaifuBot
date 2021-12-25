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
	listState    listState
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
		err := mux.RegisterCommand(c, corde.GuildOpt(b.GuildID))
		if err != nil {
			log.Err(err).Msg("error creating command")
		}
	}

	mux.Command("roll", b.Roller)
	mux.Command("search/char", b.SearchChar)
	mux.Command("search/user", b.SearchUser)
	mux.Command("search/manga", b.SearchManga)
	mux.Command("search/anime", b.SearchAnime)
	mux.Command("list", b.List())
	mux.Button("list/back", b.listBack)
	mux.Button("list/next", b.listNext)

	log.Info().Msg("Gateway connected")
	mux.ListenAndServe(":8070")
}
