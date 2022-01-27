package anilist

import (
	"context"
	"errors"
	"fmt"

	"github.com/machinebox/graphql"
)

type Title struct {
	Romaji string `json:"romaji"`
}

type Cover struct {
	Large string `json:"large"`
	Color string `json:"color"`
}

type Media struct {
	Title       Title  `json:"title"`
	Description string `json:"description"`
	Siteurl     string `json:"siteUrl"`
	CoverImage  Cover  `json:"coverImage"`
	BannerImage string `json:"bannerImage"`
	ID          int    `json:"id"`
}

func (a *Anilist) Anime(name string) ([]Media, error) {
	return a.media(name, "ANIME")
}

func (a *Anilist) Manga(name string) ([]Media, error) {
	return a.media(name, "MANGA")
}

func (a *Anilist) media(name, typ string) ([]Media, error) {
	var q struct {
		Page struct {
			Media []Media `json:"media"`
		} `json:"page"`
	}
	req := graphql.NewRequest(fmt.Sprintf(`
query ($name: String) {
  Page {
    media(search: $name, type: %s) {
      id
      title {
        romaji
      }
      description
      siteUrl
      coverImage {
        large
        color
      }
      bannerImage
    }
  }
}
	`, typ))
	req.Var("name", name)

	err := a.c.Run(context.Background(), req, &q)
	if err != nil {
		return nil, err
	}
	if len(q.Page.Media) < 1 {
		return nil, errors.New("no media found")
	}

	return q.Page.Media, nil
}
