package query

import (
	"context"

	"github.com/machinebox/graphql"
)

// CharSearchStruct handles data from CharByName queries
type CharSearchStruct struct {
	Character struct {
		ID      int64  `json:"id"`
		SiteURL string `json:"siteUrl"`
		Name    struct {
			Full string `json:"full"`
		}
		Image struct {
			Large string `json:"large"`
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

// CharSearchInput is used to input the arguments you want to search
type CharSearchInput struct {
	ID   int
	Name string
}

// CharSearch makes a query to the anilist API based on the name//ID you input
func (input CharSearchInput) CharSearch() (CharSearchStruct, error) {
	var res CharSearchStruct

	// build query
	client := graphql.NewClient(GraphURL)
	req := graphql.NewRequest(`
	query ($id: Int, $name: String) {
		Character(id: $id, search: $name, sort: SEARCH_MATCH) {
		  id
		  siteUrl
		  name {
			full
		  }
		  image {
			large
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
	`)
	// Add variable
	if input.ID != 0 {
		req.Var("id", input.ID)
	} else {
		req.Var("name", input.Name)
	}

	ctx := context.Background()
	err := client.Run(ctx, req, &res)

	return res, err
}
