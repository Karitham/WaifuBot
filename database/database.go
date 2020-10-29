package database

import (
	"context"
	"log"
	"time"

	"github.com/Karitham/WaifuBot/config"
	"github.com/Karitham/WaifuBot/query"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Initialise is used to start the database
func Initialise(config *config.ConfStruct) {
	// Connect to MongoDB
	clientOptions := options.Client().ApplyURI(config.MongoURL)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	collection = client.Database("waifubot").Collection("waifus")
	log.Println("Connected to WaifuDB!")
}

var collection *mongo.Collection

// CharID represent a characterID
type CharID uint

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

// Favorite represent the favorite char
type Favorite query.CharSearchStruct

// Quote represent the data needed to change user quote
type Quote string

// CharStruct alias query.CharStruct
type CharStruct query.CharStruct
