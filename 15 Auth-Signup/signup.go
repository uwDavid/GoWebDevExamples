package main

import (
	"html/template"
	"log"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

// To set up in-memory storage for user info
type UserInfo struct {
	UserName string
	Password []byte
	First    string
	Last     string
}

var dbUsers = map[string]UserInfo{} // map of UserName to UserInfo

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}
func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/signup", signup)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	tpl.ExecuteTemplate(w, "index.html", dbUsers)
}

func signup(w http.ResponseWriter, req *http.Request) {
	var userinfo UserInfo

	if req.Method == http.MethodPost {
		un := req.FormValue("username")
		pwd := req.FormValue("password")
		hash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.MinCost)
		if err != nil {
			log.Print("Error generating hash:", err)
		}

		f := req.FormValue("firstname")
		l := req.FormValue("lastname")
		userinfo = UserInfo{un, hash, f, l}
		dbUsers[un] = userinfo
	}
	tpl.ExecuteTemplate(w, "signup.html", userinfo)
}
