package bd

import (
	"context"

	"github.com/miafate/twigo/models"
	"go.mongodb.org/mongo-driver/bson"
)

func GetProducts() ([]*models.GetProducts, bool) {

	db := MongoCN.Database(DatabaseName)
	col := db.Collection("products")

	var result []*models.GetProducts

	cursor, err := col.Find(context.TODO(), bson.M{})
	if err != nil {
		return result, false
	}

	for cursor.Next(context.TODO()) {
		var products models.GetProducts
		err := cursor.Decode(&products)
		if err != nil {
			return result, false
		}
		result = append(result, &products)
	}

	return result, true
}
