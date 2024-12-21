package main

import (
	"context"
	"log"

	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"

	"github.com/miafate/twigo/bd"
	"github.com/miafate/twigo/configuration"
	"github.com/miafate/twigo/models"
)

func main() {
	err := bd.Connect(configuration.Ctx)
	if err != nil {
		log.Fatal(err)
		return
	}
	godotenv.Load(.env)
	configuration.Ctx = context.WithValue(configuration.Ctx, models.Key("method"), request.HTTPMethod)
	configuration.Ctx = context.WithValue(configuration.Ctx, models.Key("user"), SecretModel.User)

}
