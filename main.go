package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"

	"github.com/miafate/twigo/bd"
	"github.com/miafate/twigo/configuration"
	"github.com/miafate/twigo/models"
)

func main() {

	godotenv.Load(".env")
	configuration.SetContext(context.Background())
	configuration.Ctx = context.WithValue(configuration.Ctx, models.Key("db_host"), os.Getenv("DB_HOST"))
	configuration.Ctx = context.WithValue(configuration.Ctx, models.Key("db_user"), os.Getenv("DB_USERNAME"))
	configuration.Ctx = context.WithValue(configuration.Ctx, models.Key("db_password"), os.Getenv("DB_PASSWORD"))
	configuration.Ctx = context.WithValue(configuration.Ctx, models.Key("db_name"), os.Getenv("DB_NAME"))
	configuration.Ctx = context.WithValue(configuration.Ctx, models.Key("site_title"), os.Getenv("SITE_TITLE"))

	fmt.Println("Site title: " + os.Getenv("SITE_TITLE"))
	// checkeo conexión a la base de datos
	err := bd.Connect(configuration.Ctx)
	if err != nil {
		log.Fatal(err)
		return
	}

}
