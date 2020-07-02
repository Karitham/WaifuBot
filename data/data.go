package data

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// InputStruct represents how to send data to the database
type InputStruct struct {
	UserID interface{} `bson:"_id"`
	Date   time.Time   `bson:"Date"`
	Waifu  int64       `bson:"Waifus"`
}

// OutputStruct is a representation of the data inside the database, it's used to retrieve data
type OutputStruct struct {
	UserID int       `bson:"_id"`
	Date   time.Time `bson:"Date"`
	Waifu  []int64   `bson:"Waifus"`
}

var client *mongo.Client
var collection *mongo.Collection
var ctx = context.TODO()

// InitDB is used to start the database
func InitDB() {

	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	// Connect to MongoDB
	var err error
	client, err = mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	// Connection successfull
	collection = client.Database("waifubot").Collection("waifus")
	fmt.Println("Connected to MongoDB!")
}

// Drop a user via USER ID
func Drop(input InputStruct) mongo.DeleteResult {
	deleteOneResult, err := collection.DeleteOne(ctx, bson.M{"UserID": input.UserID})
	if err != nil {
		fmt.Println(err)
	}
	return *deleteOneResult
}

// AddWaifu adds a waifu to the user each time he has a new one
func AddWaifu(input InputStruct) {
	var decoded bson.M
	opts := options.FindOneAndUpdate().SetUpsert(true)
	err := collection.FindOneAndUpdate(
		ctx,
		bson.M{
			"_id": input.UserID,
		},
		bson.M{
			"$push": bson.M{"Waifus": input.Waifu},
			"$set":  bson.M{"Date": input.Date},
		},
		opts,
	).Decode(&decoded)
	if err != nil {
		fmt.Println(err)
	}
}

// SeeWaifus returns a list of waifus the user has collected
func SeeWaifus(id interface{}) []int64 {
	var output OutputStruct
	bytesWaifu, err := collection.FindOne(ctx, bson.M{"_id": id}).DecodeBytes()
	if err != nil {
		fmt.Println(err)
	}
	err = bson.Unmarshal(bytesWaifu, &output)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(output.Waifu)
	return output.Waifu
}
