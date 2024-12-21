package bd

import (
	"context"
	"fmt"

	"github.com/miafate/twigo/models"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoCN *mongo.Client
var DatabaseName string

func Connect(ctx context.Context) error {
	user := ctx.Value(models.Key("db_user")).(string)
	password := ctx.Value(models.Key("db_password")).(string)
	host := ctx.Value(models.Key("db_host")).(string)
	connStr := fmt.Sprintf("mongodb+srv://%s:%s@%s/?retryWrites=true&w=majority", user, password, host)

	var clientOptions = options.Client().ApplyURI(connStr)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		fmt.Println("Error al conectar a la base de datos: " + err.Error())
		return err
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		fmt.Println("Error al conectar a la base de datos: " + err.Error())
		return err
	}

	fmt.Println("Conexión exitosa a la base de datos")
	MongoCN = client
	DatabaseName = ctx.Value(models.Key("db_name")).(string)
	return nil

}

func CheckConnection() bool {
	err := MongoCN.Ping(context.TODO(), nil)
	return err == nil
}