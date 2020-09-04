package database

import (
	"context"
	"log"

	"github.com/andersfylling/disgord"
	"github.com/andersfylling/snowflake/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// FavouriteStruct represents how to send data to the database
type FavouriteStruct struct {
	UserID    snowflake.Snowflake `bson:"_id"`
	Favourite CharLayout
}

// NewQuote represent the data needed to change user quote
type NewQuote struct {
	UserID disgord.Snowflake
	Quote  string
}

// SetFavourite adds a waifu to the user each time he has a new one
func (input FavouriteStruct) SetFavourite() {
	opts := options.FindOneAndUpdate().SetUpsert(true)
	_, err := collection.FindOneAndUpdate(
		context.TODO(),
		bson.M{
			"_id": input.UserID,
		},
		bson.M{
			"$set": bson.M{"Favourite": input.Favourite},
		},
		opts,
	).DecodeBytes()
	if err != mongo.ErrNoDocuments && err != nil {
		log.Println(err)
	}
}

// SetQuote set the user quote on his profile
func (input NewQuote) SetQuote() {
	opts := options.FindOneAndUpdate().SetUpsert(true)
	_, err := collection.FindOneAndUpdate(
		context.TODO(),
		bson.M{
			"_id": input.UserID,
		},
		bson.M{
			"$set": bson.M{"Quote": input.Quote},
		},
		opts,
	).DecodeBytes()
	if err != mongo.ErrNoDocuments && err != nil {
		log.Println(err)
	}
}
