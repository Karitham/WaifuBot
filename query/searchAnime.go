package query

import (
	"context"

	"github.com/machinebox/graphql"
)

// AnimeSearchStruct handles data from CharByName queries
type AnimeSearchStruct struct {
	Media struct {
		ID      int    `json:"id"`
		SiteURL string `json:"siteUrl"`
		Title   struct {
			Romaji string `json:"romaji"`
		}
		CoverImage struct {
			Large string `json:"large"`
		}
		Status       string `json:"status"`
		Episodes     int    `json:"episodes"`
		Description  string `json:"description"`
		AverageScore int    `json:"averageScore"`
		IsAdult      bool   `json:"isAdult"`
	}
}

// SearchAnime makes a query to the anilist API based on the name//ID you input
func (args CharSearchInput) SearchAnime() (AnimeSearchStruct, error) {
	var res AnimeSearchStruct

	// build query
	client := graphql.NewClient(GraphURL)
	req := graphql.NewRequest(`
	query ($query: String, $type: MediaType) {
		Media(search: $query, type: $type) {
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
	`)

	// Add variable
	if args.ID != 0 {
		req.Var("id", args.ID)
	} else {
		req.Var("query", args.Name)
	}

	ctx := context.Background()
	err := client.Run(ctx, req, &res)
	return res, err
}
