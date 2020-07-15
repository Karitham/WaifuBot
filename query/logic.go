package query

import (
	"fmt"
	"math/rand"
	"time"
)

// RandomCharQuery make the request and setups correctly the page total
func RandomCharQuery(maxCharQuery int) CharStruct {
	// set seeds & roll
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	random := r.Intn(maxCharQuery)

	// get the response
	res, err := RandomChar(random)
	if err != nil {
		fmt.Println(err)
	}
	return res
}
