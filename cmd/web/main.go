package main

import (
	"log"
	"net/http"
)

const port = "4000"

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /{$}", home)
	mux.HandleFunc("GET /snippet/view/{id}", snippetView)
	mux.HandleFunc("GET /snippet/create", snippetCreate)
	mux.HandleFunc("POST /snippet/create", snippetCreatePost)

	log.Printf("Starting server on :%s", port)

	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
