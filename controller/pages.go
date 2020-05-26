package controller

import (
	"github.com/startng/sensei-poultry-management/views"
	"fmt"
	"net/http"
	"github.com/gorilla/sessions"
)


// Signin renders the signin page
func Signin(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hello word")
	//w.Write([]byte("Hello world"))
	w.Header().Set("Content-Type", "text/html; charset=utf8")
	http.Redirect(w, r, "/static/sign_in.html", http.StatusSeeOther)
	

} 

// getUser returns a user from session s
// on error returns an empty user
func getUser(s *sessions.Session) views.User {
	val := s.Values["user"]

	var user = views.User{}
	user, ok := val.(views.User)
	if !ok {
		return views.User{Authenticated: false}
	}
	return user
}

// Forbidden handles invalid requests todo: improve implementation
func Forbidden(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "avely")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	flashMessages := session.Flashes()
	err = session.Save(r, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tpl.ExecuteTemplate(w, "forbidden.gohtml", flashMessages)
}

// // Secret renders a hidden message chore: To be removed
// func Secret(w http.ResponseWriter, r *http.Request) {
// 	session, err := store.Get(r, "avely")
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	user := getUser(session)

// 	if auth := user.Authenticated; !auth {
// 		session.AddFlash("You don't have access!")
// 		err = session.Save(r, w)
// 		if err != nil {
// 			http.Error(w, err.Error(), http.StatusInternalServerError)
// 			return
// 		}
// 		http.Redirect(w, r, "/forbidden", http.StatusFound)
// 		return
// 	}

// 	tpl.ExecuteTemplate(w, "secret.gohtml", user.Email)
// }