package routes

import (
	"github.com/gorilla/mux"
)

func GetRouter() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/generateUrl", GenerateUrl)
	router.HandleFunc("/getUrl", GetUrl)

	return router
}
