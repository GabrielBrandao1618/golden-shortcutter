package routes

import (
	"github.com/gorilla/mux"
)

func GetRouter() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/generateUrl", GenerateUrl).Methods("POST")
	router.HandleFunc("/getUrl/{name}", GetUrl).Methods("GET")

	return router
}
