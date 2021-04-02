package anilist

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
		IsAdult   bool   `json:"isAdult"`
		Title     struct {
			Romaji string `json:"romaji"`
		} `json:"title"`
		CoverImage struct {
			Medium string `json:"medium"`
		} `json:"coverImage"`
		Description string `json:"description"`
	} `json:"Media"`
}

// MediaSearch makes a query to the Anilist API to look after the wanted anime/manga (ID or Name accepted)
func MediaSearch(name string, format string) (response MediaSearchStruct, err error) {
	// build request
	req := graphql.NewRequest(`
	query ($name: String, $type: MediaType) {
		Media(search: $name, type: $type) {
		  siteUrl
		  status
		  meanScore
		  isAdult	
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
