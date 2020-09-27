package database

import (
	"context"
	"log"

	"github.com/diamondburned/arikawa/discord"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// CharDelStruct is what data to send to remove a waifu from the database
type CharDelStruct struct {
	UserID discord.UserID `bson:"_id"`
	CharID int
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

// DropUser a user via USER ID
func (input InputChar) DropUser() (deletedResult *mongo.DeleteResult) {
	deletedResult, err := collection.DeleteOne(context.TODO(), bson.M{"UserID": input.UserID})
	if err != nil {
		log.Println(err)
	}
	return
}
