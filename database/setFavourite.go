package database

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// FavouriteStruct represents how to send data to the database
type FavouriteStruct struct {
	UserID    interface{} `bson:"_id"`
	Favourite struct {
		ID    interface{} `bson:"ID"`
		Name  string      `bson:"Name"`
		Image string      `bson:"Image"`
	}
}

// SetFavourite adds a waifu to the user each time he has a new one
func SetFavourite(input FavouriteStruct) {
	var decoded bson.M
	opts := options.FindOneAndUpdate().SetUpsert(true)
	err := Collection.FindOneAndUpdate(
		context.TODO(),
		bson.M{
			"_id": input.UserID,
		},
		bson.M{
			"$set": bson.M{"Favourite": input.Favourite},
		},
		opts,
	).Decode(&decoded)
	if err != mongo.ErrNoDocuments && err != nil {
		fmt.Println(err)
	}
}
