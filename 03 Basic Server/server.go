package main

import (
	"fmt"
	"net/http"
)

// Step 1: Initiate our hotdog type
type hotdog int

// Step 2: Implement ServeHTTP for hotdog type
// Now hotdog type implements the Handler interface
func (m hotdog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Insert code here")
}

// Step 3: Now we can pass hotdog 'hd' as the Handler for our server
func main() {
	var hd hotdog
	http.ListenAndServe(":8080", hd)
}
