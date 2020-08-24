package query

import (
	"context"

	"github.com/machinebox/graphql"
)

// MangaSearchStruct handles data from CharByName queries
type MangaSearchStruct struct {
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

// SearchManga makes a query to the Anilist API based on the name you input
func SearchManga(name string) (response MangaSearchStruct, err error) {
	// build request
	req := graphql.NewRequest(`
	query ($name: String) {
		Media(search: $name, type: MANGA) {
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