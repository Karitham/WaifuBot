package discord

import (
	"github.com/Karitham/corde"
	"github.com/Karitham/corde/components"
	"github.com/rs/zerolog/log"
)

func (b *Bot) give(m *corde.Mux) {
	m.SlashCommand("", trace(b.giveCommand))
	m.Autocomplete("id", b.profileEditFavoriteComplete)
}

func (b *Bot) giveCommand(w corde.ResponseWriter, i *corde.Request[components.SlashCommandInteractionData]) {
	user, errUserOK := i.Data.OptionsUser("user")
	if errUserOK != nil {
		w.Respond(rspErr("select a user to give to"))
		return
	}
	charID, errCharOK := i.Data.Options.Int("id")
	if errCharOK != nil {
		w.Respond(rspErr("select a character to give"))
		return
	}
	log.Trace().Stringer("src", i.Member.User.ID).Stringer("dst", user.ID).Int("charID", charID).Send()

	err := b.Store.GiveUserChar(i.Context, user.ID, i.Member.User.ID, int64(charID))
	if err != nil {
		w.Respond(newErrf("error giving character %d to user %s", charID, user.Username))
		return
	}

	w.Respond(components.NewResp().Contentf("You successfully gave %d to %s", charID, user.Username))
}