package anilist

import (
	"context"

	"github.com/machinebox/graphql"
)

// UserQueryStruct represent the struct of the user query
type UserQueryStruct struct {
	User struct {
		SiteURL string `json:"siteUrl"`
	} `json:"User"`
}

// User queries the anilist user
func User(name string) (response UserQueryStruct, err error) {
	req := graphql.NewRequest(`
	query ($name: String) {
	User(search: $name) {
		siteUrl
		}
	}
	`)
	req.Var("name", name)
	err = graphql.NewClient(graphURL).Run(context.Background(), req, &response)
	return
}
