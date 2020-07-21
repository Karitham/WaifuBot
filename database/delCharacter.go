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
func (input DelWaifuStruct) DelChar() bool {
	var decoded bson.M

	// Find the character and delete it
	err := Collection.FindOneAndUpdate(
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
	switch {
	case err == mongo.ErrNoDocuments:
		return false
	case err != nil:
		fmt.Println(err)
		return false
	case err == nil:
		return true
	}
	return false
}
