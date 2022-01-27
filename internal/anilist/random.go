package anilist

import (
	"context"
	"errors"

	"github.com/machinebox/graphql"
	"github.com/rs/zerolog/log"
)

// CharAndMedia represent character object
type CharAndMedia struct {
	Character
	MediaTitle string
}

func (a *Anilist) RandomChar(notIn ...int64) (CharAndMedia, error) {
	c := a.internalCache["random"]
	c.Lock()
	defer c.Unlock()
	rest := []CharAndMedia{}
two:
	for id, char := range c.cache {
		for _, n := range notIn {
			if n == id {
				continue two
			}
		}

		rest = append(rest, char.(CharAndMedia))
	}

	if len(c.cache) < 20 {
		for i := 0; i < 5; i++ {
			go func() {
				ch, err := a.randomChar(notIn...)
				if err != nil {
					return
				}
				c.Lock()
				defer c.Unlock()
				c.cache[ch.ID] = ch
			}()
		}
	}

	if len(rest) > 0 {
		char := rest[a.seed.Int63()%(int64(len(rest)))]
		log.Trace().Str("char", char.Name.Full).Int("cache size", len(c.cache)).Msg("Hit cache")
		delete(c.cache, char.ID)
		return char, nil
	}

	return a.randomChar(notIn...)
}

func (a *Anilist) randomChar(notIn ...int64) (CharAndMedia, error) {
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
