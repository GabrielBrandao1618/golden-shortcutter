package main

import (
	"net/http"

	"github.com/gorilla/handlers"

	"github.com/GabrielBrandao1618/golden-shortcutter/database"
	"github.com/GabrielBrandao1618/golden-shortcutter/routes"
)

func main() {
	database.Migrate()

	allowedOrigins := handlers.AllowedOrigins([]string{"*"})
	allowedHeaders := handlers.AllowedHeaders([]string{"Content-Type", "Access-Control-Allow-Origin"})
	allowedMethods := handlers.AllowedMethods([]string{"POST", "GET"})

	http.ListenAndServe(":8080", handlers.CORS(allowedHeaders, allowedOrigins, allowedMethods)(routes.GetRouter()))
}
