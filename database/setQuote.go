package database

import (
	"context"
	"fmt"

	"github.com/andersfylling/disgord"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// NewQuote represent the data needed to change user quote
type NewQuote struct {
	UserID disgord.Snowflake
	Quote  string
}

// SetQuote set the user quote on his profile
func (input NewQuote) SetQuote() {
	opts := options.FindOneAndUpdate().SetUpsert(true)
	_, err := Collection.FindOneAndUpdate(
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
		fmt.Println(err)
	}
}
