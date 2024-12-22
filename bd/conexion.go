package bd

import (
	"context"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoCN *mongo.Client
var DatabaseName string

func Connect() error {
	user := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	connStr := fmt.Sprintf("mongodb+srv://%s:%s@%s/?retryWrites=true&w=majority", user, password, host)

	var clientOptions = options.Client().ApplyURI(connStr)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		fmt.Println("Error al conectar a la base de datos: " + err.Error())
		return err
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		fmt.Println("Error al conectar a la base de datos: " + err.Error())
		return err
	}

	fmt.Println("Conexión exitosa a la base de datos")
	MongoCN = client
	DatabaseName = os.Getenv("DB_NAME")
	return nil

}

func CheckConnection() bool {
	err := MongoCN.Ping(context.TODO(), nil)
	return err == nil
}
