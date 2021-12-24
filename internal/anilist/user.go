package anilist

import (
	"context"
	"errors"

	"github.com/machinebox/graphql"
)

type Avatar struct {
	Large string `json:"large"`
}

type User struct {
	Name        string `json:"name"`
	About       string `json:"about"`
	Siteurl     string `json:"siteUrl"`
	Bannerimage string `json:"bannerImage"`
	Avatar      Avatar `json:"avatar"`
	ID          int    `json:"id"`
}

// User queries the anilist user.
func (a *Anilist) User(name string) ([]User, error) {
	var q struct {
		Page struct {
			Users []User `json:"users"`
		} `json:"page"`
	}
	req := graphql.NewRequest(`
query ($name: String) {
  Page {
    users(search: $name) {
      id
      name
      about
      siteUrl
      avatar {
        large
      }
      bannerImage
    }
  }
}
	`)
	req.Var("name", name)

	err := a.c.Run(context.Background(), req, &q)
	if err != nil {
		return nil, err
	}
	if len(q.Page.Users) < 1 {
		return nil, errors.New("no users found")
	}

	return q.Page.Users, nil
}
