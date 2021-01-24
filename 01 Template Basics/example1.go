package main

import (
	"log"
	"os"
	"text/template"
)

// This is a basic example of using templates.
// We will output our file to our console log.

func main() {
	// tpl := template.Must(template.Files("tpl.gohtml"))
	tpl := template.Must(template.ParseGlob("*gohtml"))

	err := tpl.Execute(os.Stdout, 42) //We pass 42 as data into template
	if err != nil {
		log.Fatalln(err)
	}
}
