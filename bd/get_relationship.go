package bd

import (
	"github.com/miafate/twigo/models"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/net/context"
)

func GetRelationship(t models.Relationship) bool {
	db := MongoCN.Database(DatabaseName)
	col := db.Collection("relationship")

	condition := bson.M{
		"userid":   t.UserId,
		"friendid": t.FriendId,
	}
	var result models.Relationship
	err := col.FindOne(context.TODO(), condition).Decode(&result)
	return err == nil
}
