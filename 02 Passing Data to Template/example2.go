package main

import (
	"log"
	"os"
	"text/template"
)

type Product struct {
	Name  string
	Count int
}

func main() {
	//Step 1: We will set up our dummy data
	milk := Product{Name: "milk", Count: 10}
	bread := Product{Name: "bread", Count: 15}

	inventory := []Product{milk, bread}

	//Step 2: We use ParseGlob to parse all ".tp" extensions
	//We can later select which template ExecuteTemplate() uses
	tpl := template.Must(template.ParseGlob("*tp"))

	//Step 3: We will use os.Create() to make "index.html" file
	file, err := os.Create("index.html")
	if err != nil {
		log.Println("Error creating file,", err)
	}
	defer file.Close()

	//Step 4: Select which template to execute
	err2 := tpl.ExecuteTemplate(file, "two.tp", inventory)
	if err != nil {
		log.Fatalln(err2)
	}
	//Step 5: run this file, and see our "index.html"
}
