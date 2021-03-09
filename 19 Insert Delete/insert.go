package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

// In order to export these fields, we have to capitalize them
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

	// Insert - using Exec()
	_, err2 := db.Exec("INSERT INTO person (first_name, last_name, email, gender, date_of_birth, country_of_brith) values($1, $2, $3, $4, $5, $6)", "Insert", "Example", "inserted@example.com", "Male", "1963/07/13", "Unicorn Land")
	if err2 != nil {
		panic(err)
	}

	// To ensure we added this new person
	// Read - using QueryRow()
	country := "Unicorn Land"
	row := db.QueryRow("SELECT * FROM person WHERE country_of_birth= $1", country)
	onePerson := Person{}
	err3 := row.Scan(&onePerson.id, &onePerson.first_name, &onePerson.last_name, &onePerson.email,
		&onePerson.gender, &onePerson.date_of_birth, &onePerson.country_of_birth)
	if err3 == sql.ErrNoRows {
		fmt.Printf("Person from country %s does not exist.", country)
		panic(err)
	}
	if err3 != nil {
		panic(err)
	}
	fmt.Printf("Insert Example ID: %d, First Name: %s, Last Name: %s, Country of Birth: %s\n", onePerson.id, onePerson.first_name, onePerson.last_name, onePerson.country_of_birth)

	//
}
