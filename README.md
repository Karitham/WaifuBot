# Waifu Bot

<p align="center">
  <a href="https://discord.com/oauth2/authorize?client_id=712332547694264341&permissions=1074097217&scope=bot"><img alt="Add it to your discord" src="https://img.shields.io/badge/WaifuBot-ADD%20IT-green?style=for-the-badge"></a>
  
  <img alt="Docker Pulls" src="https://img.shields.io/docker/pulls/karithamdocker/go-waifubot?color=green&style=for-the-badge">

  <img alt="Docker Image Size (latest by date)" src="https://img.shields.io/docker/image-size/karithamdocker/go-waifubot?color=green&style=for-the-badge">
</p>

This is a waifu/husbando discord bot

It uses [Arikawa](https://github.com/diamondburned/arikawa), [Anilist's GraphQL API](https://github.com/AniList/ApiV2-GraphQL-Docs) and [mongoDB](https://mongodb.com)

You can add the bot to your server [using this link](https://discord.com/api/oauth2/authorize?client_id=733399598126596128&permissions=1074097217&scope=bot)

## Commands

| command         | alias | description                                           |
| --------------- | ----- | ----------------------------------------------------- |
| `roll`          | `r`   | Rolls a random waifu                                  |
| `list`          | `l`   | Lists all waifus you claimed / rolled                 |
| `claim`         | `c`   | Claims a randomly dropped waifu                       |
| `profile`       | `p`   | View your profile                                     |
| `give`          | `g`   | Give a character to someone                           |
| `search`        | `s`   | Searches for an anime character                       |
| `favourite`     | `f`   | Adds a character as your favourite                    |
| `quote`         | `q`   | Adds a custom quote on your profile                   |
| `trendingAnime` | `ta`  | View currently trending anime (on AniList)            |
| `trendingManga` | `tm`  | View currently trending manga (on AniList)            |
| `searchAnime`   | `sa`  | Search for an anime                                   |
| `searchManga`   | `sm`  | Search for an manga                                   |
| `invite`        | `i`   | Send an invite link to invite the bot to your server  |
| `verify`        | `v`   | Verifies if mentioned user has got the waifu you want |

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
- [docker-compose](https://docs.docker.com/compose/install/) (latest version recommended)
- [A discord bot token](https://discord.com/developers)

## Setup

Rename `configExample.toml` to `config.toml` and change the values according to your needs

## Run

`docker-compose up -d`

# Thanks

Much thanks to people from the gopher discord for all the help with go
