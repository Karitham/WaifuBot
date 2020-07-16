package database

import (
	"bot/config"
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Collection defines the which collection is waifu
var Collection *mongo.Collection

// Client is the mongo Client
var Client *mongo.Client

// Init is used to start the database
func Init(config config.ConfJSONStruct) {

	// Set client options
	clientOptions := options.Client().ApplyURI(config.MongoURL)
	// Connect to MongoDB
	var err error
	Client, err = mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = Client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	// Connection successfull
	Collection = Client.Database("waifubot").Collection("waifus")
	fmt.Println("Connected to MongoDB!")
}
