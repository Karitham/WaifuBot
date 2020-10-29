package database

import (
	"context"
	"strings"
	"time"

	"github.com/Karitham/WaifuBot/query"
	"github.com/diamondburned/arikawa/discord"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// CharStruct alias query.CharStruct
type CharStruct query.CharStruct

// AddRolled adds a waifu to the user each time he has a new one
func (c CharStruct) AddRolled(uID discord.UserID, date time.Time) error {
	opts := options.FindOneAndUpdate().SetUpsert(true)
	err := collection.FindOneAndUpdate(
		context.TODO(),
		bson.M{
			"_id": uID,
		},
		bson.M{
			"$set": bson.M{"Date": time.Now()},
			"$push": bson.M{"Waifus": CharLayout{
				ID:    c.Page.Characters[0].ID,
				Image: c.Page.Characters[0].Image.Large,
				Name:  strings.Join(strings.Fields(c.Page.Characters[0].Name.Full), " "),
			}},
		},
		opts,
	).Decode(&bson.M{})
	if err != mongo.ErrNoDocuments && err != nil {
		return err
	}
	return nil
}

// AddClaimed adds a waifu to the user each time he has a new one
func (c CharStruct) AddClaimed(uID discord.UserID) error {
	opts := options.FindOneAndUpdate().SetUpsert(true)
	err := collection.FindOneAndUpdate(
		context.TODO(),
		bson.M{
			"_id": uID,
		},
		bson.M{
			"$inc": bson.M{"ClaimedWaifus": 1},
			"$push": bson.M{"Waifus": CharLayout{
				ID:    c.Page.Characters[0].ID,
				Image: c.Page.Characters[0].Image.Large,
				Name:  strings.Join(strings.Fields(c.Page.Characters[0].Name.Full), " "),
			}},
		},
		opts,
	).Decode(&bson.M{})
	if err != mongo.ErrNoDocuments && err != nil {
		return err
	}
	return nil
}

// Add adds a waifu to the user each time he has a new one
func (c CharLayout) Add(uID discord.UserID) error {
	opts := options.FindOneAndUpdate().SetUpsert(true)
	err := collection.FindOneAndUpdate(
		context.TODO(),
		bson.M{
			"_id": uID,
		},
		bson.M{
			"$inc": bson.M{"ClaimedWaifus": 1},
			"$push": bson.M{"Waifus": CharLayout{
				ID:    c.ID,
				Image: c.Image,
				Name:  strings.Join(strings.Fields(c.Name), " "),
			}},
		},
		opts,
	).Decode(&bson.M{})
	if err != mongo.ErrNoDocuments && err != nil {
		return err
	}
	return nil
}
