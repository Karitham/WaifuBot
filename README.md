# Waifu Bot

<p align="center">
  <img alt="Travis (.com)" src="https://img.shields.io/travis/com/karitham/waifubot?style=for-the-badge">
  
  <img alt="Docker Pulls" src="https://img.shields.io/docker/pulls/karithamdocker/go-waifubot?style=for-the-badge">

  <img alt="Docker Image Size (latest by date)" src="https://img.shields.io/docker/image-size/karithamdocker/go-waifubot?style=for-the-badge">
</p>

## About

This is a waifu/husbando discord bot, it is pretty stable but some commands might change.

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
| `trendingManga` | `tm`  | View currently trending manga                     |
| `searchAnime`   | `sa`  | Search for an anime                               |
| `searchManga`   | `sm`  | Search for an manga                               |
| `invite`        | `i`   | Send invite link to invite the bot to your server |

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
- [A discord bot token](discordapp.com/developers)

## Setup

Rename `configExample.json` to `config.json` and change the values according to your needs

```js
{
  "Prefix": "w.", // This is the prefix for the bot
  "Bot_Token": "thIsIsaDiscorDToken.Xw41Og.Ix0IhJlhgEbEyK_CJwrGgucn9rk", // Place your discord bot token here
  "Mongo_URL": "mongodb://db:27017", // This is the mongoDB URI, don't change this unless you know what you are doing
  "Bot_Status": "WaifuBot | w.help", // This is the status displayed by the bot when running
  "Max_Character_Roll": 15000, // Mean the roll function can query in the 15000 most popular character. DOn't go higher than 75000
  "Max_Character_Drop": 5000, // Just like the roll function, but for dropping characters (so they are more popular)
  "Delete_Illegal_Roll_After": 3, // Delete illegal roll message to have a cleaner chat. Duration in minutes
  "Delete_Wrong_Claim_After": 1, // Delete wrong claims message. Duration in minute
  "List_Max_Update_Time": 2, // How long can the bot update the last list embed sent. Duration in minute
  "Time_Between_Rolls": 12, // Time interval between each user roll
  "Drops_On_Interact": 20 // Lowering this value increases drop rate, the opposite works too
}
```

## Run

`docker-compose up -d`

# Thanks

Much thanks to people from the gopher discord for all the help with go and Anders for its library.
