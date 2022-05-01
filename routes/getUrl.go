package routes

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/GabrielBrandao1618/golden-shortcutter/database"
	"github.com/GabrielBrandao1618/golden-shortcutter/model/shortedLink"
)

func GetUrl(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	ref := database.GetUrlByCustomName(vars["name"])

	finalJson, _ := json.Marshal(shortedLink.ShortedLink{Name: vars["name"], Ref: ref})

	w.Write(finalJson)
}
