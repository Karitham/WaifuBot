package disc

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/Karitham/WaifuBot/database"
	"github.com/diamondburned/arikawa/bot"
	"github.com/diamondburned/arikawa/discord"
	"github.com/diamondburned/arikawa/gateway"
)

type messageEvent struct {
	ChannelID discord.ChannelID
	GuildID   discord.GuildID
	MessageID discord.MessageID
	UserID    discord.UserID
}

// ListMap is used as a map for storing mssage events
var ListMap = make(map[messageEvent]int)

// List shows the user's list
func (b *Bot) List(m *gateway.MessageCreateEvent, page bot.RawArguments) (*discord.Embed, error) {
	var p int
	var err error

	if page != "" {
		p, err = strconv.Atoi(string(page))
		if err != nil {
			return nil, err
		}
	}

	uData, err := database.ViewUserData(m.Author.ID)
	if err != nil {
		return nil, err
	}

	embed, err := createListEmbed(m.Author, p, uData.Waifus)
	if err != nil {
		return nil, err
	}

	// Init message in map
	ListMap[messageEvent{
		ChannelID: m.ChannelID,
		GuildID:   m.GuildID,
		MessageID: m.ID,
		UserID:    m.Author.ID,
	}] = p

	return embed, nil
}

// UpdateList is used to update the list
func (b *Bot) UpdateList(m *gateway.MessageCreateEvent) {
	msgCh := make(chan *gateway.MessageCreateEvent)
	cancel := b.Ctx.AddHandler(msgCh)
	for {
		select {
		case <-time.After(1 * time.Second):
			cancel()
			return
		case msg := <-msgCh:
			if msg.Author.ID != b.Ctx.Ready.User.ID {
				continue
			}

			fmt.Println(msg)
		}
	}

}

func (b *Bot) handleReactions(m gateway.MessageReactionAddEvent) {
	if m.Emoji.Name != "➡️" || m.Emoji.Name != "⬅️" {
		return
	}
	if p, ok := ListMap[messageEvent{
		ChannelID: m.ChannelID,
		GuildID:   m.GuildID,
		MessageID: m.MessageID,
		UserID:    m.UserID,
	}]; ok {
		if m.Emoji.Name == "➡️" {
			p++
		} else if m.Emoji.Name == "⬅️" && p > 0 {
			p--
		}
		uData, err := database.ViewUserData(m.Member.User.ID)
		if err != nil {
			log.Println(err)
		}

		embed, err := createListEmbed(m.Member.User, p, uData.Waifus)
		if err != nil {
			log.Println(err)
		}
		b.Ctx.Client.DeleteAllReactions(m.ChannelID, m.MessageID)
		b.Ctx.React(m.ChannelID, m.MessageID, "⬅️")
		b.Ctx.React(m.ChannelID, m.MessageID, "➡️")
		b.Ctx.Client.EditEmbed(m.ChannelID, m.MessageID, *embed)
	}
}

func createListEmbed(user discord.User, page int, list []database.CharLayout) (embed *discord.Embed, err error) {
	return &discord.Embed{
		Title: fmt.Sprintf("%s's list page %d", user.Username, page),
		Description: func(l []database.CharLayout) (d string) {
			if len(l) > page*c.ListLen+c.ListLen {
				l = l[page*c.ListLen : page*c.ListLen+c.ListLen]
			} else if len(l) > page*c.ListLen {
				l = l[page*c.ListLen:]
			}
			for _, waifu := range l {
				d += fmt.Sprintf("%d - %s\n", waifu.ID, waifu.Name)
			}
			return
		}(list),
	}, nil
}
