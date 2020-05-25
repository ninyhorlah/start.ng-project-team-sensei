package model

import "sensei-poultry-management/views"

// CreateUser creates a new user
func CreateUser(username, password, email string) error {
	// before creating user check if the user name exits
	if _, err := db.Query("INSERT INTO forthebirds (username, password, email) VALUES ($1,$2,$3)", username, password, email); err != nil {
		return err
	}

	return nil
}

// GetUserCredential gets the existing entry present in the database for the given username
func GetUserCredential(username string) (*views.Credentials, error) {
	result := db.QueryRow("SELECT password FROM avelival WHERE username= $1", username)
	// We create another instance of `Credentials` to store the credentials we get from the database
	hashedCreds := &views.Credentials{}

	// Store the obtained password in `storedCreds`
	err := result.Scan(&hashedCreds.Password)
	if err != nil {
		return nil, err
	}

	return hashedCreds, nil
}
