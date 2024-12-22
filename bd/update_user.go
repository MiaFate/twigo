package bd

import (
	"context"

	"github.com/miafate/twigo/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func UpdateUser(u models.Usuario, id string) (bool, error) {

	db := MongoCN.Database(DatabaseName)
	col := db.Collection("users")

	objID, _ := primitive.ObjectIDFromHex(id)
	// fmt.Println("objID: ", objID)
	// condition := bson.M{"_id": id}

	// var resultado models.Usuario
	// notFound := col.FindOne(context.TODO(), condition).Decode(&resultado)
	// fmt.Println("resultado: ", resultado)
	// if notFound != nil {
	// 	fmt.Println("Registro no encontrado")
	// 	return false, notFound
	// }

	registro := make(map[string]interface{})
	if len(u.Nombre) > 0 {
		registro["nombre"] = u.Nombre
	}
	if len(u.Apellidos) > 0 {
		registro["apellidos"] = u.Apellidos
	}
	registro["fecha_nacimiento"] = u.FechaNacimiento
	if len(u.Avatar) > 0 {
		registro["avatar"] = u.Avatar
	}
	if len(u.Banner) > 0 {
		registro["banner"] = u.Banner
	}
	if len(u.Biografia) > 0 {
		registro["biografia"] = u.Biografia
	}
	if len(u.Ubicacion) > 0 {
		registro["ubicacion"] = u.Ubicacion
	}
	if len(u.SitioWeb) > 0 {
		registro["sitio_web"] = u.SitioWeb
	}

	updtString := bson.M{
		"$set": registro,
	}
	filter := bson.M{"_id": bson.M{"$eq": objID}}

	var resultado models.Usuario
	err := col.FindOneAndUpdate(context.TODO(), filter, updtString).Decode(&resultado)
	//_, err := col.UpdateOne(context.TODO(), filter, updtString)
	if err != nil {
		return false, err
	}

	return true, nil
}
