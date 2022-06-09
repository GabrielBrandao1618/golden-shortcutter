package routes

import (
	"encoding/json"
	"net/http"
)

func HandleRootRoute(w http.ResponseWriter, r *http.Request){
	response := map[string] string{
		"createLinkUrl":"/createLink",
		"getUrl":"getUrl/$customName",
	}
	responseJson, _ := json.Marshal(response)
	w.Write(responseJson)
}