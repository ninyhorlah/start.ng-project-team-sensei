package controller

import (
	"sensei-poultry-management/model"
	"sensei-poultry-management/views"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"html/template"

	"golang.org/x/crypto/bcrypt"
)

const hashCost = 8

//SignupGet renders signup ppage
func SignupGet(w http.ResponseWriter, r *http.Request){
	tmpl, _ :=template.ParseFiles("./templates/signup.html")
	tmpl.Execute(w, nil)
}



// SignupPost registers an ew user in the database
func SignupPost(w http.ResponseWriter, r *http.Request){
	// Parse and decode the request body into a new `Credentials` instance
	creds := &views.Credentials{}
	err := json.NewDecoder(r.Body).Decode(creds)
	fmt.Println(creds)
	if err != nil {
		// If there is something wrong with the request body, return a 400 status
		w.WriteHeader(http.StatusBadRequest)
		return 
	}
	// Salt and hash the password using the bcrypt algorithm
	// The second argument is the cost of hashing, which we arbitrarily set as 8 (this value can be more or less, depending on the computing power you wish to utilize)
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(creds.Password), hashCost)
	username := creds.Username
	passkey := string(hashedPassword)
	email :=creds.Email
	// Next, insert the username, along with the hashed , password into the database
	if err = model.CreateUser(username, passkey, email); err != nil {
		// If there is any issue with inserting into the database, return a 500 error
		w.WriteHeader(http.StatusInternalServerError)
		log.Fatal(err)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(creds.Username)
	// We reach this point if the credentials we correctly stored in the database, and the default status of 200 is sent back
}