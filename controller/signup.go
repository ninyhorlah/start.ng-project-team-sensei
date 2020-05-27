package controller

import (
	"html/template"
	"log"
	"net/http"

	"github.com/startng/sensei-poultry-management/model"
	"golang.org/x/crypto/bcrypt"
)

const hashCost = 8

//Signup sWitches between requests get and post in SignupGet and SignupPost
func Signup(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		SignupGet(w, r)
	case "POST":
		SignupPost(w, r)
		http.Redirect(w, r, "/login", http.StatusFound)

	}
}

//SignupGet renders the signup ppage
func SignupGet(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("./templates/index.html")
	tmpl.Execute(w, nil)
}

// SignupPost registers a new user in the database
// SignupPost registers a new user in the database
func SignupPost(w http.ResponseWriter, r *http.Request) {
	// Parse and decode the request body into a new `Credentials` instance

	username := r.FormValue("name")
	email := r.FormValue("email")
	password := r.FormValue("password")

	// Salt and hash the password using the bcrypt algorithm
	// The second argument is the cost of hashing, which we arbitrarily set as 8 (this value can be more or less, depending on the computing power you wish to utilize)
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), hashCost)
	passkey := string(hashedPassword)
	// Next, insert the username, along with the email hashed , password into the database
	if err = model.CreateUser(username, email, passkey); err != nil {
		// If there is any issue with inserting into the database, return a 500 error
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}
	w.Header().Set("Content-Type", "text/html;charset=utf8")
	// We reach this point if the credentials were correctly stored in the database, and the default status of 200 is sent back
}
