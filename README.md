# Waifu Bot

<p align="center">
  <img src="https://github.com/Karitham/WaifuBot/workflows/golangci-lint/badge.svg">
</p>

## About

This is a waifu/husbando bot in developpement. It is unstable and will stay unstable until 1.0

It uses [Disgord go lib](github.com/andersfylling/disgord), [Anilist's GraphQL API](https://github.com/AniList/ApiV2-GraphQL-Docs) and [mongoDB](https://mongodb.com)

You can add the bot to your server [using this link](https://discord.com/oauth2/authorize?scope=bot&client_id=712332547694264341&permissions=0)

## Commands

| command         | alias | description                                       |
| --------------- | ----- | ------------------------------------------------- |
| `roll`          | `r`   | Rolls a random waifu                              |
| `list`          | `l`   | Lists all waifu claimed / rolled                  |
| `claim`         | `c`   | Claims a randomly dropped waifu                   |
| `profile`       | `p`   | View your profile                                 |
| `give`          | `g`   | Give a character to someone                       |
| `search`        | `s`   | Searches for an anime character                   |
| `favourite`     | `f`   | Adds a character as your favourite                |
| `quote`         | `q`   | Adds a custom quote on your profile               |
| `trendingAnime` | `ta`  | View currently trending anime                     |
| `searchAnime`   | `sa`  | Search for an anime                               |
| `invite`        |       | Send invite link to invite the bot to your server |

## Feature Requests

I'm open to new feature requests, please follow the template given and/or formulate a clear request

## Contribution

If you want to contribute, just try to follow a clear style just as the rest of the code

Try to space your code and follow general programming guidelines

About comments : The code should speak for itself, only comment big processes / dense places.

See [Project](https://github.com/Karitham/WaifuBot/projects/1) to view the current progress of the bot

# Using it for yourself

## Requirements

- A recent [Golang](https://golang.org/) version
- A working [MongoDB](https://mongodb.com) database / URL to a cluster
- A [discord bot token](discordapp.com/developers)

## Setup

Rename `configExample.json` to `config.json` and change the values according to your needs

## Run

You can run the bot by either doing `go run .` or building ( `go build` ) it and then running the `bot` executable binary.

# Docker

## Requirements

- Docker
- A working [MongoDB](https://mongodb.com) database / URL to a cluster
- A [discord bot token](discordapp.com/developers)

## Setup

Rename `configExample.json` to `config.json` and change the values according to your needs

Create a `config.json` file in *~/WaifuBot/Config.json* looking like `configExample.json` but replacing with your own values

## Run

`docker run --network host -d -v ~/WaifuBot/config.json:/home/waifubot/config.json --name waifubot --restart always karithamdocker/go-waifubot`

# Thanks

Much thanks to people from the gopher discord for all the help with go and Anders for its library.
