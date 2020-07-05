package query

import (
	"context"

	"github.com/machinebox/graphql"
)

// RandomCharStruct handles data for RandomChar function
type RandomCharStruct struct {
	Page struct {
		Characters []struct {
			ID      int64  `json:"id"`
			SiteURL string `json:"siteUrl"`
			Image   struct {
				Medium string `json:"medium"`
				Large  string `json:"large"`
			}
			Name struct {
				Full string `json:"full"`
			}
		}
	}
}

// RandomChar outputs the character you want based on their number on the page list
func RandomChar(id int) (RandomCharStruct, error) {
	var res RandomCharStruct
	graphURL := "https://graphql.anilist.co"
	client := graphql.NewClient(graphURL)
	req := graphql.NewRequest(`
		query ($pageNumber: Int) {
			Page(perPage: 1, page: $pageNumber) {
				characters(sort: FAVOURITES_DESC) {
					id
					siteUrl
					image {
						medium
						large
					}
					name {
						full
					}
				}
				}
			}
		`)

	req.Var("pageNumber", id)
	ctx := context.Background()
	err := client.Run(ctx, req, &res)

	return res, err
}
