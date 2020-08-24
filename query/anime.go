package query

import (
	"context"

	"github.com/machinebox/graphql"
)

// MediaSearchStruct handles data from either manga or anime searches
type MediaSearchStruct struct {
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

// SearchMedia makes a query to the anilist API based on the name you input
func SearchMedia(name string, format string) (response MediaSearchStruct, err error) {
	// build request
	req := graphql.NewRequest(`
	query ($name: String, $type: MediaType = ANIME) {
		Media(search: $name, type: $type) {
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
	req.Var("type", format)

	// Make request
	err = graphql.NewClient(graphURL).Run(context.Background(), req, &response)
	return
}
