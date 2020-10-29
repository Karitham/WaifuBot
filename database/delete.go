package database

import (
	"context"

	"github.com/diamondburned/arikawa/discord"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// CharDelStruct is what data to send to remove a waifu from the database
type CharDelStruct struct {
	UserID discord.UserID `bson:"_id"`
	CharID uint
}

// DelChar removes a waifu from the database
func (input CharDelStruct) DelChar() (ChangedData UserDataStruct, err error) {
	// Find the character and delete it
	err = collection.FindOneAndUpdate(
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
	).Decode(&ChangedData)
	if err != nil {
		return UserDataStruct{}, err
	}
	return ChangedData, nil
}
