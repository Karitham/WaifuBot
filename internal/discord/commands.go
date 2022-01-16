package discord

import "github.com/Karitham/corde"

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
