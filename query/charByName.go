package query

import (
	"context"

	"github.com/machinebox/graphql"
)

// CharByNameSruct handles data from CharByName queries
type CharByNameSruct struct {
	Character struct {
		ID      int    `json:"id"`
		SiteURL string `json:"siteUrl"`
		Name    struct {
			Full string `json:"full"`
		}
		Image struct {
			Large string `json:"large"`
		}
		Description string `json:"description"`
	}
}

// CharByName makes a query to the anilist API based on the name you input
func CharByName(name string) (CharByNameSruct, error) {
	var res CharByNameSruct
	graphURL := "https://graphql.anilist.co"
	client := graphql.NewClient(graphURL)
	req := graphql.NewRequest(`
	query ($search: String) {
		Character(search: $search) {
			id
			siteUrl
			name {
				full
			}
			image {
				large
			}
			description
		}
	}
	`)
	req.Var("search", name)
	ctx := context.Background()
	err := client.Run(ctx, req, &res)

	return res, err
}
