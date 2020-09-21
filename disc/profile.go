package disc

import (
	"fmt"
	"time"

	"github.com/Karitham/WaifuBot/database"
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
			uData.Quote, m.Author.Username, time.Now().Sub(uData.Date).Truncate(time.Minute), len(uData.Waifus), uData.ClaimedWaifus, uData.Favorite.Name,
		),
		Thumbnail: &discord.EmbedThumbnail{URL: uData.Favorite.Image},
	}, nil
}
