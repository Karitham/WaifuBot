package query

import (
	"context"
	"fmt"

	"github.com/shurcooL/graphql"
)

// GetCharByPopularityStruct represents the struct for the GetCharByPopularity func
type GetCharByPopularityStruct struct {
	Page struct {
		Characters []struct {
			ID      int64  `graphql:"id"`
			SiteURL string `graphql:"siteUrl"`
			Image   struct {
				Large string `graphql:"large"`
			} `graphql:"image"`
			Name struct {
				Full string `graphql:"full"`
			} `graphql:"name"`
			Media struct {
				Nodes []struct {
					Title struct {
						Romaji string `graphql:"romaji"`
					} `graphql:"title"`
				} `graphql:"nodes"`
			} `graphql:"media(perPage: 1, sort: POPULARITY_DESC)"`
		} `graphql:"characters(sort: FAVOURITES_DESC)"`
	} `graphql:"Page(perPage: 1, page: $pageNumber)"`
}

// GetCharByPopularity is used to get characters by popularity order.
// This makes it easy to always get characters and have a randomized feature
func GetCharByPopularity(pageNumber int) GetCharByPopularityStruct {
	// q represents the data sent back by the query
	var q GetCharByPopularityStruct

	// insert variables
	variables := map[string]interface{}{
		"pageNumber": graphql.Int(pageNumber),
	}

	// Query
	err := client.Query(context.Background(), &q, variables)
	if err != nil {
		fmt.Println(err)
	}

	return q
}
