package bd

import (
	"context"

	"github.com/miafate/twigo/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetPosts(id string, pag int64) ([]*models.GetPosts, bool) {

	db := MongoCN.Database(DatabaseName)
	col := db.Collection("posts")

	var result []*models.GetPosts

	condition := bson.M{
		"userid": id,
	}

	opts := options.Find()
	opts.SetLimit(10)
	opts.SetSort(bson.D{{Key: "date", Value: -1}})
	opts.SetSkip((pag - 1) * 20)

	cursor, err := col.Find(context.TODO(), condition, opts)
	if err != nil {
		return result, false
	}

	for cursor.Next(context.TODO()) {
		var posts models.GetPosts
		err := cursor.Decode(&posts)
		if err != nil {
			return result, false
		}
		result = append(result, &posts)
	}

	return result, true
}
