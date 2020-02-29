package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

const (
	hostname = "localhost"
	hostPort = 5432
	username = "postgres"
	password = "Circumstances123"
	dbname   = "golanguage"
)

func Login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Login Endpoint Reached")
}

func Register(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Register Endpoint Reached")
}

//DB Connection
func ConnectDB() {

	connString := fmt.Sprintf("port=%d host=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		hostPort, hostname, username, password, dbname)

	db, err := sql.Open("postgres", connString)
	if err != nil {
		panic(err)
	}
	// fmt.Println("You are connected indeed")
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")
}

func main() {
	ConnectDB()

	myRouter := mux.NewRouter()
	r.Use(apiMiddleware)
	myRouter.HandleFunc("/", handleBase).Methods("GET")
	myRouter.HandleFunc("/login", Login).Methods("POST")
	myRouter.HandleFunc("/register", Register).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", myRouter))

}
