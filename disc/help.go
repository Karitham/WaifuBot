package disc

import (
	"fmt"

	"github.com/diamondburned/arikawa/discord"
	"github.com/diamondburned/arikawa/gateway"
)

// Help prints the default help message.
func (bot *Bot) Help(_ *gateway.MessageCreateEvent) (*discord.Embed, error) {
	return &discord.Embed{
		Title: "Help",
		Description: fmt.Sprintf(
			"Available functions :\n\n%v\n",
			bot.Ctx.Help(),
		),
		Footer: &discord.EmbedFooter{
			Text: "https://github.com/Karitham/WaifuBot",
			Icon: "https://upload.wikimedia.org/wikipedia/commons/thumb/9/91/Octicons-mark-github.svg/1200px-Octicons-mark-github.svg.png",
		},
	}, nil
}
