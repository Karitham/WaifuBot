package query

import (
	"context"

	"github.com/machinebox/graphql"
)

// AnimeSearchStruct handles data from CharByName queries
type AnimeSearchStruct struct {
	Media struct {
		SiteURL   string `json:"siteUrl"`
		Status    string `json:"status"`
		MeanScore int    `json:"meanScore"`
		Title     struct {
			Romaji string `json:"romaji"`
		} `json:"title"`
		CoverImage struct {
			Medium string `json:"medium"`
		} `json:"coverImage"`
		Description string `json:"description"`
	} `json:"Media"`
}

// SearchAnime makes a query to the anilist API based on the name you input
func SearchAnime(name string) (response AnimeSearchStruct, err error) {
	// build request
	req := graphql.NewRequest(`
	query ($name: String) {
		Media(search: $name, type: ANIME) {
		  siteUrl
		  status
		  meanScore
		  title {
			romaji
		  }
		  coverImage {
			medium
		  }
		  description
		}
	  }
	`)

	// Add variables
	req.Var("name", name)

	// Make request
	err = graphql.NewClient(graphURL).Run(context.Background(), req, &response)
	return
}
