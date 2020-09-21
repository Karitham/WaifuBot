package disc

import (
	"errors"
	"fmt"

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
	sub.ChangeCommandInfo("User", "", "search for an anilist user")
}

// Manga is a subcommand of Search
func (s *Search) Manga(_ *gateway.MessageCreateEvent, name bot.RawArguments) (*discord.Embed, error) {
	if name == "" {
		return nil, errors.New("missing manga name")
	}

	r, err := query.MediaSearch(string(name), "MANGA")
	if err != nil {
		return nil, err
	}

	return &discord.Embed{
		Title:       r.Media.Title.Romaji,
		URL:         r.Media.SiteURL,
		Description: r.Media.Description,
		Thumbnail: &discord.EmbedThumbnail{
			URL: r.Media.CoverImage.Medium,
		},
	}, nil
}

// Anime is a subcommand of Search
func (s *Search) Anime(_ *gateway.MessageCreateEvent, name bot.RawArguments) (*discord.Embed, error) {
	if name == "" {
		return nil, errors.New("missing anime name")
	}

	r, err := query.MediaSearch(string(name), "ANIME")
	if err != nil {
		return nil, err
	}

	return &discord.Embed{
		Title:       r.Media.Title.Romaji,
		URL:         r.Media.SiteURL,
		Description: r.Media.Description,
		Thumbnail: &discord.EmbedThumbnail{
			URL: r.Media.CoverImage.Medium,
		},
	}, nil
}

// Character is a subcommand of Search
func (s *Search) Character(_ *gateway.MessageCreateEvent, name bot.RawArguments) (*discord.Embed, error) {
	if name == "" {
		return nil, errors.New("missing character name / ID")
	}

	// Parse args
	n, id := parseArgs(name)
	searchArgs := query.CharSearchInput{
		ID:   id,
		Name: n,
	}

	// Search for character
	r, err := query.CharSearch(searchArgs)
	if err != nil {
		return nil, err
	}

	return &discord.Embed{
		Title: r.Character.Name.Full,
		URL:   r.Character.SiteURL,
		Description: fmt.Sprintf(
			"Found character ID `%d`\nThis character appears in:\n- %s",
			r.Character.ID,
			r.Character.Media.Nodes[0].Title.Romaji,
		),
		Thumbnail: &discord.EmbedThumbnail{
			URL: r.Character.Image.Large,
		},
	}, nil
}

// User is a subcommand of Search
func (s *Search) User(_ *gateway.MessageCreateEvent, name bot.RawArguments) (*discord.Embed, error) {
	if name == "" {
		return nil, errors.New("missing user name")
	}

	r, err := query.User(string(name))
	if err != nil {
		return nil, err
	}

	return &discord.Embed{
		Title:       r.User.Name,
		URL:         r.User.SiteURL,
		Description: r.User.About,
		Thumbnail:   &discord.EmbedThumbnail{URL: r.User.Avatar.Medium},
		Footer: &discord.EmbedFooter{
			Text: fmt.Sprintf(
				"Chapters read : %d | Episode watched : %d",
				r.User.Statistics.Manga.ChaptersRead,
				r.User.Statistics.Anime.EpisodesWatched,
			),
		},
	}, nil
}
