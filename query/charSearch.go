package query

import (
	"context"

	"github.com/machinebox/graphql"
)

// SearchOpts defines what search terms you can use
type SearchOpts struct {
	Name string
	ID   int
}

// CharSearchStruct handles data from CharByName queries
type CharSearchStruct struct {
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

// CharSearch makes a query to the anilist API based on the name//ID you input
func CharSearch(so SearchOpts) (CharSearchStruct, error) {
	var res CharSearchStruct

	graphURL := "https://graphql.anilist.co"
	client := graphql.NewClient(graphURL)
	req := graphql.NewRequest(`
	query ($id: Int, $name:String) {
		Character(id: $id,search: $name) {
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

	// Add variable
	if so.ID != 0 {
		req.Var("id", so.ID)
	} else {
		req.Var("name", so.Name)
	}

	ctx := context.Background()
	err := client.Run(ctx, req, &res)

	return res, err
}
