package main

import (
	"io"
	"net/http"
)

type hotdog int
type hotcat int

func (d hotdog) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "dog dog")
}

func (d hotcat) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "cat cat")
}

func main() {
	var d hotdog
	var c hotcat

	// Instead of mux.Handle, we change it to http.Handle
	http.Handle("/dog/", d)
	http.Handle("/cat/", c)

	http.ListenAndServe(":8080", nil)
	// Instead of the sepcific handler, we pass 'nil' here to use the default mux
}
