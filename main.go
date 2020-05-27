package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath" // so that we can make path joins compatible on all OS
	"github.com/startng/sensei-poultry-management/controller"
	"github.com/startng/sensei-poultry-management/model"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"strconv"
	"os"
	_ "github.com/lib/pq"
)

var tmpl = template.New("")

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")
	if port == ""{
		port = strconv.Itoa(8000)
	}
	fmt.Println(filepath.Join(".", "templates", "*.html"))
	fmt.Println(filepath.Join(".", "templates", "*.css"))

	fs := http.FileServer(http.Dir("templates/css/"))
	// Registering routes and handler that we will implement
	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/", controller.Signin).Methods("GET") // root address renders signin page
	r.HandleFunc("/login", controller.Login).Methods("POST")
	r.HandleFunc("/logout", controller.Logout)
	//r.HandleFunc("/refresh", controller.Refresh)
	r.HandleFunc("/forbidden", controller.Forbidden)
	//r.HandleFunc("/secret", controller.Secret)
	r.HandleFunc("/signup", controller.Signup).Methods("GET", "POST")
	r.PathPrefix("/css/").Handler(http.StripPrefix("/css/", fs))

	
	db := model.InitDB()
	defer db.Close()
	// start the server on port 8000
	fmt.Printf("Listening and serving on port %s.....", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
