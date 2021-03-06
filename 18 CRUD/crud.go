package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type Person struct {
	id               int64
	first_name       string
	last_name        string
	email            sql.NullString
	gender           string
	date_of_birth    string
	country_of_birth string
}

func main() {
	// check sql documentation on this connection string
	connStr := "postgres://tester:password@localhost/test?sslmode=disable"

	// Open database connection
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Optional: ping database to ensure we are connected
	err = db.Ping()
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Connected to db.")
	}

	// Querying database
	// READ
	fmt.Printf("Running Query: SELCT * From person LIMIT 10.\n")
	rows, err := db.Query("SELECT * FROM person LIMIT 10;")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	people := make([]Person, 0)
	for rows.Next() {
		person := Person{}
		err := rows.Scan(&person.id, &person.first_name, &person.last_name, &person.email,
			&person.gender, &person.date_of_birth, &person.country_of_birth)
		if err != nil {
			panic(err)
		}
		people = append(people, person)
	}

	for _, person := range people {
		fmt.Printf("ID: %d, First Name: %s, Last Name: %s \n", person.id, person.first_name, person.last_name)
	}
	fmt.Printf("End of SELECT query.\n")
}
