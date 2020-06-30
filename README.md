# Waifu Bot

This is a waifu / husbando bot in developpement.

I used the [Disgord go lib](github.com/andersfylling/disgord) & [Anilist's GraphQL API](https://github.com/AniList/ApiV2-GraphQL-Docs)

## Requirements

* Golang (latest version recommended)
* A discord bot token

## SETUP

Create a `config.json` file and put your token in the form of a string in it

### Exemple

```json
{
    "Prefix": "w.",
    "Bot_Token": "thIsIsaDiscorDToken.$dzahuidladsdazadgegdj",
    "Max_Character_Roll": 7000
}
```

Run `go mod download` to download the depedencies needed.

Do not touch them in case you do not know what you are doing.

To run the bot, do `go run .`

## Building

To build for the platform you are on, just do

`go build`

To build for raspberry Pi, do, please check ARM version, this exemple works on raspberry Pi 4, but I don't know about the other models

`env GOOS=linux GOARCH=arm GOARM=7 go build`