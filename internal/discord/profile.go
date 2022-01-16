package discord

import (
	"fmt"
	"time"

	"github.com/Karitham/corde"
	"github.com/rs/zerolog/log"
)

type Profile struct {
	User
	CharacterCount int
	Favorite       Character
}

func (b *Bot) profile(w corde.ResponseWriter, i *corde.InteractionRequest) {
	user := i.Member.User
	if len(i.Data.Resolved.Users) > 0 {
		user = i.Data.Resolved.Users.First()
	}

	data, err := b.Store.Profile(i.Context, user.ID)
	if err != nil {
		log.Err(err).Msg("Error getting user's profile")
		w.Respond(corde.NewResp().Content("An error occurred dialing the database, please try again later").Ephemeral())
		return
	}

	resp := corde.NewEmbed().
		Titlef("%s' Profile", user.Username).
		URL(fmt.Sprintf("https://waifugui.kar.moe/#/list/%s", user.ID.String())).
		Descriptionf(
			"%s\n%s last rolled %s ago and have %d tokens.\nThey have %d characters.\nTheir favorite character is %s",
			data.Quote,
			user.Username,
			time.Since(data.Date.UTC()).Truncate(time.Second),
			data.Tokens,
			data.CharacterCount,
			data.Favorite.Name,
		)
	if data.Favorite.Image != "" {
		resp.Thumbnail(corde.Image{URL: data.Favorite.Image})
	}

	w.Respond(resp)
}
