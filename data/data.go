package data

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// UserBson represents how the data is stored in the database
type UserBson struct {
	UserID int64  `json:"UserID"`
	Date   string `json:"Date"`
	Waifus []int  `json:"Waifus"`
}

// InputStruct serves as an input scheme
type InputStruct struct {
	UserID int `json:"UserID"`
	Waifu  int `json:"Waifu"`
}

var client *mongo.Client

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

	fmt.Println("Connected to MongoDB!")
}

// Store handle the incoming data for the database
func Store(input InputStruct) {
	collection := client.Database("waifu").Collection("users")

	insertResult, err := collection.InsertOne(context.TODO(), input)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted a single document: ", insertResult.InsertedID)
}

// ResetDB resets the database
func ResetDB() {
	client.Database("waifu").Drop(context.TODO())
	fmt.Println("whole database emptied")
}
