package model

import "sensei-poultry-management/views"

// CreateUser creates a new user
func CreateUser(username, email, password string) error {
	// before creating user check if the user name exits
	if _, err := db.Query("INSERT INTO forthebirds (username, email, password) VALUES ($1,$2,$3)", username, email, password); err != nil {
		return err
	}

	return nil
}

// GetUserCredential gets the existing entry present in the database for the given username
func GetUserCredential(username string) (*views.Credentials, error) {
	result := db.QueryRow("SELECT password FROM forthebirds WHERE username= $1", username)
	// We create another instance of `Credentials` to store the credentials we get from the database
	hashedCreds := &views.Credentials{}

	// Store the obtained password in `storedCreds`
	err := result.Scan(&hashedCreds.Password)
	if err != nil {
		return nil, err
	}

	return hashedCreds, nil
}
