package disc

import (
	"github.com/diamondburned/arikawa/v2/discord"
	"github.com/diamondburned/arikawa/v2/gateway"
)

// Help prints the default help message.
func (b *Bot) Help(_ *gateway.MessageCreateEvent) (*discord.Embed, error) {
	return &discord.Embed{
		Description: b.Ctx.Help(),
		Footer: &discord.EmbedFooter{
			Text: "https://github.com/Karitham/WaifuBot",
			Icon: "https://upload.wikimedia.org/wikipedia/commons/thumb/9/91/Octicons-mark-github.svg/1200px-Octicons-mark-github.svg.png",
		},
	}, nil
}
