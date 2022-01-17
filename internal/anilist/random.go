package anilist

import (
	"context"
	"errors"

	"github.com/machinebox/graphql"
)

// CharAndMedia represent character object
type CharAndMedia struct {
	Character
	MediaTitle string
}

func (a *Anilist) RandomChar(notIn ...int64) (CharAndMedia, error) {
	type CharNMediaTmp struct {
		Character
		Media struct {
			Nodes []struct {
				Title struct {
					Romaji string `json:"romaji"`
				}
			}
		}
	}

	var q struct {
		Page struct {
			Characters []CharNMediaTmp `json:"characters"`
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
		return CharAndMedia{}, err
	}
	if len(q.Page.Characters) < 1 {
		return CharAndMedia{}, errors.New("no characters found")
	}

	media := ""
	if len(q.Page.Characters[0].Media.Nodes) > 0 {
		media = q.Page.Characters[0].Media.Nodes[0].Title.Romaji
	}

	return CharAndMedia{
		Character:  q.Page.Characters[0].Character,
		MediaTitle: media,
	}, nil
}
