package disc

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/Karitham/WaifuBot/database"
	"github.com/Karitham/WaifuBot/query"
	"github.com/diamondburned/arikawa/bot"
	"github.com/diamondburned/arikawa/discord"
	"github.com/diamondburned/arikawa/gateway"
)

// Name represent the name of a character
type Name bot.RawArguments

// Quote represent a quote
type Quote bot.RawArguments

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
func (b *Bot) Favorite(m *gateway.MessageCreateEvent, name ...string) (string, error) {
	if name == nil {
		return "", errors.New("no character name entered")
	}
	n := strings.Join(name, " ")

	n, id := parseArgs(n)
	searchArgs := query.CharSearchInput{
		ID:   id,
		Name: n,
	}

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
func (b *Bot) Quote(m *gateway.MessageCreateEvent, quote ...string) (string, error) {
	if quote == nil {
		return "", errors.New("no quote entered")
	}

	q := strings.Join(quote, " ")

	database.NewQuote{
		UserID: m.Author.ID,
		Quote:  q,
	}.SetQuote()

	return fmt.Sprintf("New quote set :\n%s", q), nil
}