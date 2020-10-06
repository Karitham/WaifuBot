package database

import (
	"context"
	"time"

	"github.com/diamondburned/arikawa/discord"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// InputClaimedChar represents how to send data to the database
type InputClaimedChar struct {
	UserID   discord.UserID `bson:"_id"`
	CharList CharLayout
}

// InputChar represents how to send data to the database
type InputChar struct {
	UserID   discord.UserID `bson:"_id"`
	Date     time.Time      `bson:"Date"`
	CharList CharLayout
}

// AddChar adds a waifu to the user each time he has a new one
func (input InputClaimedChar) AddChar() error {
	var decoded bson.M

	opts := options.FindOneAndUpdate().SetUpsert(true)

	err := collection.FindOneAndUpdate(
		context.TODO(),
		bson.M{
			"_id": input.UserID,
		},
		bson.M{
			"$inc":  bson.M{"ClaimedWaifus": 1},
			"$push": bson.M{"Waifus": input.CharList},
		},
		opts,
	).Decode(&decoded)
	if err != mongo.ErrNoDocuments && err != nil {
		return err
	}
	return nil
}

// AddChar adds a waifu to the user each time he has a new one
func (input InputChar) AddChar() error {
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
		return err
	}
	return nil
}
