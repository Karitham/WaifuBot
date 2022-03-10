package discord

import (
	"github.com/Karitham/corde"
	"github.com/Karitham/corde/components"
)

func (b *Bot) verify(m *corde.Mux) {
	m.SlashCommand("", trace(b.verifyCommand))
	m.Autocomplete("id", b.profileEditFavoriteComplete)
}

func (b *Bot) verifyCommand(w corde.ResponseWriter, i *corde.Request[components.SlashCommandInteractionData]) {
	user := i.Member.User
	if len(i.Data.Resolved.Members) > 0 {
		user = i.Data.Resolved.Users.First()
	}
	charOpt, _ := i.Data.Options.Int64("id")

	char, _ := b.Store.VerifyChar(i.Context, user.ID, charOpt)
	if char.ID == charOpt {
		w.Respond(newErrf("%s owns %s", user.Username, char.Name))
		return
	}

	w.Respond(newErrf("%s doesn't own this character", user.Username))
}
