package database

import (
	"context"
	"fmt"

	"github.com/andersfylling/disgord"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// OwnsCharacter verify if a user owns a character, if not, returns false
func OwnsCharacter(UserID disgord.Snowflake, CharacterID int) bool {
	// Test if a user owns a character
	_, err := Collection.FindOne(
		context.TODO(),
		bson.D{
			{"_id", UserID},
			{"Waifus.ID", CharacterID},
		}).DecodeBytes()
	// Return false if the character was not found
	switch {
	case err == mongo.ErrNoDocuments:
		return false
	case err != nil:
		fmt.Println(err)
		return false
	case err == nil:
		return true
	}
	return false
}
