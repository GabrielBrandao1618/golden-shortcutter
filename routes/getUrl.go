package routes

import (
	"encoding/json"
	"net/http"

	"github.com/GabrielBrandao1618/golden-shortcutter/database"
	"github.com/GabrielBrandao1618/golden-shortcutter/model/shortedLink"
)

type getUrlRequestBody struct {
	HashCode string `json:"hashCode"`
}

func GetUrl(w http.ResponseWriter, r *http.Request) {
	var body getUrlRequestBody

	json.NewDecoder(r.Body).Decode(&body)

	ref := database.GetUrlByHashCode(body.HashCode)

	finalJson, _ := json.Marshal(shortedLink.ShortedLink{HashCode: body.HashCode, Ref: ref})

	w.Write(finalJson)
}