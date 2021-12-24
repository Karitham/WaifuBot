package anilist

import (
	"context"
	"errors"

	"github.com/machinebox/graphql"
)

func (a *Anilist) Random(notIn []int) (Character, error) {
	var q struct {
		Page struct {
			Characters []Character `json:"characters"`
		} `json:"Page"`
	}

	req := graphql.NewRequest(`
	query ($pageNumber: Int, $not_in: [Int]) {
		Page(perPage: 1, page: $pageNumber) {
		  characters(sort: FAVOURITES_DESC, id_not_in: $not_in) {
			id
			siteUrl
			image {
			  large
			}
			name {
			  full
			}
			media(perPage: 1, sort: POPULARITY_DESC) {
			  nodes {
				title {
				  romaji
				}
			  }
			}
		  }
		}
	  }	
	`)

	req.Var("pageNumber", a.seed.Int63()%a.MaxChars)
	req.Var("not_in", notIn)

	err := a.c.Run(context.Background(), req, &q)
	if err != nil {
		return Character{}, err
	}
	if len(q.Page.Characters) < 1 {
		return Character{}, errors.New("no characters found")
	}

	return q.Page.Characters[0], nil
}
