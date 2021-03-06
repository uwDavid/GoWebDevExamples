package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	// OR connect using a URL
	// check sql documentation on this connection string
	connStr := "postgres://tester:password@localhost/test?sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Now connected to db.")
	}
	fmt.Println("end of program.")
}
