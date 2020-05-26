package model

// CreateUser creates a new user
func CreateUser(username, email, password string) error {
	// before creating user check if the user name exits
	if _, err := db.Query("INSERT INTO user_info (username, email, password) VALUES ($1,$2,$3)", username, email, password); err != nil {
		return err
	}

	return nil
}