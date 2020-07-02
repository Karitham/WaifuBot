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

// UserBson represents how the data is stored in the database
type UserBson struct {
	UserID int64     `bson:"UserID"`
	Date   time.Time `bson:"Date"`
	Waifus []int     `bson:"Waifus"`
}

var client *mongo.Client
var collection = client.Database("waifubot").Collection("waifus")

// InitDB is used to start the database
func InitDB() {

	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// Connect to MongoDB
	var err error
	client, err = mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	// Connection successfull
	fmt.Println("Connected to MongoDB!")
}

// Store handles incoming data
func Store(input UserBson) (bson.M, mongo.InsertOneResult) {
	var exists bson.M
	var insertOneReturn *mongo.InsertOneResult

	// Check if user exist in a document
	err := collection.FindOne(context.TODO(), bson.M{"UserID": input.UserID}).Decode(&exists)
	if err != nil {
		fmt.Println(err)
	}

	// Insert it if it doesn't exist
	if exists == nil {
		insertOneReturn, err = collection.InsertOne(context.TODO(), input)
		if err != nil {
			fmt.Println(err)
		}
	} else {
		fmt.Println(exists)
	}

	return exists, *insertOneReturn
}

// Drop a user via USER ID
func Drop(input UserBson) mongo.DeleteResult {
	deleteOneResult, err := collection.DeleteOne(context.TODO(), bson.M{"UserID": input.UserID})
	if err != nil {
		fmt.Println(err)
	}
	return *deleteOneResult
}
