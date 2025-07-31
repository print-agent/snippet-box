package main

import (
	"log"
	"net/http"
)

func handlerHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from SnippetBox"))
}

func handlerSnippetView(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a specific snippet..."))
}

func handlerSnippetCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a form for creating a new snippet..."))
}

func main() {
	const port = "4000"

	mux := http.NewServeMux()

	mux.HandleFunc("/", handlerHome)
	mux.HandleFunc("/snippet/view", handlerSnippetView)
	mux.HandleFunc("/snippet/create", handlerSnippetCreate)

	log.Println("Starting on server :%s", port)

	err := http.ListenAndServe(":"+port, mux)
	log.Fatal(err)
}
