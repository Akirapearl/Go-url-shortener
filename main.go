package main

import (
	"crypto/md5"
	"fmt"
	"log"
	"net/http"
	"strings"
)

var urlStore = make(map[string]string)

func shorten(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	url := r.Form.Get("url")
	shortURL := fmt.Sprintf("%x", md5.Sum([]byte(url)))[:5]
	urlStore[shortURL] = url
	fmt.Fprintf(w, "http://localhost:8080/%s\n", shortURL)
}

func redirect(w http.ResponseWriter, r *http.Request) {
	shortURL := strings.TrimPrefix(r.URL.Path, "/")
	//vars := mux.Vars(r)
	//oUrl, ok := urlStore[vars["shortURL"]]
	oUrl, ok := urlStore[shortURL]
	if ok {
		http.Redirect(w, r, oUrl, http.StatusMovedPermanently)
	} else {
		http.NotFound(w, r)
	}
}

func main() {
	r := http.NewServeMux()
	r.HandleFunc("POST /create", shorten)
	r.HandleFunc("GET /{shortURL}", redirect)
	log.Print("Starting server on port :8080...")
	log.Fatal(http.ListenAndServe(":8080", r))

}
