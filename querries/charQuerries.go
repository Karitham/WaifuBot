//Package querries define what querries you can use and the response structs
package querries

import (
	"context"
	"fmt"

	"github.com/machinebox/graphql"
)

// ResponseChar : Struct for handling character requests from anilists API
type ResponseChar struct {
	Page struct {
		Characters []struct {
			ID    int
			Image struct {
				Large  string
				Medium string
			}
			Name struct {
				First  string
				Full   string
				Last   string
				Native string
			}
		}
		PageInfo struct {
			LastPage int
		}
	}
}

var graphURL = "https://graphql.anilist.co"
var client = graphql.NewClient(graphURL)

// Char : makes a character query
func Char(id int) ResponseChar {
	req := graphql.NewRequest(`
	query ($pageNumber : Int) {
		Page(perPage: 1, page: $pageNumber) {
			pageInfo {
				lastPage
			}
			characters {
					id
					image {
						large
						medium
					}
					name {
						first
						last
						full
						native
						}
					}
				
		}
	}
	`)
	req.Var("pageNumber", 5849)
	ctx := context.Background()
	var res ResponseChar
	if err := client.Run(ctx, req, &res); err != nil {
		fmt.Println(err)
	}
	return res
}
