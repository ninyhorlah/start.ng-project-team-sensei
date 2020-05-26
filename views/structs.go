package views

// Credentials is a struct that models the structure of a user, both in the request body, and in the DB
type Credentials struct {
	Email string `json:"username" db:"email"`  //todo: create this table in the database
	Password string `json:"password" db:"password"`
}

// User holds user state information 
type User struct {
	Email string
	Authenticated bool
}