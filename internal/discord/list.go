package discord

import (
	"fmt"

	"github.com/Karitham/corde"
)

func (b *Bot) list(w corde.ResponseWriter, i *corde.Request[corde.SlashCommandInteractionData]) {
	user := i.Member.User
	if len(i.Data.Resolved.Members) > 0 {
		user = i.Data.Resolved.Users.First()
	}

	chars, err := b.Store.Chars(i.Context, user.ID)
	if err != nil {
		w.Respond(rspErr("An error occurred dialing the database, please try again later"))
		return
	}

	if len(chars) == 0 {
		w.Respond(rspErr("This user doesn't appear to have any characters"))
		return
	}

	embed := corde.NewEmbed().
		Titlef("%s's List", user.Username).
		Thumbnail(corde.Image{URL: user.AvatarPNG()}).
		URL(fmt.Sprintf("https://waifugui.karitham.dev/#/list/%s", user.ID.String()))

	if len(chars) > 18 {
		chars = chars[:18]
	}

	for _, c := range chars {
		embed.FieldInline(c.Name, fmt.Sprintf("%d — %s", c.ID, c.Date.Format("02/01")))
	}

	w.Respond(corde.NewResp().Embeds(embed).Ephemeral())
}

func list(chars []Character) []corde.Field {
	f := make([]corde.Field, 0, len(chars))

	return f
}
