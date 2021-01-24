package main

import (
	"html/template"
	"log"
	"net/http"
)

type hotdog int

func (m hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	err := req.ParseForm() // parseform
	if err != nil {
		log.Fatalln(err)
	}
	tpl.ExecuteTemplate(w, "index.gohtml", req.Form)
	// After parsing form, pass data to ExecuteTemplate
	// request.Form passes data from the URL and the body
	// request.PostForm only passes data from the body
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

func main() {
	var d hotdog
	http.ListenAndServe(":8080", d)
}
