package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	http.HandleFunc("/", foo)
	// http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, req *http.Request) {
	var s string

	// if method is POST, we will handle the submission
	// http.MethodPost is a constant = "POST"
	if req.Method == http.MethodPost {
		f, h, err := req.FormFile("q")
		//Here we use FormFile vs. FormValue - it catches the file
		//the identifier is "q"

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer f.Close()

		// Server log info
		fmt.Println("\nfile:", f, "\nheader:", h, "\nerr", err)

		// Read the file
		bs, err := ioutil.ReadAll(f) // this is passing the file, returns a byte slice
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		s = string(bs)

		dst, err := os.Create(filepath.Join("./user/", h.Filename))
		// create a file with Filename
		// dst = destination
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer dst.Close()

		_, err = dst.Write(bs)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
	// We will just write our html directly here
	// enctype = "multipart/form-data" allows us to upload files.
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `
	<form method="POST" enctype="multipart/form-data">
	<input type="file" name="q">
	<input type="submit">
	</form>
	<br>`+s)
}
