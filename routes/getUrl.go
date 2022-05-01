package routes

import (
	"encoding/json"
	"net/http"

	"github.com/GabrielBrandao1618/golden-shortcutter/database"
	"github.com/GabrielBrandao1618/golden-shortcutter/model/shortedLink"
)

type getUrlRequestBody struct {
	Name string `json:"name"`
}

func GetUrl(w http.ResponseWriter, r *http.Request) {
	var body getUrlRequestBody

	json.NewDecoder(r.Body).Decode(&body)

	ref := database.GetUrlByCustomName(body.Name)

	finalJson, _ := json.Marshal(shortedLink.ShortedLink{Name: body.Name, Ref: ref})

	w.Write(finalJson)
}
