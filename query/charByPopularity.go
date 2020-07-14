package query

import (
	"context"

	"github.com/machinebox/graphql"
)

// CharStruct handles data for RandomChar function
type CharStruct struct {
	Page struct {
		Characters []struct {
			ID      int    `json:"id"`
			SiteURL string `json:"siteUrl"`
			Image   struct {
				Large string `json:"large"`
			}
			Name struct {
				Full string `json:"full"`
			}
			Media struct {
				Nodes []struct {
					Title struct {
						Romaji string `json:"romaji"`
					}
				}
			}
		}
	}
}

// RandomChar outputs the character you want based on their number on the page list
func RandomChar(id int) (CharStruct, error) {
	var res CharStruct
	graphURL := "https://graphql.anilist.co"
	client := graphql.NewClient(graphURL)
	req := graphql.NewRequest(`
	query ($pageNumber: Int) {
		Page(perPage: 1, page: $pageNumber) {
		  characters(sort: FAVOURITES_DESC) {
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

	req.Var("pageNumber", id)
	ctx := context.Background()
	err := client.Run(ctx, req, &res)

	return res, err
}
