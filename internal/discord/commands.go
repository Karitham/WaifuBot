package discord

import "github.com/Karitham/corde"

var commands = []corde.Command{
	{
		Name:        "roll",
		Description: "roll a random character",
		Type:        corde.COMMAND_CHAT_INPUT,
	},
	{
		Name:        "list",
		Description: "list owned characters",
		Type:        corde.COMMAND_CHAT_INPUT,
		Options: []corde.Option{
			{
				Name:        "user",
				Type:        corde.OPTION_USER,
				Description: "optional user to list characters for",
			},
		},
	},
	{
		Name:        "search",
		Description: "search for anything on anilist",
		Type:        corde.COMMAND_CHAT_INPUT,
		Options: []corde.Option{
			{
				Name:        "anime",
				Description: "search for an anime",
				Type:        corde.OPTION_SUB_COMMAND,
				Options: []corde.Option{
					{
						Name:        "search",
						Description: "title of the anime",
						Required:    true,
						Type:        corde.OPTION_STRING,
					},
				},
			},
			{
				Name:        "manga",
				Description: "search for a manga",
				Type:        corde.OPTION_SUB_COMMAND,
				Options: []corde.Option{
					{
						Name:        "search",
						Description: "title of the manga",
						Required:    true,
						Type:        corde.OPTION_STRING,
					},
				},
			},
			{
				Name:        "char",
				Description: "search for a character",
				Type:        corde.OPTION_SUB_COMMAND,
				Options: []corde.Option{
					{
						Name:        "search",
						Description: "name of the char",
						Required:    true,
						Type:        corde.OPTION_STRING,
					},
				},
			},
			{
				Name:        "user",
				Description: "search for a user",
				Type:        corde.OPTION_SUB_COMMAND,
				Options: []corde.Option{
					{
						Name:        "search",
						Description: "name of the user",
						Required:    true,
						Type:        corde.OPTION_STRING,
					},
				},
			},
		},
	},
}
