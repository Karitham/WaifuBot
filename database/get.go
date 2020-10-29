package database

import (
	"context"
	"log"

	"github.com/diamondburned/arikawa/discord"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// ViewUserData returns a list of waifus the user has collected
func ViewUserData(id discord.UserID) (userData UserDataStruct, err error) {
	bytesWaifu, err := collection.FindOne(context.TODO(), bson.M{"_id": id}).DecodeBytes()
	if err != mongo.ErrNoDocuments {
		err := bson.Unmarshal(bytesWaifu, &userData)
		if err != bson.ErrDecodeToNil && err != nil {
			log.Println(err)
		}
	} else {
		return UserDataStruct{}, err
	}
	return
}

// VerifyWaifu verifies if the mentioned account has got the Waifu he asked for.
func (wID CharID) VerifyWaifu(uID discord.UserID) (WaifuExists bool, userData UserDataStruct) {
	return !(collection.FindOne(
		context.TODO(),
		bson.D{
			primitive.E{Key: "_id", Value: uID},
			primitive.E{Key: "Waifus.ID", Value: wID},
		},
	).Decode(&userData) == mongo.ErrNoDocuments), userData
}
