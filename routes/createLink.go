package routes

import (
	"encoding/json"
	"net/http"

	"github.com/GabrielBrandao1618/golden-shortcutter/database"
	"github.com/GabrielBrandao1618/golden-shortcutter/model/shortedLink"
)

type createLinkRequestBody struct {
	Ref  string `json:"ref"`
	Name string `json:"name"`
}

func CreateLink(w http.ResponseWriter, r *http.Request) {
	var body shortedLink.ShortedLink

	json.NewDecoder(r.Body).Decode(&body)
	url := shortedLink.ShortedLink{Ref: body.Ref, Name: body.Name}

	result := database.CreateUrl(url.Ref, body.Name)
	finalJson, _ := json.Marshal(result)

	w.Write(finalJson)
}
