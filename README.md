# Waifu Bot

This is a waifu/husbando bot in developpement. It is unstable and will stay unstable until 1.0

It uses [Disgord go lib](github.com/andersfylling/disgord), [Anilist's GraphQL API](https://github.com/AniList/ApiV2-GraphQL-Docs) and [mongoDB](https://mongodb.com)

## Requirements

* [Golang](https://golang.org/) (latest version recommended)
* [MongoDB](https://mongodb.com) (A working url to your cluster / database)
* A discord bot token [(go there to get one](discordapp.com/developers)

## Setup

Rename `configExample.json` to `config.json` and change the values according to your needs

Run `go mod init bot` to initialise a bot.

To run the bot, do `go run .`

## Building

To build for the platform you are on, just do

`go build`

To build for other models, specify the OS & architecture, if nothing is specified it defaults to the current OS / architecture

### Exemple for Raspberry Pi 4

`env GOOS=linux GOARCH=arm GOARM=7 go build`

If it doesn't work, check ARM version, this exemple works on raspberry Pi 4, I'm not sure about other models

## Contribution

If you want to contribute, just try to follow a clear style just as the rest of the code

Check the Project section to see advance in the project v1

## Feature Requests

I'm open to new feature requests, please follow the template given and/or formulate a clear request
