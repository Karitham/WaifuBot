package disc

import (
	"errors"
	"fmt"

	"github.com/Karitham/WaifuBot/database"
	"github.com/Karitham/WaifuBot/query"
	"github.com/diamondburned/arikawa/bot"
	"github.com/diamondburned/arikawa/gateway"
)

// Favorite sets a waifu as favorite
func (b *Bot) Favorite(m *gateway.MessageCreateEvent, name bot.RawArguments) (string, error) {
	if name == "" {
		return "", errors.New("no character name entered")
	}

	// Parse args
	n, id := parseArgs(name)
	searchArgs := query.CharSearchInput{
		ID:   id,
		Name: n,
	}

	// Search for character
	char, err := query.CharSearch(searchArgs)
	if err != nil {
		return "", err
	}

	database.FavoriteStruct{UserID: m.Author.ID, Favorite: database.CharLayout{
		ID:    char.Character.ID,
		Image: char.Character.Image.Large,
		Name:  char.Character.Name.Full,
	}}.SetFavorite()

	return fmt.Sprintf("New waifu set, check your profile\n<%s>", char.Character.SiteURL), nil
}
