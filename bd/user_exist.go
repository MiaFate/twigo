package bd

import (
	"context"

	"github.com/miafate/twigo/models"
	"go.mongodb.org/mongo-driver/bson"
)

func UserExist(email string) (models.Usuario, bool, string) {
	ctx := context.TODO()

	db := MongoCN.Database(DatabaseName)
	col := db.Collection("users")

	condition := bson.M{"email": email}
	var resultado models.Usuario

	err := col.FindOne(ctx, condition).Decode(&resultado)
	Id := resultado.Id.Hex()
	if err != nil {
		return resultado, false, Id
	}
	return resultado, true, Id
}
