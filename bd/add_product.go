package bd

import (
	"github.com/miafate/twigo/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/net/context"
)

func AddProduct(product models.Product) (string, bool, error) {

	db := MongoCN.Database(DatabaseName)
	col := db.Collection("products")

	register := bson.M{
		"message": product.Product,
		"date":    product.Date,
	}
	result, err := col.InsertOne(context.TODO(), register)
	if err != nil {
		return "", false, err
	}

	objId := result.InsertedID.(primitive.ObjectID)

	return objId.Hex(), true, nil

}
