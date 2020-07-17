package query

import "github.com/shurcooL/graphql"

// GraphURL represent the URL which is accessed when making a query
const GraphURL = "https://graphql.anilist.co"

// init client
var client = graphql.NewClient(GraphURL, nil)
