package bd

import (
	"context"

	"github.com/miafate/twigo/models"
	"go.mongodb.org/mongo-driver/bson"
)

func GetFriendsPosts(id string, page int64) ([]models.GetFriendsPosts, bool) {
	db := MongoCN.Database(DatabaseName)
	col := db.Collection("relationship")

	skip := (page - 1) * 20

	conditions := make([]bson.M, 0)
	conditions = append(conditions, bson.M{"$match": bson.M{"userid": id}})
	conditions = append(conditions, bson.M{
		"$lookup": bson.M{
			"from":         "posts",
			"localField":   "friendid",
			"foreignField": "userid",
			"as":           "post",
		},
	})
	conditions = append(conditions, bson.M{"$unwind": "$post"})
	conditions = append(conditions, bson.M{"$sort": bson.M{"post.date": -1}})
	conditions = append(conditions, bson.M{"$skip": skip})
	conditions = append(conditions, bson.M{"$limit": 20})

	var result []models.GetFriendsPosts

	cursor, err := col.Aggregate(context.TODO(), conditions)
	if err != nil {
		return result, false
	}

	err = cursor.All(context.TODO(), &result)
	if err != nil {
		return result, false
	}

	return result, true

}
