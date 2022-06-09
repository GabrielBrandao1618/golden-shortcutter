package routes

import (
	"github.com/gorilla/mux"
)

func GetRouter() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/createLink", CreateLink).Methods("POST")
	router.HandleFunc("/getUrl/{name}", GetUrl).Methods("GET")
	router.HandleFunc("/", HandleRootRoute)

	return router
}
