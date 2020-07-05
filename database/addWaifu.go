package database

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// InputStruct represents how to send data to the database
type InputStruct struct {
	UserID interface{} `bson:"_id"`
	Date   time.Time   `bson:"Date"`
	Waifu  int64       `bson:"Waifus"`
}

// AddWaifu adds a waifu to the user each time he has a new one
func AddWaifu(input InputStruct) {
	var decoded bson.M
	opts := options.FindOneAndUpdate().SetUpsert(true)
	err := Collection.FindOneAndUpdate(
		context.TODO(),
		bson.M{
			"_id": input.UserID,
		},
		bson.M{
			"$push": bson.M{"Waifus": input.Waifu},
			"$set":  bson.M{"Date": input.Date},
		},
		opts,
	).Decode(&decoded)
	switch err != nil {
	case err == mongo.ErrNoDocuments:
	default:
		fmt.Println(err)
	}

}
