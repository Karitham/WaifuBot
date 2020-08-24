package query

import (
	"context"

	"github.com/machinebox/graphql"
)

// TvTrendingStruct handles data from the queries
type TvTrendingStruct struct {
	Page struct {
		Media []struct {
			Title struct {
				Romaji string `json:"romaji"`
			}
		}
	}
}

// TrendingSearch makes a query to the AniList GraphQL API to scrape the 10 best trending animes right now
func TrendingSearch() (response TvTrendingStruct, err error) {
	// Build request
	req := graphql.NewRequest(`
	query {
		Page(perPage: 10, page: 1) {
		  media(type: ANIME, sort: [TRENDING_DESC, POPULARITY_DESC]) {
			title {
			  romaji
			}
		  }
		}
	  }
	  
	`)

	err = graphql.NewClient(graphURL).Run(context.Background(), req, &response)
	return
}
