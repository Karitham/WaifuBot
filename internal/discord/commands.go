package discord

import (
	"github.com/Karitham/corde"
	"github.com/rs/zerolog/log"
)

var nameOpt = corde.NewStringOption("name", "name you wish to search", true)

var listCmd = corde.NewSlashCommand("list", "list owned characters",
	corde.NewUserOption("user", "user to list characters for", false),
)

var rollCmd = corde.NewSlashCommand("roll", "roll a random character")

var searchCmd = corde.NewSlashCommand("search", "search for anything on anilist",
	corde.NewSubcommand("anime", "search for an anime", nameOpt),
	corde.NewSubcommand("manga", "search for a manga", nameOpt),
	corde.NewSubcommand("char", "search for a character", nameOpt),
	corde.NewSubcommand("user", "search for a user", nameOpt),
)

var profileCmd = corde.NewSlashCommand("profile", "interact with your profile or view someone else's",
	corde.NewSubcommand("view", "view a user's profile",
		corde.NewUserOption("user", "user to view profile for", false),
	),
	corde.NewSubcommandGroup("edit", "edit your profile",
		corde.NewSubcommand("favorite", "set your favorite character",
			corde.NewIntOption("id", "id of the character", true).CanAutocomplete(),
		),
		corde.NewSubcommand("quote", "set your quote",
			corde.NewStringOption("value", "quote value to set", true),
		),
	),
)

var giveCmd = corde.NewSlashCommand("give", "give a character to someone",
	corde.NewUserOption("user", "user to give character to", true),
	corde.NewIntOption("id", "id of the character", true).CanAutocomplete(),
)

var infoCmd = corde.NewSlashCommand("info", "information about the bot")

func cmdEqual(existing corde.Command, new corde.CreateCommand) bool {
	if existing.Name != new.Name || existing.Description != new.Description || new.Type != existing.Type {
		return false
	}

	if len(existing.Options) != len(new.Options) {
		return false
	}

	return true
}

func (b *Bot) registerCommands() error {
	actual := []corde.CreateCommand{searchCmd, rollCmd, listCmd, profileCmd, giveCmd, infoCmd}

	commands, err := b.mux.GetCommands(corde.GuildOpt(b.GuildID))
	if err != nil {
		log.Err(err).Msg("Failed to get commands")
	}

	if b.ForceRegisterCMD {
		var toRegister []corde.CreateCommander
		for _, c := range actual {
			toRegister = append(toRegister, c)
		}

		log.Info().Msg("Forcing register of CMD")
		return b.mux.BulkRegisterCommand(toRegister, corde.GuildOpt(b.GuildID))
	}

	for _, c := range commands {
		for i, r := range actual {
			if cmdEqual(c, r) {
				actual = remove(actual, i)
				break
			}
		}
	}

	if len(actual) != 0 {
		var toRegister []corde.CreateCommander
		for _, c := range actual {
			toRegister = append(toRegister, c)
		}

		return b.mux.BulkRegisterCommand(toRegister, corde.GuildOpt(b.GuildID))
	}

	return nil
}

func (b *Bot) RemoveUnknownCommands(r corde.ResponseWriter, i *corde.InteractionRequest) {
	r.Respond(corde.NewResp().Content("I don't know what that means, you shouldn't be able to do that").Ephemeral())
	b.mux.DeleteCommand(i.ID, corde.GuildOpt(b.GuildID))
}

func remove[T any](s []T, i int) []T {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}
