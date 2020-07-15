package database

import (
	"context"
	"fmt"

	"github.com/andersfylling/disgord"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// OwnsCharacter verify if a use owns a character, if not, returns false
func OwnsCharacter(UserID disgord.Snowflake, CharacterID int) bool {
	_, err := Collection.FindOne(context.TODO(),
		bson.M{
			"$eq": bson.M{
				"_id": UserID,
				"Waifus": bson.M{
					"ID": CharacterID,
				},
			},
		},
	).DecodeBytes()
	switch {
	case err == mongo.ErrNoDocuments:
		return false
	case err != nil:
		fmt.Println(err)
	default:
		return true
	}
	return false
}
