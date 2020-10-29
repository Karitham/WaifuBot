package database

import (
	"context"
	"strings"

	"github.com/Karitham/WaifuBot/query"
	"github.com/diamondburned/arikawa/discord"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

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

// Set adds a waifu to the user each time he has a new one
func (fav Favorite) Set(uID discord.UserID) error {
	opts := options.FindOneAndUpdate().SetUpsert(true)
	_, err := collection.FindOneAndUpdate(
		context.TODO(),
		bson.M{
			"_id": uID,
		},
		bson.M{
			"$set": bson.M{"Favourite": CharLayout{
				ID:    fav.Character.ID,
				Image: fav.Character.Image.Large,
				Name:  strings.Join(strings.Fields(fav.Character.Name.Full), " "),
			}},
		},
		opts,
	).DecodeBytes()
	if err != mongo.ErrNoDocuments && err != nil {
		return err
	}
	return nil
}

// Set set the user quote on his profile
func (quote Quote) Set(uID discord.UserID) error {
	opts := options.FindOneAndUpdate().SetUpsert(true)
	_, err := collection.FindOneAndUpdate(
		context.TODO(),
		bson.M{
			"_id": uID,
		},
		bson.M{
			"$set": bson.M{"Quote": quote},
		},
		opts,
	).DecodeBytes()
	if err != mongo.ErrNoDocuments && err != nil {
		return err
	}
	return nil
}
