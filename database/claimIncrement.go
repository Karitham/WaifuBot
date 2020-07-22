package database

import (
	"context"
	"fmt"

	"github.com/andersfylling/disgord"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// ClaimIncrementStruct is an input struct for the ClaimIncrement function
type ClaimIncrementStruct struct {
	UserID    disgord.Snowflake
	Increment int
}

// ClaimIncrement is used to increment the number of claims the user has done
func (input ClaimIncrementStruct) ClaimIncrement() {
	opts := options.FindOneAndUpdate().SetUpsert(true)
	_, err := Collection.FindOneAndUpdate(
		context.TODO(),
		bson.M{
			"_id": input.UserID,
		},
		bson.M{
			"$inc": bson.M{"ClaimedWaifus": input.Increment},
		},
		opts,
	).DecodeBytes()
	if err != mongo.ErrNoDocuments && err != nil {
		fmt.Println(err)
	}
}
