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

// CharSearch makes a query to the anilist API based on the name//ID you input
func CharSearch(args []string) (CharSearchStruct, error) {
	var res CharSearchStruct

	// build query
	graphURL := "https://graphql.anilist.co"
	client := graphql.NewClient(graphURL)
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
