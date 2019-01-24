package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Book Struct (Model)
type Book struct {
	ID        string  `json:"id"`
	Isbn      string  `json:"isbn"`
	Title     string  `json:"title"`
	Author    *Author `json:"author"`
	Publisher string  `json:"publisher"`
	Audiobook bool    `json:"audiobook"`
}

// Author Struct (Model)
type Author struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

// Init books var as a slice Book struct
var books []Book

// Every handler has to have a Response Writer and Request param

// Gets all books
func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

// Get a single book
func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	params := mux.Vars(r) // Get params

	for _, item := range books {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Book{})
}

// Update book
func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			var book Book
			_ = json.NewDecoder(r.Body).Decode(&book)
			book.ID = params["id"]
			books = append(books, book)
			json.NewEncoder(w).Encode(book)
			return
		}
	}
	json.NewEncoder(w).Encode(books)
}

// Create book
func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID = strconv.Itoa(rand.Intn(10000000)) // Not safe DON'T USE IN PRODUCTION!!!
	books = append(books, book)
	json.NewEncoder(w).Encode(book)
}

// Remove book
func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
		}
		break
	}
	json.NewEncoder(w).Encode(books)
}

func handleRequest() {
	r := mux.NewRouter()

	// Mock Data
	books = append(books, Book{
		ID:    "1",
		Isbn:  "234255",
		Title: "How to get Money!",
		Author: &Author{
			FirstName: "Reggie",
			LastName:  "Davis",
		},
		Publisher: "Magnified Group",
		Audiobook: false,
	},
	)

	books = append(books, Book{
		ID:    "2",
		Isbn:  "h43842",
		Title: "How to get Sued!",
		Author: &Author{
			FirstName: "Jordan",
			LastName:  "Franklin",
		},
		Publisher: "Magnified Group",
		Audiobook: true,
	},
	)

	r.HandleFunc("/api/v1/books", getBooks).Methods("GET")
	r.HandleFunc("/api/v1/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/v1/books", createBook).Methods("POST")
	r.HandleFunc("/api/v1/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/api/v1/books/{id}", deleteBook).Methods("DELETE")

	fmt.Println("Go to localhost:8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}

func main() {
	handleRequest()
}
