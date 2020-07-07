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

// InputWaifu represents how to send data to the database
type InputWaifu struct {
	UserID    snowflake.Snowflake `bson:"_id"`
	Date      time.Time           `bson:"Date"`
	WaifuList CharLayout
}

// AddWaifu adds a waifu to the user each time he has a new one
func AddWaifu(input InputWaifu) {
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
