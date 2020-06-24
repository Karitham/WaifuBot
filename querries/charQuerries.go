//Package querries define what querries you can use and the response structs
package querries

import (
	"context"
	"fmt"

	"github.com/machinebox/graphql"
)

// ResponseChar : Struct for handling character requests from anilists API
type ResponseChar struct {
	Character struct {
		ID    int64 `json:"id"`
		Image struct {
			Large  string `json:"large"`
			Medium string `json:"medium"`
		}
		Name struct {
			First  string `json:"first"`
			Full   string `json:"full"`
			Last   string `json:"last"`
			Native string `json:"native"`
		}
	}
}

var graphURL = "https://graphql.anilist.co"
var client = graphql.NewClient(graphURL)

// Char : makes a character query
func Char(id int) ResponseChar {
	req := graphql.NewRequest(`
	query ($id: Int) {
		Character(id: $id) {
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
	`)
	req.Var("id", id)
	ctx := context.Background()
	var RespChar ResponseChar
	if err := client.Run(ctx, req, &RespChar); err != nil {
		fmt.Println(err)
	}
	return RespChar
}
