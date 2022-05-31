package main

import (
	"log"
	"net/http"

	"github.com/gorilla/handlers"
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

	allowedOrigins := handlers.AllowedOrigins([]string{"*"})
	allowedHeaders := handlers.AllowedHeaders([]string{"Content-Type", "Access-Control-Allow-Origin"})
	allowedMethods := handlers.AllowedMethods([]string{"POST", "GET"})

	http.ListenAndServe(":8080", handlers.CORS(allowedHeaders, allowedOrigins, allowedMethods)(routes.GetRouter()))
}
