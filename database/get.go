package database

import (
	"context"
	"log"
	"time"

	"github.com/andersfylling/disgord"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// UserDataStruct is a representation of the data inside the database, it's used to retrieve data
type UserDataStruct struct {
	ID            int          `bson:"_id"`
	Quote         string       `bson:"Quote,omitempty"`
	Favourite     CharLayout   `bson:"Favourite,omitempty"`
	ClaimedWaifus int          `bson:"ClaimedWaifus,omitempty"`
	Date          time.Time    `bson:"Date,omitempty"`
	Waifus        []CharLayout `bson:"Waifus,omitempty"`
}

// CharLayout is how each character is stored
type CharLayout struct {
	ID    int64  `bson:"ID"`
	Name  string `bson:"Name"`
	Image string `bson:"Image"`
}

// ViewUserData returns a list of waifus the user has collected
func ViewUserData(id disgord.Snowflake) (userData UserDataStruct) {
	bytesWaifu, err := collection.FindOne(context.TODO(), bson.M{"_id": id}).DecodeBytes()
	if err != mongo.ErrNoDocuments {
		err := bson.Unmarshal(bytesWaifu, &userData)
		if err != bson.ErrDecodeToNil && err != nil {
			log.Println(err)
		}
	}
	return
}

// CheckWaifuStruct is what data to send to check if a waifu is owned from another user in the database
type CheckWaifuStruct struct {
	UserID disgord.Snowflake `bson:"_id"`
	CharID int
}

// CheckWaifuData is what data to send to check if a waifu is owned from another user in the database
func (input CheckWaifuStruct) CheckWaifu() (WaifuExists bool) {
	var decoded bson.M

	// Find the character and delete it
	var err = collection.FindOne(
		context.TODO(),
		bson.D{
			primitive.E{Key: "_id", Value: input.UserID},
			primitive.E{Key: "Waifus.ID", Value: input.CharID},
		},
	).Decode(&decoded) // If the database found something, returns true
	if err != nil {
		fmt.Println("There was an error when checking for a waifu :", err)
		return false
	}
	return true
}
