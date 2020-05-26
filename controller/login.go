package controller

import (
	"github.com/startng/sensei-poultry-management/model"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"html/template"
	//"github.com/twinj/uuid"
	//"time"
	"github.com/startng/sensei-poultry-management/views"
	"golang.org/x/crypto/bcrypt"
	"encoding/gob"
	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
	//"github.com/badoux/checkmail"
	//redisStore "gopkg.in/boj/redistore.v1"
)




var (
	//store *redisStore.RediStore
	
	tpl *template.Template
	err error
	store *sessions.CookieStore
)
func init() {
	authKeyOne := securecookie.GenerateRandomKey(64)
	encryptionKeyOne := securecookie.GenerateRandomKey(32)

	store = sessions.NewCookieStore(
		authKeyOne,
		encryptionKeyOne,
	)


	//store,err =  redisStore.NewRediStore(5, "tcp", ":6379","",authKeyOne,encryptionKeyOne)
	//pleasenaawor :=  NewRediStore(5, "tcp", ":6379","",authKeyOne,encryptionKeyOne)
		// if err != nil {
		// 	log.Println(err)
		// }
		// fmt.Println("Connected to cache memory")
	//defer store.Close()

	store.Options = &sessions.Options{
		MaxAge:   60 * 15,
		HttpOnly: true,
	}

	gob.Register(views.User{})

	
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}


// Login authenticates user login credentials
func Login(w http.ResponseWriter, r *http.Request){
	// Parse and decode the form
	email := r.FormValue("email")
	password := r.FormValue("password")
	fmt.Println(email)

	// create a new session for user  chore: better error handling log it properly for deployment
	session, err := store.Get(r, "avely")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// // Validate user email
	// err = checkmail.ValidateFormat(email)
	// 	if err != nil {
	// 		session.AddFlash("Enter a valid email address")
	// 		log.Println(err)
	// 		return
	// 	}
	

	
	// We get  the credentials from the database
	hashedCreds,err := model.GetUserCredential(email)
	if err != nil {
		// If an entry with the email does not exist, send an "Unauthorized"(401) status
		if err == sql.ErrNoRows {
			fmt.Println("in here")
			session.AddFlash("Email not found")
			log.Println("Email not in database",err)
			//http.Redirect(w, r, "/forbidden", http.StatusNotFound)
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		// If the error is of any other type, send a 500 status
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Somethin wrong with database",err)
		return
	}
	
	// Compare the stored hashed password, with the hashed version of the password that was received
	if err = bcrypt.CompareHashAndPassword([]byte(hashedCreds.Password), []byte(password)); err != nil {
		// write to session, todo: will probaly use to keep count of attempts 
		err = session.Save(r, w)
		if err != nil {
			log.Println("Error saving session")
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		log.Println("Error with authentication, incorrect password")
		// If the two passwords don't match, return a 401 status
		w.WriteHeader(http.StatusUnauthorized)
		// Render error message on page
		session.AddFlash("The code was incorrect")
		//http.Redirect(w, r, "/forbidden", http.StatusNotFound)
		return
	}
	 

	// Create a new struct to store User Id and authentication state
	user := &views.User{
		Email:      email,
		Authenticated: true,
	}

	// save the state of the user in the cookie
	session.Values["user"] = user

	// Save the cookie to the cookie store
	err = session.Save(r, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// On succeful login return OK and write to output chore: To be change to dashboard
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(struct {
		Status string `json:"status"`
	}{"User Logged In"})

	// todo: render dashboard

	//http.Redirect(w, r, "/secret", http.StatusFound)


	// If we reach this point, that means the users password was correct, and that they are authorized

}

// Logout revokes authentication for a user
func Logout(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "avely")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Gets the user hitting the endpoint sets it to an empty user(deleting the state) and deletes the cookie from memory
	session.Values["user"] = views.User{}
	session.Options.MaxAge = -1

	// Save the cookie back to the cookie store
	err = session.Save(r, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Redirects to signin page
	http.Redirect(w, r, "/static/sign_in.html", http.StatusFound)  // redirect to login page
}

