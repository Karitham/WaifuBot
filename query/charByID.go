package query

import (
	"context"

	"github.com/machinebox/graphql"
)

// CharByIDStruct handles data from CharByID queries
type CharByIDStruct struct {
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

// CharByID makes a query to the anilist API based on the name you input
func CharByID(id int) (CharByIDStruct, error) {
	var res CharByIDStruct
	graphURL := "https://graphql.anilist.co"
	client := graphql.NewClient(graphURL)
	req := graphql.NewRequest(`
	query ($id: Int) {
		Character(id: $id) {
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
	req.Var("id", id)
	ctx := context.Background()
	err := client.Run(ctx, req, &res)

	return res, err
}
