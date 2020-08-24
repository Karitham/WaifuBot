package database

import (
	"context"
	"fmt"

	"github.com/andersfylling/disgord"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// DelWaifuStruct is what data to send to remove a waifu from the database
type DelWaifuStruct struct {
	UserID disgord.Snowflake `bson:"_id"`
	CharID int
}

// DelChar removes a waifu from the database
func (input DelWaifuStruct) DelChar() (WaifuWasRemoved bool) {
	var decoded bson.M

	// Find the character and delete it
	err := collection.FindOneAndUpdate(
		context.TODO(),
		bson.D{
			primitive.E{Key: "_id", Value: input.UserID},
			primitive.E{Key: "Waifus.ID", Value: input.CharID},
		},
		bson.M{"$pull": bson.M{
			"Waifus": bson.M{
				"ID": input.CharID,
			},
		},
		},
	).Decode(&decoded)

	// If the database found something, returns true
	if err != nil {
		fmt.Println("There was an error delete waifu :", err)
		return false
	}
	return true
}

// DropUser a user via USER ID
func (input InputChar) DropUser() (deletedResult *mongo.DeleteResult) {
	deletedResult, err := collection.DeleteOne(context.TODO(), bson.M{"UserID": input.UserID})
	if err != nil {
		fmt.Println(err)
	}
	return
}
