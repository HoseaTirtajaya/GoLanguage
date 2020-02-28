package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//Book Struct
type Book struct {
	ID     string  `json:"id"`
	ISBN   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

//Author Struct
type Author struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

//Get All Books
func getBooks(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Get All Book reached")
}

//Get Single Book
func getBook(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint get single book reached")
}

//Create a New Book
func createBook(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint create a book reached")
}

//Update a book
func updateBook(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint update a book reached")
}

//Delete a book
func deleteBook(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint delete a book reached")
}
func main() {
	//INIT Router
	myRouter := mux.NewRouter()

	//Route Handlers / Endpoints
	myRouter.HandleFunc("/api/books", getBooks).Methods("GET")
	myRouter.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	myRouter.HandleFunc("/api/books", createBook).Methods("POST")
	myRouter.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	myRouter.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", myRouter))
}
