# Waifu Bot

This is a waifu/husbando bot in developpement. It is unstable and will stay unstable until 1.0

It uses [Disgord go lib](github.com/andersfylling/disgord), [Anilist's GraphQL API](https://github.com/AniList/ApiV2-GraphQL-Docs) and [mongoDB](https://mongodb.com)

## Requirements

* Golang (latest version recommended)
* MongoDB
* A discord bot token

## setup

Rename `configExemple.json` to `config.json` and change the values according to your needs

Run `go mod init bot` to initialise a bot.

To run the bot, do `go run .`

## Building

To build for the platform you are on, just do

`go build`

To build for raspberry Pi, do

`env GOOS=linux GOARCH=arm GOARM=7 go build`

If it doesn't work, check ARM version, this exemple works on raspberry Pi 4, I'm not sure about other models
