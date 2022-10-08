package discord

import (
	"context"
	"strings"
	"time"

	"github.com/Karitham/corde"
	"github.com/rs/zerolog/log"
)

func (b *Bot) drop(ctx context.Context, guildID corde.Snowflake, channelID corde.Snowflake) {
	log.Debug().Msgf("drop character to %d in %d", guildID, channelID)

	char, err := b.AnimeService.RandomChar(ctx)
	if err != nil {
		log.Error().Err(err).Msg("failed to get random character")
		return
	}

	err = b.Inter.SetChannelChar(ctx, channelID, char)
	if err != nil {
		log.Error().Err(err).Msg("failed to set channel character")
		return
	}

	log.Trace().Msgf("dropped %s", char.Name)

	// FIXME: send message to channel.
	// _, err = b.mux.SendEmbed(ctx, channelID, charEmbed(char))
}

func (b *Bot) claim(w corde.ResponseWriter, i *corde.Request[corde.SlashCommandInteractionData]) {
	name, err := i.Data.Options.String("name")
	if err != nil || name == "" {
		w.Respond(rspErr("enter a name to claim the character"))
		return
	}

	// impl claim.
	char, err := b.Inter.GetChannelChar(i.Context, i.ChannelID)
	if err != nil {
		log.Trace().Err(err).Msg("failed to get channel character")
		w.Respond(rspErr("No character to claim"))
		return
	}

	if !equalStrings(char.Name, name) {
		w.Respond(rspErr("Wrong!"))
		return
	}

	err = b.Store.PutChar(i.Context, i.Member.User.ID, Character{
		Date:   time.Now(),
		Image:  char.ImageURL,
		Name:   sanitizeName(char.Name),
		Type:   "CLAIM",
		UserID: i.Member.User.ID,
		ID:     char.ID,
	})
	if err != nil {
		w.Respond(newErrf("you already have this character!"))
		return
	}

	w.Respond(corde.NewResp().Contentf("Congrats %s! You claimed %s.", i.Member.Nick, char.Name))
}

func equalStrings(this, that string) bool {
	return strings.EqualFold(sanitizeName(this), sanitizeName(that))
}

func sanitizeName(name string) string {
	return strings.Join(strings.Fields(name), " ")
}
