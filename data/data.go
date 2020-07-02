package data
<<<<<<< HEAD

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
func Store(input UserBson) {
	collection := client.Database("waifu").Collection("users")
	var test bson.M
	err := collection.FindOne(context.TODO(), bson.D{{"UserID", input.UserID}}).Decode(&test)
	if err != nil {
		collection.InsertOne(context.TODO(), input)
	}
}

// ResetDB resets the database
func ResetDB() {
	client.Database("waifu").Drop(context.TODO())
	fmt.Println("whole database emptied")
}
=======
>>>>>>> 5a6f76adb141c254e863f291a31743c241bb2280
