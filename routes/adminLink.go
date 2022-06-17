package routes

import (
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/GabrielBrandao1618/golden-shortcutter/database"

)

type response struct {
	VisitsCount uint64 `json:"visitsCount"`
}

func AdminLink(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	name := vars["name"]

	count := database.GetVisitsCount(name)

	jsonResponse, _ := json.Marshal(response{count})
	w.Write(jsonResponse)
}