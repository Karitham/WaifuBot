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
    "Bot_Token": "thIsIsaDiscorDToken.$dzahuidladsdazadgegdj"
}
```

Run `go mod init bot` and `go mod download` to download the depedencies needed.

This will create 2 file, a `go.mod` and a `go.sum`

Do not touch them in case you do not know what you are doing.

Run the bot `go run .`
