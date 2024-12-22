package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"

	"github.com/miafate/twigo/bd"
	"github.com/miafate/twigo/router"
)

func main() {

	godotenv.Load(".env")

	fmt.Println("Site title: " + os.Getenv("SITE_TITLE"))

	// checkeo conexión a la base de datos
	err := bd.Connect()
	if err != nil {
		log.Fatal(err)
		return
	}
	r := router.SetupRouter()

	// Listen and Server in 0.0.0.0:8080
	r.Run()
}
