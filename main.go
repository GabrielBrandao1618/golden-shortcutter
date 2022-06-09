package main

import (
	"net/http"
	"os"

	"github.com/gorilla/handlers"

	"github.com/GabrielBrandao1618/golden-shortcutter/database"
	"github.com/GabrielBrandao1618/golden-shortcutter/routes"
)

func main() {
	database.Migrate()

	allowedOrigins := handlers.AllowedOrigins([]string{"*"})
	allowedHeaders := handlers.AllowedHeaders([]string{"Content-Type", "Access-Control-Allow-Origin"})
	allowedMethods := handlers.AllowedMethods([]string{"POST", "GET"})

	port := os.Getenv("PORT")

	http.ListenAndServe(":"+port, handlers.CORS(allowedHeaders, allowedOrigins, allowedMethods)(routes.GetRouter()))
}
