package disc

import (
	"errors"
	"strings"

	"github.com/Karitham/WaifuBot/anilist"
	"github.com/diamondburned/arikawa/v2/bot"
	"github.com/diamondburned/arikawa/v2/gateway"
	"github.com/rs/zerolog/log"
)

// Name represent the name of a character
type Name = string

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

	r, err := anilist.MediaSearch(n, "MANGA")
	if err != nil {
		log.Debug().
			Err(err).
			Str("Title", n).
			Str("Type", "MANGA SEARCH").
			Msg("Could not get manga")

		return "", errors.New("error getting manga, please check the name entered is right and retry later")
	}

	return r.Media.SiteURL, nil
}

// Anime is a subcommand of Search
func (s *Search) Anime(_ *gateway.MessageCreateEvent, name ...Name) (string, error) {
	if len(name) < 1 {
		return "", errors.New("missing anime name")
	}

	n := strings.Join(name, " ")

	r, err := anilist.MediaSearch(n, "ANIME")
	if err != nil {
		log.Debug().
			Err(err).
			Str("Title", n).
			Str("Type", "ANIME SEARCH").
			Msg("Could not get anime")

		return "", errors.New("error getting anime, please check the name entered is right and retry later")
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
	searchArgs := anilist.CharSearchInput{
		ID:   id,
		Name: n,
	}

	// Search for character
	r, err := anilist.CharSearch(searchArgs)
	if err != nil {
		log.Debug().
			Err(err).
			Str("Name/ID", n).
			Str("Type", "CHAR SEARCH").
			Msg("Could not get character")

		return "", errors.New("error getting character, please check the name entered is right and retry later")
	}

	return r.Character.SiteURL, nil
}

// User is a subcommand of Search
func (s *Search) User(_ *gateway.MessageCreateEvent, name ...Name) (string, error) {
	if len(name) < 1 {
		return "", errors.New("missing user name")
	}

	n := strings.Join(name, " ")

	r, err := anilist.User(n)
	if err != nil {
		log.Debug().
			Err(err).
			Str("Name", n).
			Str("Type", "USER SEARCH").
			Msg("Could not get user")

		return "", errors.New("error getting user, please check the name entered is right and retry later")
	}

	return r.User.SiteURL, nil
}
