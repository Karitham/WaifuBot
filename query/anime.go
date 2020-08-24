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
		} `json:"title"`
		CoverImage struct {
			Large string `json:"large"`
		} `json:"coverImage"`
		Status       string `json:"status"`
		Episodes     int    `json:"episodes"`
		Description  string `json:"description"`
		AverageScore int    `json:"averageScore"`
		IsAdult      bool   `json:"isAdult"`
	} `json:"Media"`
}

// SearchAnime makes a query to the anilist API based on the name you input
func SearchAnime(name string) (response AnimeSearchStruct, err error) {
	// build request
	req := graphql.NewRequest(`
	query ($name: String) {
		Media(search: $name, type: ANIME) {
		  id
		  siteUrl
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

	// Add variables
	req.Var("name", name)

	// Make request
	err = graphql.NewClient(graphURL).Run(context.Background(), req, &response)
	return
}
