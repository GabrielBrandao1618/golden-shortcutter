package routes

import (
	"encoding/json"
	"net/http"

	"github.com/GabrielBrandao1618/golden-shortcutter/database"
	"github.com/GabrielBrandao1618/golden-shortcutter/model/shortedLink"
)

type generateUrlRequestBody struct {
	Ref string `json:"ref"`
}

func GenerateUrl(w http.ResponseWriter, r *http.Request) {
	var body generateUrlRequestBody

	json.NewDecoder(r.Body).Decode(&body)
	url := shortedLink.New(body.Ref)

	result := database.CreateUrl(url.Ref)
	finalJson, _ := json.Marshal(result)

	w.Write(finalJson)
}
