package database

import (
	"context"
	"log"
	"time"

	"github.com/diamondburned/arikawa/discord"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// UserDataStruct is a representation of the data inside the database, it's used to retrieve data
type UserDataStruct struct {
	ID            uint         `bson:"_id"`
	Quote         string       `bson:"Quote,omitempty"`
	Favorite      CharLayout   `bson:"Favourite,omitempty"`
	ClaimedWaifus int          `bson:"ClaimedWaifus,omitempty"`
	Date          time.Time    `bson:"Date,omitempty"`
	Waifus        []CharLayout `bson:"Waifus,omitempty"`
}

// CharLayout is how each character is stored
type CharLayout struct {
	ID    uint   `bson:"ID"`
	Name  string `bson:"Name"`
	Image string `bson:"Image"`
}

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
func VerifyWaifu(wID uint, uID uint) (WaifuExists bool, userData UserDataStruct) {
	return !(collection.FindOne(
		context.TODO(),
		bson.D{
			primitive.E{Key: "_id", Value: uID},
			primitive.E{Key: "Waifus.ID", Value: wID},
		},
	).Decode(&userData) == mongo.ErrNoDocuments), userData
}
