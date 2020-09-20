package query

import (
	"context"

	"github.com/machinebox/graphql"
)

// UserQueryStruct represent the struct of the user query
type UserQueryStruct struct {
	User struct {
		Avatar struct {
			Medium string `json:"medium"`
		} `json:"avatar"`
		About      string `json:"about"`
		Name       string `json:"name"`
		Statistics struct {
			Anime struct {
				EpisodesWatched int64 `json:"episodesWatched"`
			} `json:"anime"`
			Manga struct {
				ChaptersRead int64 `json:"chaptersRead"`
			} `json:"manga"`
		} `json:"statistics"`
		SiteURL string `json:"siteUrl"`
	} `json:"User"`
}

// User queries the anilist user
func User(name string) (response UserQueryStruct, err error) {
	req := graphql.NewRequest(`
	query ($name: String) {
	User(search: $name) {
		avatar {
		medium
		}
		about
		name
		statistics {
		anime {
			episodesWatched
		}
		manga {
			chaptersRead
		}
		}
		siteUrl
	}
	}
	`)
	req.Var("name", name)
	err = graphql.NewClient(graphURL).Run(context.Background(), req, &response)
	return
}
