package main

import (
	"log"
	"net/http"
	"text/template"

	"github.com/google/uuid"
)

// To set up in-memory storage for user info
type UserInfo struct {
	UserName string
	First    string
	Last     string
}

var dbUsers = map[string]UserInfo{}  // map of UserName to UserInfo
var dbSessions = map[string]string{} // map of session ID to UserName

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/bar", bar)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {

	// get cookie
	c, err := req.Cookie("session")
	// If no cookie, set cookie
	if err != nil {
		sID := uuid.NewString()
		c = &http.Cookie{
			Name:  "sessionCookie",
			Value: sID,
		}
		http.SetCookie(w, c)
	}

	// if the user exists already, get user
	var userinfo UserInfo
	username, ok := dbSessions[c.Value]
	if ok {
		userinfo = dbUsers[username]
	}

	// process POST submission and save value to database
	if req.Method == http.MethodPost {
		un := req.FormValue("username")
		f := req.FormValue("firstname")
		l := req.FormValue("lastname")
		userinfo = UserInfo{un, f, l}
		dbSessions[c.Value] = un
		dbUsers[un] = userinfo
	}

	tpl.ExecuteTemplate(w, "index.html", userinfo)
}

func bar(w http.ResponseWriter, req *http.Request) {

	// get cookie
	c, err := req.Cookie("sessionCookie")
	// return to index if do not have cookie
	if err != nil {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		log.Println("User did not have a cookie, returned to index")
		return
	}
	// return to index if there's no userinfo/username in database
	un, ok := dbSessions[c.Value]
	if !ok {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		log.Println("There is no UserInfo, returned to index")
		return
	}
	u := dbUsers[un]
	tpl.ExecuteTemplate(w, "bar.html", u)
}
