package routes

import (
	"encoding/json"
	"net/http"

	"github.com/GabrielBrandao13/golden-shortcutter/database"
	"github.com/GabrielBrandao13/golden-shortcutter/model/shortedLink"
)

type generateUrlRequestBody struct {
	Ref string `json:"ref"`
}

func GenerateUrl(w http.ResponseWriter, r *http.Request) {
	var body generateUrlRequestBody

	json.NewDecoder(r.Body).Decode(&body)
	url := shortedLink.New(body.Ref)

	database.CreateUrl(url)
	finalJson, _ := json.Marshal(url)

	w.Write(finalJson)
}

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
