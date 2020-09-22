package disc

import (
	"errors"

	"github.com/Karitham/WaifuBot/query"
	"github.com/diamondburned/arikawa/bot"
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
func (s *Search) Manga(_ *gateway.MessageCreateEvent, name bot.RawArguments) (string, error) {
	if name == "" {
		return "", errors.New("missing manga name")
	}

	r, err := query.MediaSearch(string(name), "MANGA")
	if err != nil {
		return "", err
	}

	return r.Media.SiteURL, nil
}

// Anime is a subcommand of Search
func (s *Search) Anime(_ *gateway.MessageCreateEvent, name bot.RawArguments) (string, error) {
	if name == "" {
		return "", errors.New("missing anime name")
	}

	r, err := query.MediaSearch(string(name), "ANIME")
	if err != nil {
		return "", err
	}

	return r.Media.SiteURL, nil
}

// Character is a subcommand of Search
func (s *Search) Character(_ *gateway.MessageCreateEvent, name bot.RawArguments) (string, error) {
	if name == "" {
		return "", errors.New("missing character name / ID")
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
		return "", err
	}

	return r.Character.SiteURL, nil
}

// User is a subcommand of Search
func (s *Search) User(_ *gateway.MessageCreateEvent, name bot.RawArguments) (string, error) {
	if name == "" {
		return "", errors.New("missing user name")
	}

	r, err := query.User(string(name))
	if err != nil {
		return "", err
	}

	return r.User.SiteURL, nil
}
