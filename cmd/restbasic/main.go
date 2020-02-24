package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"github.com/jackc/pgx/v4"
	"github.com/gorilla/mux"
)

type Article struct {
	Title   string `json:"Title"`
	Desc    string `json:"Desc"`
	Content string `json:"Content"`
}

type Articles []Article

func allArticles(w http.ResponseWriter, r *http.Request) {
	articles := Articles{
		Article{
			Title:   "Go Language",
			Desc:    "Golang Tutorial",
			Content: "Hello World",
		},
	}
	fmt.Println("ENDPOINT HIT: All Articles Endpoint")
	json.NewEncoder(w).Encode(articles)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Homepage Endpoint Hit")
}

func postArticle(w http.ResponseWriter, r *http.Request) {
	fmt.Println("POST endpoint reached")
}

func handleRequest() {
	myRouter := mux.NewRouter().StrictSlash(true)
	// http.HandleFunc("/", homePage)
	myRouter.HandleFunc("/articles", allArticles).Methods("GET")
	myRouter.HandleFunc("/PostArticle", postArticle).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", myRouter))
}

func main() {
	handleRequest()
}
