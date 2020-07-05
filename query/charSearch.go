package query

import (
	"context"
	"strconv"
	"strings"

	"github.com/machinebox/graphql"
)

// CharSearchStruct handles data from CharByName queries
type CharSearchStruct struct {
	Character struct {
		ID      int    `json:"id"`
		SiteURL string `json:"siteUrl"`
		Name    struct {
			Full string `json:"full"`
		}
		Image struct {
			Medium string `json:"medium"`
			Large  string `json:"large"`
		}
		Description string `json:"description"`
	}
}

// CharSearch makes a query to the anilist API based on the name//ID you input
func CharSearch(args []string) (CharSearchStruct, error) {
	var res CharSearchStruct

	// build query
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
			medium
			large
		  }
		  description
		}
	  }
	`)
	// Parse the arguments to check if an ID or a Name was entered
	arg := strings.Join(args, " ")
	id, err := strconv.Atoi(arg)

	// Add variable
	if id != 0 {
		req.Var("id", id)
	} else {
		req.Var("name", arg)
	}

	ctx := context.Background()
	err = client.Run(ctx, req, &res)

	return res, err
}
