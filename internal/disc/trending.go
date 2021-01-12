package disc

import (
	"errors"
	"strings"

	"github.com/Karitham/WaifuBot/internal/anilist"
	"github.com/diamondburned/arikawa/v2/bot"
	"github.com/diamondburned/arikawa/v2/discord"
	"github.com/diamondburned/arikawa/v2/gateway"
	"github.com/rs/zerolog/log"
)

// Trending Initializes a subcommand
type Trending struct {
	Ctx *bot.Context
}

// Setup setups the subcommand
func (s *Trending) Setup(sub *bot.Subcommand) {
	sub.Command = "trending"
	sub.Description = "View trending manga and anime"

	sub.ChangeCommandInfo("Manga", "", "search for a manga")
	sub.ChangeCommandInfo("Anime", "", "search for an anime")
}

// Manga is a subcommand of Search
func (s *Trending) Manga(_ *gateway.MessageCreateEvent) (*discord.Embed, error) {
	r, err := anilist.TrendingMediaQuery("MANGA")
	if err != nil {
		log.Err(err).
			Str("Type", "TRENDING MANGA").
			Msg("Could not query trending manga")

		return nil, errors.New("couldn't query trending manga")
	}

	return &discord.Embed{
		Title:       "Trending Manga",
		Description: formatTrending(r),
	}, nil
}

// Anime is a subcommand of Search
func (s *Trending) Anime(_ *gateway.MessageCreateEvent) (*discord.Embed, error) {
	r, err := anilist.TrendingMediaQuery("ANIME")
	if err != nil {
		log.Err(err).
			Str("Type", "TRENDING ANIME").
			Msg("Could not query trending anime")

		return nil, errors.New("couldn't query trending anime")
	}

	return &discord.Embed{
		Title:       "Trending Anime",
		Description: formatTrending(r),
	}, nil
}

func formatTrending(list anilist.TrendingMediaStruct) string {
	sb := new(strings.Builder)

	for _, v := range list.Page.Media {
		sb.WriteString(v.Title.Romaji)
		sb.WriteString("\n")
	}

	return sb.String()
}
