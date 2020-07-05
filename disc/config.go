package disc

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// configFromJSON reads config from file
func configFromJSON(file string) ConfigT {
	var config ConfigT
	body, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println(err)
	}
	json.Unmarshal(body, &config)
	return config
}
