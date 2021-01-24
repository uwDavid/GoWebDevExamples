package main

import (
	"io"
	"net/http"
)

// The Handler function just have to implement this signature
func d(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "dog dog")
}

func c(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "cat cat")
}

func main() {
	// Instead of mux.Handle, we change it to http.Handle
	// Finally, we change it to http.HandleFunc
	http.HandleFunc("/dog/", d)
	http.HandleFunc("/cat/", c)

	// This function is the same as:
	// http.Handle("/cat/", http.HandlerFunc(c))
	// Note the explicit type conversion here. Just calling http.Handle("/cat/", c) will not work.

	http.ListenAndServe(":8080", nil)
	// Instead of the sepcific handler, we pass 'nil' here to use the default mux
}
