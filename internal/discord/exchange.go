package discord

import (
	"github.com/Karitham/corde"
)

func (b *Bot) exchange(m *corde.Mux) {
	m.SlashCommand("", trace(b.exchangeCommand))
	m.Autocomplete("id", b.profileEditFavoriteComplete)
}

func (b *Bot) exchangeCommand(w corde.ResponseWriter, i *corde.Request[corde.SlashCommandInteractionData]) {
	user := i.Member.User
	char, _ := i.Data.Options.Int64("id")
	var c Character

	if err := b.Store.Tx(func(s Store) error {
		var err error
		c, err = s.DeleteChar(i.Context, user.ID, char)
		if err != nil {
			w.Respond(newErrf("%s doesn't own this character; you can't exchange it!", user.Username))
			return err
		}

		return s.AddDropToken(i.Context, user.ID)
	}); err != nil {
		return
	}

	w.Respond(newErrf("Good job %s, you just exchanged %s for a token!", user.Username, c.Name))
}
