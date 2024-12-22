package bd

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func DeletePost(id string, userId string) error {

	db := MongoCN.Database(DatabaseName)
	col := db.Collection("posts")

	objId, _ := primitive.ObjectIDFromHex(id)

	condition := bson.M{
		"_id":    objId,
		"userid": userId,
	}

	result, err := col.DeleteOne(context.TODO(), condition)
	if result.DeletedCount == 0 {
		//return new error
		//mt.Println("error %w", err)
		return errors.New("post not found")
	}
	return err
}
