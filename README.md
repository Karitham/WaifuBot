# Waifu Go Bot

This is a waifu / husbando bot in developpement.

I used the [Disgord go lib](github.com/andersfylling/disgord) & [Anilist's GraphQL API](https://github.com/AniList/ApiV2-GraphQL-Docs)

## Requirements

* Golang (latest version recommended)
* A discord bot token

## SETUP

Create a `token.json` file and put your token in the form of a string in it

### Exemple

```json
"ThisIsAdiscordToken.ù$dzajodpzaddzadzad4898dza"
```

Run `go mod init bot` and `go mod download` to download the depedencies needed.

This will create 2 file, a `go.mod` and a `go.sum`

Do not touch them in case you do not know what you are doing.

Run the bot `go run .`

## Quick start (LINUX)

This will open vi. if you do not have vi either install it or follow the [SETUP](#setup)

Paste your token with `Ctrl+Maj+V` in the terminal. Quit & save it by typing `:wq`

```sh
git clone https://github.com/Karitham/WaifuGoBot
cd WaifuGoBot
touch token.json
vi token.json
go mod init bot
go mod download
go run .
```
