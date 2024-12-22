package bd

import (
	"context"

	"github.com/miafate/twigo/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetUsers() []models.Usuario {
	//ctx := context.TODO()

	db := MongoCN.Database(DatabaseName)
	col := db.Collection("users")

	var resultado []models.Usuario

	cursor, err := col.Find(context.TODO(), bson.D{}, options.Find().SetSort(bson.D{{Key: "name", Value: 1}}))
	if err != nil {
		panic(err)
	}
	if err = cursor.All(context.TODO(), &resultado); err != nil {
		panic(err)
	}

	return resultado

}
