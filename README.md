# Waifu Bot

<p align="center">
  <img alt="Travis (.com)" src="https://img.shields.io/travis/com/karitham/waifubot?style=for-the-badge">
  
  <img alt="Docker Pulls" src="https://img.shields.io/docker/pulls/karithamdocker/go-waifubot?style=for-the-badge">

  <img alt="Docker Image Size (latest by date)" src="https://img.shields.io/docker/image-size/karithamdocker/go-waifubot?style=for-the-badge">
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

# Deploying Yourself

## Requirements

- [Docker](https://docs.docker.com/get-docker/)
- [docker-compose](https://docs.docker.com/compose/install/)
- [A discord bot token](discordapp.com/developers)

## Setup

Rename `configExample.json` to `config.json` and change the values according to your needs

## Run

`docker-compose up -d`

# Developpement

## Requirements

- A recent [Golang](https://golang.org/) version
- A working [MongoDB](https://mongodb.com) database / URL to a cluster
- A [discord bot token](discordapp.com/developers)

## Setup

Rename `configExample.json` to `config.json` and change the values according to your needs

## Run

You can run the bot by either doing `go run .` or building ( `go build` ) it and then running the `bot` executable binary.

# Thanks

Much thanks to people from the gopher discord for all the help with go and Anders for its library.
