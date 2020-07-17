package query

import (
	"context"

	"github.com/shurcooL/graphql"
)

// TrendingSearchStruct handles data from the queries
type TrendingSearchStruct struct {
	Page struct {
		Media []struct {
			ID      int64  `graphql:"id"`
			SiteURL string `graphql:"siteUrl"`
			Title   struct {
				UserPreferred string `graphql:"userPreferred"`
			} `graphql:"title"`
			CoverImage struct {
				Large string `graphql:"large"`
			} `graphql:"coverImage"`
			AverageScore int `graphql:"averageScore"`
			Popularity   int `graphql:"popularity"`
		} `graphql:"media(type: ANIME, sort: [TRENDING_DESC, POPULARITY_DESC])"`
	} `graphql:"Page(perPage: 10, page: $page)"`
}

// TrendingSearch makes a query to the AniList GraphQL API to scrape the 10 best trending animes right now
func TrendingSearch() (TrendingSearchStruct, error) {
	// insert variables
	variables := map[string]interface{}{
		"page": graphql.Int(0),
	}

	// q represents the data sent back by the query
	var q TrendingSearchStruct

	// Query
	err := client.Query(context.Background(), &q, variables)

	return q, err
}
