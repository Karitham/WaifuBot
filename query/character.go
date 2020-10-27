package query

import (
	"context"

	"github.com/machinebox/graphql"
)

const graphURL string = "https://graphql.anilist.co"

// CharSearchStruct handles data from CharByName queries
type CharSearchStruct struct {
	Character CharacterStruct
}

// CharStruct handles data for RandomChar function
type CharStruct struct {
	Page struct {
		Characters []CharacterStruct
	}
}

// CharacterStruct represent character object
type CharacterStruct struct {
	ID      uint   `json:"id"`
	SiteURL string `json:"siteUrl"`
	Image   struct {
		Large string `json:"large"`
	}
	Name struct {
		Full string `json:"full"`
	}
	Media struct {
		Nodes []struct {
			Title struct {
				Romaji string `json:"romaji"`
			}
		}
	}
}

// CharSearchInput is used to input the arguments you want to search
type CharSearchInput struct {
	ID   int
	Name string
}

// CharSearch makes a query to the Anilist API based on the name/ID you input
func CharSearch(input CharSearchInput) (response CharSearchStruct, err error) {
	// build query
	req := graphql.NewRequest(`
	query ($id: Int, $name: String) {
		Character(id: $id, search: $name, sort: SEARCH_MATCH) {
		  id
		  siteUrl
		  name {
			full
		  }
		  image {
			large
		  }
		  media(perPage: 1, sort: POPULARITY_DESC) {
			nodes {
			  title {
				romaji
			  }
			}
		  }
		}
	  }
	`)

	if input.ID != 0 {
		req.Var("id", input.ID)
	} else {
		req.Var("name", input.Name)
	}
	err = graphql.NewClient(graphURL).Run(context.Background(), req, &response)
	return
}

// CharSearchByPopularity outputs the character you want based on their number on the page list
func CharSearchByPopularity(id int) (response CharStruct, err error) {
	// Create request
	req := graphql.NewRequest(`
	query ($pageNumber: Int) {
		Page(perPage: 1, page: $pageNumber) {
		  characters(sort: FAVOURITES_DESC) {
			id
			siteUrl
			image {
			  large
			}
			name {
			  full
			}
			media(perPage: 1, sort: POPULARITY_DESC) {
			  nodes {
				title {
				  romaji
				}
			  }
			}
		  }
		}
	  }
	`)

	req.Var("pageNumber", id)

	// Make request
	err = graphql.NewClient(graphURL).Run(context.Background(), req, &response)
	return
}
