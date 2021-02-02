package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/google/uuid"

	"golang.org/x/crypto/bcrypt"
)

// To set up in-memory storage for user info
type UserInfo struct {
	UserName string
	Password []byte
	First    string
	Last     string
	Role     string
}

var dbUsers = map[string]UserInfo{}  // map of UserName to UserInfo
var dbSessions = map[string]string{} // map of session ID to UserName

var tpl *template.Template

// Set up some dummy users in memory
func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
	pw, _ := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.MinCost)
	dbUsers["dummy@email.com"] = UserInfo{"dummy@email.com", pw, "Dummy", "User", "Admin"}
	dbUsers["dummy2@email.com"] = UserInfo{"dummy2@email.com", pw, "Dummy2", "User2", "Student"}
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/signup", signup)
	http.HandleFunc("/login", login)
	http.HandleFunc("/student", student)
	http.HandleFunc("/admin", admin)
	http.HandleFunc("/secret", secret)
	http.HandleFunc("/logout", logout)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	tpl.ExecuteTemplate(w, "index.html", dbUsers)
}

func secret(w http.ResponseWriter, req *http.Request) {
	c, err := req.Cookie("LoginInfo")
	if err != nil {
		http.Error(w, "Cannot find cookie.", http.StatusForbidden)
		http.Redirect(w, req, "/login", http.StatusSeeOther)
		return
	}
	_, ok := dbSessions[c.Value]
	if !ok {
		http.Error(w, "User is not logged in.", http.StatusForbidden)
		http.Redirect(w, req, "/login", http.StatusSeeOther)
		return
	}
	tpl.ExecuteTemplate(w, "secret.html", nil)
}

func student(w http.ResponseWriter, req *http.Request) {
	c, err := req.Cookie("LoginInfo")
	if err != nil {
		http.Error(w, "User is not logged in.", http.StatusForbidden)
		return
	}

	sID, ok := dbSessions[c.Value]
	if !ok {
		http.Error(w, "User is not logged in.", http.StatusForbidden)
		return
	}
	role := dbUsers[sID].Role
	if role != "student" {
		http.Error(w, "User is not a student!", http.StatusForbidden)
		return
	}
	tpl.ExecuteTemplate(w, "student.html", nil)
}

func admin(w http.ResponseWriter, req *http.Request) {
	c, err := req.Cookie("LoginInfo")
	if err != nil {
		http.Error(w, "User is not logged in.", http.StatusForbidden)
		return
	}

	sID, ok := dbSessions[c.Value]
	if !ok {
		http.Error(w, "User is not logged in.", http.StatusForbidden)
		return
	}
	role := dbUsers[sID].Role
	if role != "admin" {
		http.Error(w, "User is not a student!", http.StatusForbidden)
		return
	}
	tpl.ExecuteTemplate(w, "admin.html", nil)
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
		r := "student"
		userinfo = UserInfo{un, hash, f, l, r}
		dbUsers[un] = userinfo
	}
	tpl.ExecuteTemplate(w, "signup.html", userinfo)
}

func alreadyLoggedIn(req *http.Request) bool {
	c, err := req.Cookie("LoginInfo")
	if err != nil {
		return false
	}

	un := dbSessions[c.Value]
	_, ok := dbUsers[un]
	return ok
}

func login(w http.ResponseWriter, req *http.Request) {
	if alreadyLoggedIn(req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}

	if req.Method == http.MethodPost {
		un := req.FormValue("username")
		pwd := req.FormValue("password")
		user, ok := dbUsers[un]
		// Check if user exists
		if !ok {
			http.Error(w, "Username is not in use.", http.StatusForbidden)
		}

		// Check if password matches
		err := bcrypt.CompareHashAndPassword(user.Password, []byte(pwd))
		if err != nil {
			http.Error(w, "Username/password do no match", http.StatusForbidden)
			return
		} else {
			log.Print("Login success!, Redirecting to secrets page.")
		}

		// Create session for user that logged in
		sID := uuid.NewString()
		c := &http.Cookie{
			Name:  "LoginInfo",
			Value: sID,
		}
		http.SetCookie(w, c)
		dbUsers[c.Value] = user

		// Redirect based on user's role
		role := dbUsers[un].Role
		if role == "student" {
			http.Redirect(w, req, "/student", http.StatusSeeOther)
		} else if role == "admin" {
			http.Redirect(w, req, "/admin", http.StatusSeeOther)
		} else {
			http.Redirect(w, req, "/", http.StatusSeeOther)
		}
	}
	tpl.ExecuteTemplate(w, "login.html", nil)
}

func logout(w http.ResponseWriter, req *http.Request) {
	if !alreadyLoggedIn(req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}

	c, _ := req.Cookie("LoginInfo")
	// Delete session from db
	delete(dbSessions, c.Value)
	// Delete cookie
	c = &http.Cookie{
		Name:   "LoginInfo",
		Value:  "",
		MaxAge: -1,
	}
	http.SetCookie(w, c)

	http.Redirect(w, req, "/login", http.StatusSeeOther)
}
