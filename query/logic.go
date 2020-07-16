package query

import (
	"fmt"
	"math/rand"
	"time"
)

// RandomCharQuery make a random rq based on the maxCharQuery
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
