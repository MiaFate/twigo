package bd

import (
	"context"

	"github.com/miafate/twigo/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetProfile(id string) (models.Usuario, error) {

	db := MongoCN.Database(DatabaseName)
	col := db.Collection("users")

	var perfil models.Usuario
	objId, _ := primitive.ObjectIDFromHex(id)

	condition := bson.M{
		"_id": objId,
	}

	err := col.FindOne(context.TODO(), condition).Decode(&perfil)
	perfil.Password = ""
	if err != nil {
		return perfil, err
	}

	return perfil, nil
}
