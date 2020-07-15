package database

import (
	"context"
	"fmt"

	"github.com/andersfylling/disgord"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// DelWaifuStruct is what data to send to remove a waifu from the database
type DelWaifuStruct struct {
	UserID  disgord.Snowflake `bson:"_id"`
	WaifuID int
}

// DelWaifu removes a waifu from the database
func DelWaifu(input DelWaifuStruct) {
	var decoded bson.M

	// Find the character and delete it
	err := Collection.FindOneAndUpdate(
		context.TODO(),
		bson.M{
			"_id": input.UserID,
		},
		bson.M{"$pull": bson.M{"Waifus": bson.M{"ID": input.WaifuID}}},
	).Decode(&decoded)
	if err != mongo.ErrNoDocuments && err != nil {
		fmt.Println(err)
	}
	fmt.Println(decoded)
}
