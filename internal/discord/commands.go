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
