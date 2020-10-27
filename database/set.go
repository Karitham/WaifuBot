package database

import (
	"context"
	"log"

	"github.com/diamondburned/arikawa/discord"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// FavoriteStruct represents how to send data to the database
type FavoriteStruct struct {
	UserID   discord.UserID `bson:"_id"`
	Favorite CharLayout
}

// NewQuote represent the data needed to change user quote
type NewQuote struct {
	UserID discord.UserID
	Quote  string
}

// SetFavorite adds a waifu to the user each time he has a new one
func (input FavoriteStruct) SetFavorite() {
	opts := options.FindOneAndUpdate().SetUpsert(true)
	_, err := collection.FindOneAndUpdate(
		context.TODO(),
		bson.M{
			"_id": input.UserID,
		},
		bson.M{
			"$set": bson.M{"Favourite": input.Favorite},
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
