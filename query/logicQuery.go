package query

import (
	"fmt"
	"math/rand"
	"time"
)

// MakeRQ make the request and setups correctly the page total
func MakeRQ(maxCharQuery int) RandomCharStruct {
	res, err := RandomChar(random(maxCharQuery))
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
