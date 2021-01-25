package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", set)
	http.HandleFunc("/read", read)
	http.HandleFunc("/abundance", abundance)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func set(w http.ResponseWriter, req *http.Request) {
	c := http.Cookie{
		Name:  "my-cookie",
		Value: "Cookie from /",
		Path:  "/",
	}
	http.SetCookie(w, &c) // this is a pointer back to cookie
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, "COOKIE WRITTEN - cookie value: "+c.Value)
	io.WriteString(w, "<br>in chrome go to: dev tools / application / cookies")
}

func read(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	c1, err := req.Cookie("my-cookie")
	if err != nil {
		log.Println(err)
	} else {
		io.WriteString(w, "YOUR COOKIE #1:"+c1.Value+"<br>")
	}

	c2, err := req.Cookie("general")
	if err != nil {
		log.Println(err)
	} else {
		io.WriteString(w, "YOUR COOKIE #2:"+c2.Value+"<br>")
	}

	c3, err := req.Cookie("specific")
	if err != nil {
		log.Println(err)
	} else {
		io.WriteString(w, "YOUR COOKIE #3:"+c3.Value+"<br>")
	}
}

func abundance(w http.ResponseWriter, req *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:  "general",
		Value: "1st cookie from /abundance",
	})

	http.SetCookie(w, &http.Cookie{
		Name:  "specific",
		Value: "2nd cookie from /abundance",
	})
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, "Additional cookies written - CHECK YOUR BROWSER")
}
