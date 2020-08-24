package query

import (
	"context"

	"github.com/machinebox/graphql"
)

// TrendingMediaStruct handles data from the queries
type TrendingMediaStruct struct {
	Page struct {
		Media []struct {
			Title struct {
				Romaji string `json:"romaji"`
			}
		}
	}
}

// TrendingMediaQuery makes a query to the AniList GraphQL API to scrape the 10 best trending animes right now
func TrendingMediaQuery(format string) (response TrendingMediaStruct, err error) {
	// Build request
	req := graphql.NewRequest(`
	query ($type: MediaType = ANIME) {
		Page(perPage: 10, page: 1) {
		  media(type: $type, sort: [TRENDING_DESC, POPULARITY_DESC]) {
			title {
			  romaji
			}
		  }
		}
	  }
	`)
	req.Var("type", format)

	err = graphql.NewClient(graphURL).Run(context.Background(), req, &response)
	return
}
