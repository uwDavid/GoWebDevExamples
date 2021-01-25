package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/bar", bar)
	http.HandleFunc("/barred", barred)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, req *http.Request) {
	//3. Lastly we will visit the index page from "/barred"
	io.WriteString(w, "Request redirected from /bar:"+req.Method)
	fmt.Print("Your request method at foo: ", req.Method, "\n\n")
}

func bar(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Your request method at bar:", req.Method)
	//2. The POST submission will here
	//But we will redirect the client to "/"
	http.Redirect(w, req, "/", http.StatusSeeOther)
}

func barred(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Your request method at barred:", req.Method)
	io.WriteString(w, `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Title</title>
</head>
<body>

<form method="POST" action="/bar">
    <input type="text" name="fname">
    <input type="submit">
</form>

</body>
</html>`)
	//1. We will first go to "/barred" and submit a form
	//This will submit a POST request to "/bar"
}
