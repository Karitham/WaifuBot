package query

import (
	"context"

	"github.com/shurcooL/graphql"
)

// SearchAnimeStruct handles data from CharByName queries
type SearchAnimeStruct struct {
	Media struct {
		ID      int    `graphql:"id"`
		SiteURL string `graphql:"siteUrl"`
		Title   struct {
			Romaji string `graphql:"romaji"`
		} `graphql:"title"`
		CoverImage struct {
			Large string `graphql:"large"`
		} `graphql:"coverImage"`
		Status       string `graphql:"status"`
		Episodes     int    `graphql:"episodes"`
		Description  string `graphql:"description"`
		AverageScore int    `graphql:"averageScore"`
		IsAdult      bool   `graphql:"isAdult"`
	} `graphql:"Media(search: $query, type: ANIME)"`
}

// SearchAnime makes a query to the anilist API based on the name//ID you input
func SearchAnime(args string) (SearchAnimeStruct, error) {

	// Set variables
	variables := map[string]interface{}{
		"query": graphql.String(args),
	}

	// q represents the data sent back by the query
	var q SearchAnimeStruct

	// Query
	err := client.Query(context.Background(), &q, variables)

	return q, err
}
