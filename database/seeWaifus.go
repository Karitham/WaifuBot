package database

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

// OutputStruct is a representation of the data inside the database, it's used to retrieve data
type OutputStruct struct {
	UserID int       `bson:"_id"`
	Date   time.Time `bson:"Date"`
	Waifus []int64   `bson:"Waifus"`
}

// SeeWaifus returns a list of waifus the user has collected
func SeeWaifus(id interface{}) OutputStruct {
	var output OutputStruct
	bytesWaifu, err := Collection.FindOne(context.TODO(), bson.M{"_id": id}).DecodeBytes()
	if err != nil {
		fmt.Println(err)
	}
	err = bson.Unmarshal(bytesWaifu, &output)
	if err != nil {
		fmt.Println(err)
	}
	return output
}
