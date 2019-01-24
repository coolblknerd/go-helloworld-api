package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Book struct {
	ID        int16  `json:"id"`
	Title     string `json:"title"`
	Author    string `json:"author"`
	Publisher string `json:"author"`
	Audiobook bool   `json:"audiobook"`
}

func handleRequest() {
	r := mux.NewRouter()
	r.HandleFunc("/api/v1/books", getBooks).Methods("GET")
	r.HandleFunc("/api/v1/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/v1/books", createBook).Methods("POST")
	r.HandleFunc("/api/v1/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/api/v1/books/{id}", deleteBook).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", r))
}

func main() {
	handleRequest()
}
