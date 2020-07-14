package query

import (
	"context"
	"strconv"
	"strings"

	"github.com/machinebox/graphql"
)

// CharSearchStruct handles data from CharByName queries
type tvTrendingStruct struct {
	Media struct {
		ID      int    `json:"id"`
		SiteURL string `json:"siteUrl"`
		Title    struct {
			UserPreferred string `json:"userPreferred"`
		}
		CoverImage struct {
			Large string `json:"large"`
		}
		AverageScore  int `json:"averageScore"`
    Popularity    int   `json:"popularity"`
		}
	}
}

// TrendingSearch makes a query to the AniList GraphQL API to scrape the 10 best trending animes right now.
func TrendingSearch(args []string) (CharSearchStruct, error) {
	var res tvTrendingStruct

	// build query
	graphURL := "https://graphql.anilist.co"
	client := graphql.NewClient(graphURL)
	req := graphql.NewRequest(`
	query ($page: Int, $type: MediaType, $sort: [MediaSort]) {
  Page(perPage: 10, page: $page) {
    media(type: $type, sort: $sort) {
      id
      siteUrl
      title {
        userPreferred
      }
      coverImage {
        large
      }
      averageScore
      popularity
    }
  }
}
	`)
	// Inject pre-made vars to get the trending animes.
  req.Var("page", 1)
  req.Var("type", "ANIME")
  req.Var("sort", "["TRENDING_DESC","POPULARITY_DESC"]")

  // Execute code
	ctx := context.Background()
	err = client.Run(ctx, req, &res)

	return res, err
}
