package query

import (
	"context"
	"fmt"

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
func CharSearchByPopularity(id int) CharStruct {
	var res CharStruct

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

	// Add varriables
	req.Var("pageNumber", id)

	// Make request
	err := graphql.NewClient(graphURL).Run(context.Background(), req, &res)
	if err != nil {
		fmt.Println("Error getting random character", err)
	}

	return res
}
