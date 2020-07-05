# Waifu Bot

This is a waifu / husbando bot in developpement.

Tt uses [Disgord go lib](github.com/andersfylling/disgord), [Anilist's GraphQL API](https://github.com/AniList/ApiV2-GraphQL-Docs) and mongoDB

## Requirements

* Golang (latest version recommended)
* A discord bot token

## setup

Create a `config.json` file and add the needed information like below

### Exemple

```json
{
    "Prefix": "w.",
    "Bot_Token": "thIsIsaDiscorDToken.$dzahuidladsdazadgegdj",
    "Max_Character_Roll": 7000
}
```

Run `go mod init bot` to initialise a bot.

To run the bot, do `go run .`

## Building

To build for the platform you are on, just do

`go build`

To build for raspberry Pi, do

`env GOOS=linux GOARCH=arm GOARM=7 go build`

If it doesn't work, check ARM version, this exemple works on raspberry Pi 4, I'm not sure about other models
