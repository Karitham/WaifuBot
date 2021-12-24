package anilist

import (
	"math/rand"
	"time"

	"github.com/machinebox/graphql"
)

const (
	Color   = 0x19212d
	IconURL = "https://anilist.co/img/icons/favicon-32x32.png"
)

type Anilist struct {
	c        *graphql.Client
	seed     rand.Source64
	URL      string
	MaxChars int64
}

func New() *Anilist {
	const graphURL = "https://graphql.anilist.co"

	return &Anilist{
		URL:      graphURL,
		c:        graphql.NewClient(graphURL),
		MaxChars: 80_000,
		seed:     rand.New(rand.NewSource(time.Now().Unix())),
	}
}
