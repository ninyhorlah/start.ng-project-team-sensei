package model

import (
	"github.com/startng/sensei-poultry-management/views"
	"fmt"
	"log"
)

// GetUserCredential gets the existing entry present in the database for the given username
func GetUserCredential(email string) (*views.Credentials,error){
	result := db.QueryRow("SELECT password FROM users_info WHERE email= $1", email)
	// We create another instance of `Credentials` to store the credentials we get from the database
	hashedCreds := &views.Credentials{}

	// Store the obtained password in `storedCreds`
	err := result.Scan(&hashedCreds.Password)
	if err != nil {
		log.Println("Error ocurred passing password", err)
		return nil,err
	}
	fmt.Println("Password returned from db")
	
	return hashedCreds,nil
}


