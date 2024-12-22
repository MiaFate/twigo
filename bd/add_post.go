package bd

import (
	"github.com/miafate/twigo/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/net/context"
)

func AddPost(post models.AddPost) (string, bool, error) {

	db := MongoCN.Database(DatabaseName)
	col := db.Collection("posts")

	register := bson.M{
		"userid":  post.UserId,
		"message": post.Message,
		"date":    post.Date,
	}
	result, err := col.InsertOne(context.TODO(), register)
	if err != nil {
		return "", false, err
	}

	objId := result.InsertedID.(primitive.ObjectID)

	return objId.Hex(), true, nil

}
