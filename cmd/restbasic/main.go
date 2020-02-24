package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	// "github.com/gorilla/mux"
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

func handleRequest() {
	// r := mux.NewRouter()
	// http.HandleFunc("/", homePage)
	http.HandleFunc("/", allArticles)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	handleRequest()
}
