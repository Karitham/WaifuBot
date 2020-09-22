package disc

import (
	"errors"
	"fmt"
	"time"

	"github.com/Karitham/WaifuBot/database"
	"github.com/Karitham/WaifuBot/query"
	"github.com/diamondburned/arikawa/bot"
	"github.com/diamondburned/arikawa/discord"
	"github.com/diamondburned/arikawa/gateway"
)

// Profile displays user profile
func (b *Bot) Profile(m *gateway.MessageCreateEvent) (*discord.Embed, error) {
	uData, err := database.ViewUserData(m.Author.ID)
	if err != nil {
		return nil, err
	}
	return &discord.Embed{
		Title: fmt.Sprintf("%s's profile", m.Author.Username),
		Description: fmt.Sprintf(
			"%s\n%s last rolled %s ago.\nThey have rolled %d waifus and claimed %d.\nFavorite waifu is %s",
			uData.Quote, m.Author.Username, time.Since(uData.Date).Truncate(time.Minute), len(uData.Waifus), uData.ClaimedWaifus, uData.Favorite.Name,
		),
		Thumbnail: &discord.EmbedThumbnail{URL: uData.Favorite.Image},
	}, nil
}

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

// Quote sets a quote on the user profile
func (b *Bot) Quote(m *gateway.MessageCreateEvent, name bot.RawArguments) (string, error) {
	if name == "" {
		return "", errors.New("no quote entered")
	}

	database.NewQuote{
		UserID: m.Author.ID,
		Quote:  string(name),
	}.SetQuote()

	return fmt.Sprintf("New quote set :\n%s", string(name)), nil
}
