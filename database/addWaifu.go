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
	UserID    interface{} `bson:"_id"`
	Date      time.Time   `bson:"Date"`
	Favourite struct {
		ID    int64  `bson:"ID"`
		Name  string `bson:"Name"`
		Image string `bson:"Image"`
	}
	WaifuList struct {
		ID    int64  `bson:"ID"`
		Name  string `bson:"Name"`
		Image string `bson:"Image"`
	}
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
			"$set":  bson.M{"Date": input.Date},
			"$push": bson.M{"Waifus": input.WaifuList},
		},
		opts,
	).Decode(&decoded)
	if err != mongo.ErrNoDocuments && err != nil {
		fmt.Println(err)
	}
}
