package disc

import (
	"errors"
	"strings"

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
	sub.Description = "Search for characters, anime, manga and users"

	sub.AddAliases("Anime", "a", "A")
	sub.AddAliases("Manga", "m", "M")
	sub.AddAliases("Character", "c", "C")
	sub.AddAliases("User", "u", "U")

	sub.ChangeCommandInfo("Manga", "", "search for a manga")
	sub.ChangeCommandInfo("Anime", "", "search for an anime")
	sub.ChangeCommandInfo("Character", "", "search for a character")
	sub.ChangeCommandInfo("User", "", "search for an anilist user")
}

// Manga is a subcommand of Search
func (s *Search) Manga(_ *gateway.MessageCreateEvent, name ...Name) (string, error) {
	if len(name) < 1 {
		return "", errors.New("missing manga name")
	}

	n := strings.Join(name, " ")

	r, err := query.MediaSearch(n, "MANGA")
	if err != nil {
		return "", err
	}

	return r.Media.SiteURL, nil
}

// Anime is a subcommand of Search
func (s *Search) Anime(_ *gateway.MessageCreateEvent, name ...Name) (string, error) {
	if len(name) < 1 {
		return "", errors.New("missing anime name")
	}

	n := strings.Join(name, " ")

	r, err := query.MediaSearch(n, "ANIME")
	if err != nil {
		return "", err
	}

	return r.Media.SiteURL, nil
}

// Character is a subcommand of Search
func (s *Search) Character(_ *gateway.MessageCreateEvent, name ...Name) (string, error) {
	if len(name) < 1 {
		return "", errors.New("missing character name / ID")
	}

	n := strings.Join(name, " ")

	// Parse args
	id := parseArgs(n)
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
func (s *Search) User(_ *gateway.MessageCreateEvent, name ...Name) (string, error) {
	if len(name) < 1 {
		return "", errors.New("missing user name")
	}

	n := strings.Join(name, " ")

	r, err := query.User(n)
	if err != nil {
		return "", err
	}

	return r.User.SiteURL, nil
}
