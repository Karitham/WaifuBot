package anilist

import (
	"math/rand"
	"sync"
	"time"

	"github.com/machinebox/graphql"
)

const (
	Color   = 0x02a9ff
	IconURL = "https://anilist.co/img/icons/favicon-32x32.png"
)

type Anilist struct {
	c             *graphql.Client
	seed          rand.Source64
	URL           string
	MaxChars      int64
	internalCache map[string]querier[any]
}

type querier[T any] struct {
	*sync.Mutex
	cache map[any]any
}

func New() *Anilist {
	const graphURL = "https://graphql.anilist.co"

	return &Anilist{
		URL:      graphURL,
		c:        graphql.NewClient(graphURL),
		MaxChars: 100_000,
		seed:     rand.New(rand.NewSource(time.Now().Unix())),
		internalCache: map[string]querier[any]{
			"random": {
				cache: make(map[any]any),
				Mutex: &sync.Mutex{},
			},
		},
	}
}
