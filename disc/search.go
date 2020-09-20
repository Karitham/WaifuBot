package disc

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/Karitham/WaifuBot/query"
	"github.com/diamondburned/arikawa/bot"
	"github.com/diamondburned/arikawa/discord"
	"github.com/diamondburned/arikawa/gateway"
)

// Search Initializes a subcommand
type Search struct {
	Ctx *bot.Context
}

// Setup setups the subcommand
func (s *Search) Setup(sub *bot.Subcommand) {
	sub.Command = "search"
	sub.Description = "Search for characters, anime and manga"

	sub.ChangeCommandInfo("Manga", "", "search for a manga")
	sub.ChangeCommandInfo("Anime", "", "search for an anime")
	sub.ChangeCommandInfo("Character", "", "search for a character")
}

// Manga is a subcommand of Search
func (s Search) Manga(_ *gateway.MessageCreateEvent, name string) (*discord.Embed, error) {
	if name == "" {
		return nil, errors.New("missing manga name")
	}
	q, err := query.MediaSearch(name, "MANGA")
	if err != nil {
		return nil, err
	}
	return &discord.Embed{
		Title:       q.Media.Title.Romaji,
		URL:         q.Media.SiteURL,
		Description: q.Media.Description,
		Color:       discord.Color(0x01663be),
		Thumbnail: &discord.EmbedThumbnail{
			URL: q.Media.CoverImage.Medium,
		},
	}, nil
}

// Anime is a subcommand of Search
func (s Search) Anime(_ *gateway.MessageCreateEvent, name string) (*discord.Embed, error) {
	if name == "" {
		return nil, errors.New("missing anime name")
	}
	q, err := query.MediaSearch(name, "ANIME")
	if err != nil {
		return nil, err
	}
	return &discord.Embed{
		Title:       q.Media.Title.Romaji,
		URL:         q.Media.SiteURL,
		Description: q.Media.Description,
		Color:       discord.Color(0x01663be),
		Thumbnail: &discord.EmbedThumbnail{
			URL: q.Media.CoverImage.Medium,
		},
	}, nil
}

// Character is a subcommand of Search
func (s Search) Character(_ *gateway.MessageCreateEvent, name string) (*discord.Embed, error) {
	if name == "" {
		return nil, errors.New("missing character name / ID")
	}

	var searchArgs = query.CharSearchInput{
		Name: name,
	}

	if id, err := strconv.Atoi(name); id != 0 && err == nil {
		searchArgs.ID = id
	}

	q, err := query.CharSearch(searchArgs)
	if err != nil {
		return nil, err
	}
	return &discord.Embed{
		Title: q.Character.Name.Full,
		URL:   q.Character.SiteURL,
		Description: fmt.Sprintf(
			"Found character ID `%d`\nThis character appears in:\n- %s",
			q.Character.ID,
			q.Character.Media.Nodes[0].Title.Romaji,
		),
		Color: discord.Color(0x01663be),
		Thumbnail: &discord.EmbedThumbnail{
			URL: q.Character.Image.Large,
		},
	}, nil
}
