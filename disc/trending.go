package disc

import (
	"errors"

	"github.com/Karitham/WaifuBot/query"
	"github.com/diamondburned/arikawa/v2/bot"
	"github.com/diamondburned/arikawa/v2/discord"
	"github.com/diamondburned/arikawa/v2/gateway"
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
	r, err := query.TrendingMediaQuery("MANGA")
	if err != nil {
		return nil, errors.New("couldn't query trending manga")
	}

	return &discord.Embed{
		Title:       "Trending Manga",
		Description: formatTrending(r),
	}, nil
}

// Anime is a subcommand of Search
func (s *Trending) Anime(_ *gateway.MessageCreateEvent) (*discord.Embed, error) {
	r, err := query.TrendingMediaQuery("ANIME")
	if err != nil {
		return nil, errors.New("couldn't query trending anime")
	}

	return &discord.Embed{
		Title:       "Trending Anime",
		Description: formatTrending(r),
	}, nil
}

func formatTrending(list query.TrendingMediaStruct) (formattedList string) {
	for _, v := range list.Page.Media {
		formattedList += v.Title.Romaji + "\n"
	}
	return
}
