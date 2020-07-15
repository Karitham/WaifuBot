package database

import (
	"context"
	"fmt"
	"time"

	"github.com/andersfylling/disgord"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// OutputStruct is a representation of the data inside the database, it's used to retrieve data
type OutputStruct struct {
	ID        int64 `bson:"_id"`
	Favourite CharLayout
	Date      time.Time `bson:"Date"`
	Waifus    []CharLayout
}

// CharLayout is how each character is stored
type CharLayout struct {
	ID    int64  `bson:"ID"`
	Name  string `bson:"Name"`
	Image string `bson:"Image"`
}

// ViewUserData returns a list of waifus the user has collected
func ViewUserData(id disgord.Snowflake) OutputStruct {
	var output OutputStruct
	bytesWaifu, err := Collection.FindOne(context.TODO(), bson.M{"_id": id}).DecodeBytes()
	if err != mongo.ErrNoDocuments {
		err = bson.Unmarshal(bytesWaifu, &output)
		if err != bson.ErrDecodeToNil && err != nil {
			fmt.Println(err)
		}
	}
	return output
}
