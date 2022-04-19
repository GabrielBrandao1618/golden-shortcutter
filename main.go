package main

import (
	"log"
	"net/http"

	"github.com/joho/godotenv"

	"github.com/GabrielBrandao13/golden-shortcutter/database"
	"github.com/GabrielBrandao13/golden-shortcutter/routes"
	"github.com/gorilla/mux"
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

	r := mux.NewRouter()
	r.HandleFunc("/generateUrl", routes.GenerateUrl)
	r.HandleFunc("/getUrl", routes.GetUrl)

	http.ListenAndServe(":8080", r)
}
