package database

import (
	"context"
	"fmt"
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
		er := bson.Unmarshal(bytesWaifu, &userData)
		if er != bson.ErrDecodeToNil && err != nil {
			fmt.Println(er)
		}
	}
	return
}
