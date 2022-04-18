package main

import (
	"net/http"

	"github.com/GabrielBrandao13/golden-shortcutter/routes"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/generateUrl", routes.GenerateUrl)
	r.HandleFunc("/getUrl", routes.GetUrl)

	http.ListenAndServe(":8080", r)
}
