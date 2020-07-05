package database

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// OutputStruct is a representation of the data inside the database, it's used to retrieve data
type OutputStruct struct {
	UserID    int64     `bson:"_id"`
	Date      time.Time `bson:"Date"`
	Favourite struct {
		FavID    int    `bson:"ID"`
		FavName  string `bson:"Name"`
		FavImage string `bson:"Image"`
	}
	WaifuList struct {
		WListID    int    `bson:"ID"`
		WListName  string `bson:"Name"`
		WListImage string `bson:"Image"`
	}
}

// SeeWaifus returns a list of waifus the user has collected
func SeeWaifus(id interface{}) OutputStruct {
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
