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
	ForceRegisterCMD bool
	mux              *corde.Mux
	listState        listState
	Store            Store
	AnimeService     AnimeService
	AppID            corde.Snowflake
	GuildID          corde.Snowflake
	BotToken         string
	PublicKey        string
}

// New runs the bot
func New(b *Bot) *corde.Mux {
	m := corde.NewMux(b.PublicKey, b.AppID, b.BotToken)
	m.OnNotFound = b.RemoveUnknownCommands
	commands, err := m.GetCommands(corde.GuildOpt(b.GuildID))
	if err != nil {
		log.Err(err).Msg("Failed to get commands")
	}

	if err := registerCommands(b, m, commands, b.ForceRegisterCMD); err != nil {
		log.Err(err).Msg("failed to register commands")
	}

	m.Route("search", b.search)
	if b.Store != nil {
		m.Command("roll", b.roll)
		m.Route("list", b.list)
	}

	return m
}

func registerCommands(b *Bot, m *corde.Mux, commands []corde.Command, force bool) error {
	if force {
		return m.BulkRegisterCommand(
			[]corde.CreateCommander{
				searchCmd,
				listCmd,
				rollCmd,
			},
			corde.GuildOpt(b.GuildID),
		)
	}

	actual := []corde.CreateCommand{searchCmd}
	if b.Store != nil {
		actual = append(actual, rollCmd, listCmd)
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

		return m.BulkRegisterCommand(toRegister, corde.GuildOpt(b.GuildID))
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
