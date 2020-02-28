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

//Init Books variable as a slice book struct
var books []Book

//Get All Books
func getBooks(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Get All Book reached")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

//Get Single Book
func getBook(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint get single book reached")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r) // Get Params

	//Loop thru books to get to correct id
	for _, item := range books {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Book{})
}

//Create a New Book
func createBook(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint create a book reached")
	w.Header().Set("Content-Type", "application/json")

	var book Book

	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID = strconv.Itoa(rand.Intn(1000000)) // Mock ID
	books = append(books, book)
	json.NewEncoder(w).Encode(book)
}

//Update a book
func updateBook(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint update a book reached")

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			var book Book
			_ = json.NewDecoder(r.Body).Decode(&book)
			book.ID = params["id"] // Mock ID
			books = append(books, book)
			json.NewEncoder(w).Encode(book)
			return
		}
	}
	json.NewEncoder(w).Encode(books)
}

//Delete a book
func deleteBook(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint delete a book reached")

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(books)
}
func main() {
	//INIT Router
	myRouter := mux.NewRouter()

	//Mock Data - @todo implement database
	books = append(books,
		Book{ID: "1", ISBN: "448743", Title: "Book One", Author: &Author{
			FirstName: "John", LastName: "Doe"},
		})
	books = append(books,
		Book{ID: "2", ISBN: "448744", Title: "Book Two", Author: &Author{
			FirstName: "Jane", LastName: "Doe"},
		})

	//Route Handlers / Endpoints
	myRouter.HandleFunc("/api/books", getBooks).Methods("GET")
	myRouter.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	myRouter.HandleFunc("/api/books", createBook).Methods("POST")
	myRouter.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	myRouter.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", myRouter))
}
