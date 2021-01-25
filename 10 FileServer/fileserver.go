package main

import (
	"io"
	"net/http"
)

/*
// This is the previous example but using FileServer instead
//FileServer takes a directory, "." means current directory
//FileServer is inadequate because it will serve up all the content in the folder, including the source code
func main() {
	//Anything at root, the handler will be FileServer
	http.Handle("/", http.FileServer(http.Dir(".")))
	http.HandleFunc("/dog", dog)
	http.ListenAndServe(":8080", nil)
}

func dog(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `<img src="Marty.png">`)
	// this will look for a /Marty route
	// it will find match in the "/" route and will serve Marty.png
}
*/

// We now modify the above code and use StripPrefix along with FileServer
func main() {
	http.HandleFunc("/", dog)
	http.Handle("/resources/", http.StripPrefix("/resources", http.FileServer(http.Dir("./assets"))))
	// http.Handle takes a route, and a handler
	// "resources/" => anything down this path is caught using trailing /
	// stripPrefix takes a prefix to strip off, and a handler
	// stripPrefix strips off the "/resources" and we are left with /Marty.png
	// here, fileServer will serve everything in "./assets" folder and looks for Marty.png
	http.ListenAndServe(":8080", nil)
}

func dog(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `<img src="/resources/Marty.png">`)
	// This will now direct to /resources/Marty.png
	// And this is caught by the "/resources", and will be handled by the FileServer
}
