# Waifu Bot

**[Add it to your server](https://discord.com/oauth2/authorize?scope=bot&client_id=712332547694264341&permissions=92224)**

Waifu bot is discord waifu/husbando gacha bot.

It also has a read-only API to get data from. See [go-waifubot/API](https://github.com/go-waifubot/API).

There is also a website linked to the API & the bot at [go-waifubot/WaifuGUI](https://github.com/go-waifubot/WaifuGUI).

View your list at [waifugui.kar.moe](https://waifugui.kar.moe) and retrieve any data you need at [waifuapi.kar.moe](https://waifuapi.kar.moe) (see repo for more information).

## Commands

```help
Help
---
Commands
      claim|c|C string...: claim a dropped character
      favorite|f|F string...: set a char as favorite
      give|g|G int64 @user: give a char to a user
      help|h|H: display general help
      list|l|L @user...: display user characters
      profile|p|P @user...: display user profile
      quote|q|Q string...: set profile quote
      roll|r|R: roll a random character
      verify|v|V int64 @user...: check if a user owns the waifu
---
Subcommands
      search|s: Search for characters, anime, manga and users
            search anime|a|A string...: search for an anime
            search character|c|C string...: search for a character
            search manga|m|M string...: search for a manga
            search user|u|U string...: search for an anilist user
      trending|t: View trending manga and anime
            trending anime: search for an anime
            trending manga: search for a manga
```

## Feature Requests

I'm open to new feature requests, please follow the template given and/or formulate a clear request

## Deploying Yourself

### Requirements

[Docker](https://docs.docker.com/get-docker/)

[docker-compose](https://docs.docker.com/compose/install/)

[A discord bot token](https://discord.com/developers)

### Setup

Rename `.env.example` to `.env` and change the values according to your needs

### Run

`docker-compose up -d`

## Development

The database part is generated using `sqlc`. See the documentation on [the repo](https://github.com/kjconroy/sqlc)

If you make any modifications to the files in `/sql` you can re-generate the code using either the docker image like so,

```sh
docker run --rm -v $(pwd):/src -w /src kjconroy/sqlc generate
```

or just the command line tool

```sh
sqlc generate
```
