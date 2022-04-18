package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func handleRootRoute(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Hello world!"))
}

type url struct {
	ref      string
	hashCode uuid.UUID
}
type urlFromClient struct {
	Ref string `json:"ref"`
}

func generateUrl(w http.ResponseWriter, req *http.Request) {
	var parsedBody urlFromClient

	json.NewDecoder(req.Body).Decode(&parsedBody)

	urlRef := parsedBody.Ref
	hashCode := uuid.New()

	finalUrl := url{ref: urlRef, hashCode: hashCode}
	w.Write([]byte(fmt.Sprintf("HashCode: %s, ref: %s", finalUrl.hashCode, finalUrl.ref)))
}

type hashCodeFromClient struct {
	HashCode string `json:"hashCode"`
}

func getUrl(w http.ResponseWriter, req *http.Request) {
	var code hashCodeFromClient

	json.NewDecoder(req.Body).Decode(&code)

	w.Write([]byte(fmt.Sprintf("Código: %s inserido", code.HashCode)))
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", handleRootRoute)
	r.HandleFunc("/generateUrl", generateUrl)
	r.HandleFunc("/getUrl", getUrl)

	http.ListenAndServe(":8080", r)
}
