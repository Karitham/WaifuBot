# Waifu Bot

<p align="center">
  <a target="_blank" href="https://discord.com/oauth2/authorize?client_id=712332547694264341&permissions=1074097217&scope=bot" ><img alt="Add it to your discord" src="https://img.shields.io/badge/WaifuBot-ADD%20IT-brightgreen?style=for-the-badge"></a>
  <a target="_blank" href="https://www.codefactor.io/repository/github/karitham/waifubot"><img alt="code factor grade" src="https://img.shields.io/codefactor/grade/github/karitham/waifubot?color=brightgreen&style=for-the-badge"></a>
  <a target="_blank" href="https://hub.docker.com/repository/docker/karithamdocker/go-waifubot"><img alt="Docker Pulls" src="https://img.shields.io/docker/pulls/karithamdocker/go-waifubot?color=brightgreen&style=for-the-badge"></a>
  <a target="_blank" href="https://hub.docker.com/repository/docker/karithamdocker/go-waifubot"><img alt="Docker Image Size (latest by date)" src="https://img.shields.io/docker/image-size/karithamdocker/go-waifubot?color=brightgreen&style=for-the-badge"></a>
</p>

## Commands

```help
Help
---
Commands
claim string...: claim a dropped character
favorite string...: set a char as favorite
give CharacterID @user: give a char to a user
help: display general help
invite: send invite link
list @user...: display user characters
profile @user...: display user profile
quote string...: set profile quote
roll: roll a random character
verify CharacterID... @user...: check if a user owns the waifu
---
Subcommands
search: Search for characters, anime, manga and users
search anime string...: search for an anime
search character string...: search for a character
search manga string...: search for a manga
search user string...: search for an anilist user
trending: View trending manga and anime
trending anime: search for an anime
trending manga: search for a manga
```

## Feature Requests

I'm open to new feature requests, please follow the template given and/or formulate a clear request

## Contribution

If you want to contribute, just try to follow a clear style just as the rest of the code

Try to space your code and follow general programming guidelines

**About comments :**

> Good code is its own best documentation. As you’re about to add a comment, ask yourself, ‘How can I improve the code so that this comment isn’t needed?
– Steve McConnell

See [Project](https://github.com/Karitham/WaifuBot/projects/1) to view the current progress of the bot

# Deploying Yourself

## Requirements

- [Docker](https://docs.docker.com/get-docker/)
- [docker-compose](https://docs.docker.com/compose/install/)
- [A discord bot token](https://discord.com/developers)

## Setup

Rename `configExample.toml` to `config.toml` and change the values according to your needs

## Run

`docker-compose up -d`

# Dependencies

- [Arikawa](https://github.com/diamondburned/arikawa)
- [DGWidgets](https://github.com/diamondburned/dgwidgets)
- [Machinebox/graphql](https://github.com/machinebox/graphql)
- [BurntSushi/toml](github.com/BurntSushi/toml)
- [Anilist API](https://github.com/AniList/ApiV2-GraphQL-Docs)
- [MongoDB](https://mongodb.com)
- and their following deps

# LICENSE

```license
Copyright 2020 PERY "Karitham" Pierre-Louis

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
```

# Thanks

Much thanks to people from the gopher discord for all the help with go
