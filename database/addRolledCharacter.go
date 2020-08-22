package database

import (
	"context"
	"fmt"
	"time"

	"github.com/andersfylling/snowflake/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// InputChar represents how to send data to the database
type InputChar struct {
	UserID   snowflake.Snowflake `bson:"_id"`
	Date     time.Time           `bson:"Date"`
	CharList CharLayout
}

// AddChar adds a waifu to the user each time he has a new one
func (input InputChar) AddChar() {
	var decoded bson.M
	opts := options.FindOneAndUpdate().SetUpsert(true)
	err := Collection.FindOneAndUpdate(
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
		fmt.Println(err)
	}
}
