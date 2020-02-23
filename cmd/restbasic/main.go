package main

import(
    // "fmt"
    "net/http"
    "github.com/gorilla/mux"
    "encoding/json"
)

type Post struct{
    title string `json: "title"` 
}

var posts []Post

func handleGetAll(w http.ResponseWriter, r *http.Request){
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(posts)
}

func main(){
    r := mux.NewRouter()
    r.HandleFunc("/", handleGetAll).Methods("GET")
    http.ListenAndServe(":8080", r)
}