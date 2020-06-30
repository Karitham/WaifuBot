package query

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/machinebox/graphql"
)

// RespCharType : Struct for handling character requests from anilists API
type RespCharType struct {
	Page struct {
		Characters []struct {
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
}

// Char : makes a character query
func Char(id int) (RespCharType, error) {
	var res RespCharType
	graphURL := "https://graphql.anilist.co"

	client := graphql.NewClient(graphURL)
	req := graphql.NewRequest(`
	query ($pageNumber : Int) {
		Page(perPage: 1, page: $pageNumber) {
			characters (sort : FAVOURITES_DESC ){
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

	req.Var("pageNumber", id)
	ctx := context.Background()
	err := client.Run(ctx, req, &res)

	return res, err
}

// MakeRQ make the request and setups correctly the page total
func MakeRQ(maxCharQuery int) RespCharType {
	res, err := Char(random(maxCharQuery))
	if err != nil {
		fmt.Println(err)
	}
	return res
}

// random : search the char by ID entered in discord
func random(maxCharQuery int) int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	random := r.Intn(maxCharQuery)
	return random
}
