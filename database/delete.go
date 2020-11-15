package database

import (
	"context"

	"github.com/diamondburned/arikawa/v2/discord"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// DelChar removes a waifu from the database
func (c CharID) DelChar(uID discord.UserID) (ChangedData UserDataStruct, err error) {
	// Find the character and delete it
	err = collection.FindOneAndUpdate(
		context.TODO(),
		bson.D{
			primitive.E{Key: "_id", Value: uID},
			primitive.E{Key: "Waifus.ID", Value: c},
		},
		bson.M{"$pull": bson.M{
			"Waifus": bson.M{
				"ID": c,
			},
		},
		},
	).Decode(&ChangedData)
	if err != nil {
		return UserDataStruct{}, err
	}
	return ChangedData, nil
}
