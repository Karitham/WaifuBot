package database

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// DropUser a user via USER ID
func (input InputChar) DropUser() (deletedResult *mongo.DeleteResult) {
	deletedResult, err := Collection.DeleteOne(context.TODO(), bson.M{"UserID": input.UserID})
	if err != nil {
		fmt.Println(err)
	}
	return
}
