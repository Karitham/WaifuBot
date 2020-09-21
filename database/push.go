package database

import (
	"context"
	"log"
	"time"

	"github.com/andersfylling/snowflake/v4"
	"github.com/diamondburned/arikawa/discord"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// InputClaimChar represents how to send data to the database
type InputClaimChar struct {
	UserID   snowflake.Snowflake `bson:"_id"`
	CharList CharLayout
}

// ClaimIncrementStruct is an input struct for the ClaimIncrement function
type ClaimIncrementStruct struct {
	UserID    snowflake.Snowflake
	Increment int
}

// InputChar represents how to send data to the database
type InputChar struct {
	UserID   discord.UserID `bson:"_id"`
	Date     time.Time      `bson:"Date"`
	CharList CharLayout
}

// AddChar adds a waifu to the user each time he has a new one
func (input InputClaimChar) AddChar() {
	var decoded bson.M
	opts := options.FindOneAndUpdate().SetUpsert(true)
	err := collection.FindOneAndUpdate(
		context.TODO(),
		bson.M{
			"_id": input.UserID,
		},
		bson.M{
			"$push": bson.M{"Waifus": input.CharList},
		},
		opts,
	).Decode(&decoded)
	if err != mongo.ErrNoDocuments && err != nil {
		log.Println(err)
	}
}

// ClaimIncrement is used to increment the number of claims the user has done
func (input ClaimIncrementStruct) ClaimIncrement() {
	opts := options.FindOneAndUpdate().SetUpsert(true)
	_, err := collection.FindOneAndUpdate(
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
		log.Println(err)
	}
}

// AddChar adds a waifu to the user each time he has a new one
func (input InputChar) AddChar() {
	var decoded bson.M
	opts := options.FindOneAndUpdate().SetUpsert(true)
	err := collection.FindOneAndUpdate(
		context.TODO(),
		bson.M{
			"_id": input.UserID,
		},
		bson.M{
			"$set":  bson.M{"Date": input.Date},
			"$push": bson.M{"Waifus": input.CharList},
		},
		opts,
	).Decode(&decoded)
	if err != mongo.ErrNoDocuments && err != nil {
		log.Println(err)
	}
}
