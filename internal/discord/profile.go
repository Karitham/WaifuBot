package discord

import (
	"fmt"
	"strconv"
	"time"

	"github.com/Karitham/corde"
	"github.com/Karitham/corde/components"
	"github.com/rs/zerolog/log"
)

type Profile struct {
	User
	CharacterCount int
	Favorite       Character
}

func (b *Bot) profile(m *corde.Mux) {
	m.SlashCommand("view", trace(b.profileView))
	m.Route("edit", func(m *corde.Mux) {
		m.SlashCommand("quote", trace(b.profileEditQuote))
		m.Route("favorite", func(m *corde.Mux) {
			m.SlashCommand("", trace(b.profileEditFavorite))
			m.Autocomplete("id", trace(b.profileEditFavoriteComplete))
		})
	})
}

func (b *Bot) profileView(w corde.ResponseWriter, i *corde.Request[components.SlashCommandInteractionData]) {
	user := i.Member.User
	if len(i.Data.Resolved.Users) > 0 {
		user = i.Data.Resolved.Users.First()
	}

	data, err := b.Store.Profile(i.Context, user.ID)
	if err != nil {
		log.Ctx(i.Context).Err(err).Msg("Error getting user's profile")
		w.Respond(components.NewResp().Content("An error occurred dialing the database, please try again later").Ephemeral())
		return
	}

	resp := components.NewEmbed().
		Title(user.Username).
		URL(fmt.Sprintf("https://waifugui.kar.moe/#/list/%s", user.ID.String())).
		Descriptionf(
			"%s\n%s last rolled %s ago and has %d tokens.\nThey have %d characters.\nTheir favorite character is %s",
			data.Quote,
			user.Username,
			time.Since(data.Date.UTC()).Truncate(time.Second),
			data.Tokens,
			data.CharacterCount,
			data.Favorite.Name,
		)
	if data.Favorite.Image != "" {
		resp.Thumbnail(components.Image{URL: data.Favorite.Image})
	}

	w.Respond(resp)
}

func (b *Bot) profileEditFavorite(w corde.ResponseWriter, i *corde.Request[components.SlashCommandInteractionData]) {
	optID, _ := i.Data.Options.Int64("id")
	err := b.Store.SetUserFavorite(i.Context, i.Member.User.ID, optID)
	if err != nil {
		log.Ctx(i.Context).Err(err).Stringer("user", i.Member.User.ID).Int64("character", optID).Msg("Error setting user's favorite character")
		w.Respond(components.NewResp().Content("An error occurred setting this character").Ephemeral())
		return
	}

	w.Respond(components.NewResp().Contentf("Favorite character set as char id %d", optID).Ephemeral())
}

func (b *Bot) profileEditFavoriteComplete(w corde.ResponseWriter, i *corde.Request[components.AutocompleteInteractionData]) {
	id, err := i.Data.Options.String("id")
	if err != nil {
		i, _ := i.Data.Options.Int("id")
		id = strconv.Itoa(i)
	}

	chars, err := b.Store.CharsStartingWith(i.Context, i.Member.User.ID, id)
	if err != nil {
		log.Err(err).Stringer("user", i.Member.User.ID).Msg("Error getting user's characters")
		return
	}
	if len(chars) > 25 {
		chars = chars[25:]
	}

	resp := components.NewResp()
	for _, c := range chars {
		resp.Choice(c.Name, c.ID)
	}

	w.Autocomplete(resp)
}

func (b *Bot) profileEditQuote(w corde.ResponseWriter, i *corde.Request[components.SlashCommandInteractionData]) {
	quote, _ := i.Data.Options.String("value")
	if len(quote) > 1024 {
		w.Respond(components.NewResp().Content("Quote is too long").Ephemeral())
		return
	}

	err := b.Store.SetUserQuote(i.Context, i.Member.User.ID, quote)
	if err != nil {
		log.Ctx(i.Context).Err(err).Stringer("user", i.Member.User.ID).Str("quote", quote).Msg("Error setting user's favorite character")
		w.Respond(components.NewResp().Content("An error occurred setting this character").Ephemeral())
		return
	}

	w.Respond(components.NewResp().Content("Quote set").Ephemeral())
}
