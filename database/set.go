package database

import (
	"context"
	"strings"

	"github.com/diamondburned/arikawa/v2/discord"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

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
