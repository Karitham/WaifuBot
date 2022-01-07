package discord

import (
	"fmt"
	"math"
	"strings"
	"sync"
	"time"

	"github.com/Karitham/WaifuBot/internal/anilist"
	"github.com/Karitham/corde"
)

type listState struct {
	mu sync.Mutex
	// intial interaction ID : pageState
	state map[corde.Snowflake]pageState
}

type pageState struct {
	page            uint
	userID          corde.Snowflake
	user            string
	lastInteraction time.Time
}

var btnBack = corde.Component{
	Type:     corde.COMPONENT_BUTTON,
	Style:    corde.BUTTON_SECONDARY,
	Label:    "Prev",
	Emoji:    &corde.Emoji{Name: "⬅️"},
	CustomID: "list/back",
}

var btnNext = corde.Component{
	Type:     corde.COMPONENT_BUTTON,
	Style:    corde.BUTTON_SECONDARY,
	Label:    "Next",
	Emoji:    &corde.Emoji{Name: "➡️"},
	CustomID: "list/next",
}

func (b *Bot) list(m *corde.Mux) {
	m.Command("", b.listCommand())
	m.Button("back", b.listBack)
	m.Button("next", b.listNext)
}

func (b *Bot) listCommand() func(w corde.ResponseWriter, i *corde.InteractionRequest) {
	b.listState.state = make(map[corde.Snowflake]pageState)
	// Garbage collector
	go func() {
		for {
			time.Sleep(time.Minute)
			b.listState.mu.Lock()

			for k, v := range b.listState.state {
				if time.Since(v.lastInteraction) > time.Minute*5 {
					delete(b.listState.state, k)
				}
			}
			b.listState.mu.Unlock()
		}
	}()

	return func(w corde.ResponseWriter, i *corde.InteractionRequest) {
		u := i.Member.User.ID
		name := i.Member.User.Username

		if user := i.Data.Options.String("user"); user != "" {
			name = user
			u = corde.SnowflakeFromString(user)
		}

		chars, err := b.Store.Characters(u)
		if err != nil {
			w.Respond(corde.NewResp().Content("An error occurred dialing the database, please try again later").Ephemeral())
			return
		}

		if len(chars) == 0 {
			w.Respond(corde.NewResp().Content("This user doesn't appear to have any characters").Ephemeral())
			return
		}

		b.listState.mu.Lock()
		defer b.listState.mu.Unlock()
		b.listState.state[i.ID] = pageState{
			userID:          u,
			lastInteraction: time.Now(),
			user:            name,
		}

		c := chars
		if len(c) > 18 {
			c = c[:18]
		}

		w.Respond(corde.NewResp().
			Embeds(listEmbed(name, c)).
			ActionRow(setCustomID(i.ID.String(), btnBack, btnNext)...),
		)
	}
}

func (b *Bot) listNext(w corde.ResponseWriter, i *corde.InteractionRequest) {
	b.listState.mu.Lock()
	defer b.listState.mu.Unlock()

	id := getMessageID(i.Data.CustomID)
	s := b.listState.state[id]

	chars, err := b.Store.Characters(s.userID)
	if err != nil {
		w.Update(corde.NewResp().Content("An error occurred dialing the database, please try again later").Ephemeral())
		return
	}

	if len(chars) == 0 {
		w.Update(corde.NewResp().Content("No characters found").Ephemeral())
		return
	}

	s.lastInteraction = time.Now()
	s.page = (s.page - 1) % uint(math.Ceil(float64(len(chars))/18))
	b.listState.state[id] = s

	c := chars[s.page*18:]
	if len(c) > 18 {
		c = c[:18]
	}

	w.Update(corde.NewResp().
		Embeds(listEmbed(s.user, c)).
		ActionRow(setCustomID(id.String(), btnBack, btnNext)...),
	)
}

func (b *Bot) listBack(w corde.ResponseWriter, i *corde.InteractionRequest) {
	b.listState.mu.Lock()
	defer b.listState.mu.Unlock()

	id := getMessageID(i.Data.CustomID)
	s := b.listState.state[id]

	chars, err := b.Store.Characters(s.userID)
	if err != nil {
		w.Update(corde.NewResp().Content("An error occurred dialing the database, please try again later").Ephemeral())
		return
	}

	if len(chars) == 0 {
		w.Update(corde.NewResp().Content("No characters found").Ephemeral())
		return
	}

	s.lastInteraction = time.Now()
	s.page = (s.page - 1) % uint(math.Ceil(float64(len(chars))/18))
	b.listState.state[id] = s

	c := chars[s.page*18:]
	if len(c) > 18 {
		c = c[:18]
	}

	w.Update(corde.NewResp().
		Embeds(listEmbed(s.user, c)).
		ActionRow(setCustomID(id.String(), btnBack, btnNext)...),
	)
}

func setCustomID(id string, cs ...corde.Component) []corde.Component {
	var bs []corde.Component
	for _, c := range cs {
		if !strings.HasSuffix(c.CustomID, "/") {
			c.CustomID += "/"
		}
		c.CustomID += id
		bs = append(bs, c)
	}
	return bs
}

func getMessageID(customID string) corde.Snowflake {
	s := strings.Split(customID, "/")
	return corde.SnowflakeFromString(s[len(s)-1])
}

func listEmbed(name string, chars []Character) *corde.EmbedB {
	return corde.NewEmbed().
		Titlef("%s's List", name).
		Color(anilist.Color).
		Fields(list(chars)...)
}

func list(chars []Character) []corde.Field {
	f := make([]corde.Field, 0, len(chars))
	for _, c := range chars {
		f = append(f, corde.Field{
			Name:   c.Name,
			Value:  fmt.Sprintf("%d — **%s**", c.ID, c.Date.Format("01/02/06")),
			Inline: true,
		})
	}
	return f
}
