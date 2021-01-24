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

	mux := http.NewServeMux()
	mux.Handle("/dog/", d)
	// path:/dog/something/else => still runs dog code
	mux.Handle("/cat/", c)
	// if we do "/cat" => anything down the path won't be handled

	http.ListenAndServe(":8080", mux)
	// We pass Mux as our handler.
	// This is if we need a custom mux
}
