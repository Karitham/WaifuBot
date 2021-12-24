package discord

import (
	"fmt"
	"strings"
	"time"

	"github.com/Karitham/corde"
)

func (b *Bot) List(w corde.ResponseWriter, i *corde.Interaction) {
	u := i.Member.User.ID
	user, ok := i.Data.Options["user"].(string)
	if ok {
		u = corde.SnowflakeFromString(user)
	}

	chars, err := b.Store.Characters(u)
	if err != nil {
		ephemeral(w, "An error occurred dialing the database, please try again later")
		return
	}

	if len(chars) == 0 {
		ephemeral(w, "No characters found")
		return
	}

	w.WithSource(&corde.InteractionRespData{
		Embeds: []corde.Embed{
			{
				Title: "Characters",
				Description: func() string {
					s := strings.Builder{}
					s.WriteString("```txt\n")
					for _, c := range chars {
						s.WriteString(fmt.Sprintf("%d - %s - %s\n", c.ID, c.Name, c.Date.Format(time.Stamp)))
					}
					s.WriteString("```")
					return s.String()
				}(),
				Color: ColorToInt("#b00b69"),
			},
		},
	})
}
