package query

import (
	"context"

	"github.com/machinebox/graphql"
)

// CharStruct handles data for RandomChar function
type CharStruct struct {
	Page struct {
		Characters []struct {
			ID      int64  `json:"id"`
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

// CharSearchByPopularity outputs the character you want based on their number on the page list
func CharSearchByPopularity(id int) (response CharStruct, err error) {
	// Create request
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

	// Add variable
	req.Var("pageNumber", id)

	// Make request
	err = graphql.NewClient(graphURL).Run(context.Background(), req, &response)
	return
}
