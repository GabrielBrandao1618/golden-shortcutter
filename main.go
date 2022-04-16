package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func handleRootRoute(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Hello world!"))
}

type url struct {
	ref      string
	hashCode string
}

func generateUrl(w http.ResponseWriter, req *http.Request) {
	urlRef := "google.com"
	hashCode := "sjhsjh#167"

	finalUrl := url{ref: urlRef, hashCode: hashCode}
	w.Write([]byte(fmt.Sprintf("HashCode: %s, ref: %s", finalUrl.hashCode, finalUrl.ref)))
}

func getUrl(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Coletando url... sqn"))
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", handleRootRoute)
	r.HandleFunc("/generateUrl", generateUrl)
	r.HandleFunc("/getUrl", getUrl)

	http.ListenAndServe(":8080", r)
}
