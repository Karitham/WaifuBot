package query

import (
	"context"

	"github.com/machinebox/graphql"
)

// MediaSearchStruct handles data from either manga or anime searches
type MediaSearchStruct struct {
	Media struct {
		SiteURL   string `json:"siteUrl"` // Anilist URL linked to the media
		Status    string `json:"status"` // Verifies the status of the anime/manga as of today
		MeanScore int    `json:"meanScore"` // Looks after user's appreciation of the anime/manga
		IsAdult   bool   `json:"isAdult"` // Verifies if the anime/manga is for mature audiences only
		Title     struct { // Title structure
			Romaji string `json:"romaji"` // Permits the user to see the romanized name of the anime/manga.
		} `json:"title"`
		CoverImage struct { // Image structure
			Medium string `json:"medium"` // Permits the user to see the cover of the anime / manga.
		} `json:"coverImage"`
		Description string `json:"description"` // Permits to look after the description of the anime / manga.
	} `json:"Media"`
}

// SearchMedia makes a query to the Anilist API to look after the wanted anime/manga (ID or Name accepted)
func SearchMedia(name string, format string) (response MediaSearchStruct, err error) {
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
