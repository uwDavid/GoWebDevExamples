package main

import (
	"errors"
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
	loggedIn := loggedIn(req)
	if loggedIn {
		tpl.ExecuteTemplate(w, "secret.html", nil)
	} else {
		http.Redirect(w, req, "/login", http.StatusSeeOther)
		return
	}
}

func student(w http.ResponseWriter, req *http.Request) {
	loggedIn := loggedIn(req)
	if !loggedIn {
		log.Println("User not logged in, redirected to login page.")
		http.Redirect(w, req, "/login", http.StatusSeeOther)
		return
	}

	role, err := getRole(req)
	if err != nil {
		log.Println("User role not defined.")
		http.Redirect(w, req, "/login", http.StatusSeeOther)
		return
	}

	if role != "Student" {
		log.Println("User role is not student")
		http.Redirect(w, req, "/login", http.StatusSeeOther)
		return
	}
	log.Println("Entering in Studnet Page.")
	tpl.ExecuteTemplate(w, "student.html", nil)
}

func admin(w http.ResponseWriter, req *http.Request) {
	loggedIn := loggedIn(req)
	if !loggedIn {
		log.Println("User not logged in, redirected to login page.")
		http.Redirect(w, req, "/login", http.StatusSeeOther)
		return
	}

	role, err := getRole(req)
	if err != nil {
		log.Println("User role not defined.")
		http.Redirect(w, req, "/login", http.StatusSeeOther)
		return
	}

	if role != "Admin" {
		log.Println("User role is not admin")
		http.Redirect(w, req, "/login", http.StatusSeeOther)
		return
	}
	log.Println("Entering in Admin Page.")
	tpl.ExecuteTemplate(w, "admin.html", nil)
}

// Check login by checking for active session using cookie
func loggedIn(req *http.Request) bool {
	c, err := req.Cookie("LoginInfo")
	if err != nil {
		log.Println("Did not find cookie.")
		return false
	}

	un := dbSessions[c.Value]
	_, ok := dbUsers[un]
	if !ok {
		log.Println("Did not find user for cookie value")
	}
	return ok
}

// Get client's role permission
func getRole(req *http.Request) (string, error) {
	c, err := req.Cookie("LoginInfo")
	if err != nil {
		log.Println("Did not find cookie.")
		return "", err
	}

	un, ok := dbSessions[c.Value]
	if !ok {
		log.Println("Did not find session using cookie")
		return "", errors.New("Cannot find session")
	}
	user, ok := dbUsers[un]
	if !ok {
		log.Println("Did not find user in user db")
		return "", errors.New("Cannot find user")
	}
	return user.Role, nil
}

// Route Handlers
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

func login(w http.ResponseWriter, req *http.Request) {
	if loggedIn(req) {
		log.Print("User is already logged in")
		http.Redirect(w, req, "/secrets", http.StatusSeeOther)
		return
	}

	if req.Method == http.MethodPost {
		un := req.FormValue("username")
		pwd := req.FormValue("password")
		user, ok := dbUsers[un]
		// Check if user exists
		if !ok {
			log.Print("User is not found.")
			http.Error(w, "Username is not found.", http.StatusForbidden)
			return
		}

		// Check if password matches
		err := bcrypt.CompareHashAndPassword(user.Password, []byte(pwd))
		if err != nil {
			http.Error(w, "Username/password do no match", http.StatusForbidden)
			return
		}

		// Create session for user that logged in
		sID := uuid.NewString()
		c := &http.Cookie{ // remember setCookie uses the addr of a cookie
			Name:  "LoginInfo",
			Value: sID,
		}
		http.SetCookie(w, c)
		dbSessions[c.Value] = user.UserName

		log.Print("Login success! Cookie set, redirecting to secrets page.")
		http.Redirect(w, req, "/secret", http.StatusSeeOther)
	}
	tpl.ExecuteTemplate(w, "login.html", nil)
}

func logout(w http.ResponseWriter, req *http.Request) {
	if !loggedIn(req) {
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
