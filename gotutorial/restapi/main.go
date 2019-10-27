package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//Book Struct (Model)
type Book struct {
	ID     string  `json:"id"`
	Isbn   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

//Init books var as a slice Book struct
var books []Book

//Author Struct
type Author struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

// Get Single Book
func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) //Get params
	//Loop through books and find with id
	for _, item := range books {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Book{})
}

func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

//Create a New Book
func createBook(w http.ResponseWriter, r *http.Request) {

}

func updateBook(w http.ResponseWriter, r *http.Request) {

}

func deleteBooks(w http.ResponseWriter, r *http.Request) {

}

func main() {
	//Init Router
	r := mux.NewRouter()

	//Mock Data - @todo - implement DB
	books = append(books, Book{ID: "1", Isbn: "448743", Title: "Book One", Author: &Author{FirstName: "John", LastName: "Doe"}})
	books = append(books, Book{ID: "2", Isbn: "465743", Title: "Book Two", Author: &Author{FirstName: "Steave", LastName: "Mux"}})

	//Route Handlers/Endpoints
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", deleteBooks).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", r))
}
