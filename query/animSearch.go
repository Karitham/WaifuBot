package query

import (
	"context"
	"strconv"
	"strings"

	"github.com/machinebox/graphql"
)

// AnimSearchStruct handles data from CharByName queries
type AnimSearchStruct struct {
	Anime struct {
		ID      int    `json:"id"`
		SiteURL string `json:"siteUrl"`
		Title    struct {
			Romaji string `json:"romaji"`
		}
		CoverImage struct {
			Large string `json:"large"`
		}
		Status	    string `json:"status"`
		Episodes    string `json:"episodes"`
		Description string `json:"description"`
		AverageScore string `json:"averageScore"`
		IsAdult	     string `json:"isAdult"`
	}
}

// AnimSearch makes a query to the anilist API based on the name//ID you input
func AnimSearch(args []string) (CharSearchStruct, error) {
	var res AnimSearchStruct

	// build query
	graphURL := "https://graphql.anilist.co"
	client := graphql.NewClient(graphURL)
	req := graphql.NewRequest(`
	query ($query: String, $type: MediaType) {
  Page (perPage: 1) {
    media(search: $query, type: $type) {
      id
      title {
        romaji
      }
      coverImage {
        large
      }
      status
      episodes
      description
      averageScore
      isAdult
    }
  }
}
	`)
	// Parse the arguments to check if an ID or a Name was entered
	req.Var("type", "ANIME")
	arg := strings.Join(args, " ")
	id, err := strconv.Atoi(arg)

	// Add variable
	if id != 0 {
		req.Var("id", id)
	} else {
		req.Var("query", arg)
	}

	ctx := context.Background()
	err = client.Run(ctx, req, &res)

	return res, err
}
