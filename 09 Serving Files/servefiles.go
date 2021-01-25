package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", simple)
	http.HandleFunc("/dog", dog)
	http.HandleFunc("/Marty", dogPic)
	http.ListenAndServe(":8080", nil)
}

func simple(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `<body>A Simple Hello</body>`)
}
func dog(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `<img src="Marty">`)
	//This will write to html <img src .. > and serve Marty.png that is in the same folder
	//But it will go to "/Marty" route to get the imagine
}

func dogPic(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "Marty.png") // ServeFile is alot simpler than ServeContent
}

//Note: Marty.png is in the same folder
