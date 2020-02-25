package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	hostname = "localhost"
	hostPort = 5432
	username = "postgres"
	password = "Circumstances123"
	dbname   = "golanguage"
)

func main() {
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

	insertStatement := ("INSERT INTO users(username) VALUES ($1) returning id_user;")
	id_user := 0

	err = db.QueryRow(insertStatement, "Extillius").Scan(&id_user)
	if err != nil {
		panic(err)
	}
	fmt.Println("New Record is: ", id_user)

	getAllStatement := ("SELECT * FROM users")

	rows, err := db.Query(getAllStatement)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var name string

		err = rows.Scan(&id, &name)
		if err != nil {
			panic(err)
		}
		fmt.Println(id, name)
	}
	err = rows.Err()
	if err != nil {
		panic(err)
	}
}
