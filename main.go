package main

import (
	"log"
	"net/http"

	"github.com/joho/godotenv"

	"github.com/GabrielBrandao1618/golden-shortcutter/database"
	"github.com/GabrielBrandao1618/golden-shortcutter/routes"
)

func setupEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}

func main() {
	setupEnv()
	database.Migrate()

	http.ListenAndServe(":8080", routes.GetRouter())
}
