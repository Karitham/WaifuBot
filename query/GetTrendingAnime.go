package query

import (
	"context"

	"github.com/machinebox/graphql"
)

// TvTrendingStruct handles data from the queries
type TvTrendingStruct struct {
	Page struct {
		Media []struct {
			ID      int    `json:"id"`
			SiteURL string `json:"siteUrl"`
			Title   struct {
				UserPreferred string `json:"userPreferred"`
			}
			CoverImage struct {
				Large string `json:"large"`
			}
			AverageScore int `json:"averageScore"`
			Popularity   int `json:"popularity"`
		}
	}
}

// TrendingSearch makes a query to the AniList GraphQL API to scrape the 10 best trending animes right now
func TrendingSearch() (TvTrendingStruct, error) {
	var res TvTrendingStruct

	// build query
	graphURL := "https://graphql.anilist.co"
	client := graphql.NewClient(graphURL)
	req := graphql.NewRequest(`
	query ($page: Int) {
		Page(perPage: 10, page: $page) {
		  media(type: ANIME, sort: [TRENDING_DESC, POPULARITY_DESC]) {
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
	// Inject pre-made vars to get the trending animes
	req.Var("page", 1)

	// Execute code
	ctx := context.Background()
	err := client.Run(ctx, req, &res)

	return res, err
}
