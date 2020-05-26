package main

import (
	"github.com/startng/sensei-poultry-management/controller"
	"github.com/startng/sensei-poultry-management/model"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	//"path/filepath" // so that we can make path joins compatible on all OS
)


func main() {
	err := godotenv.Load()
	if err != nil{
		log.Println("Error loading .env file")
	}
	port := os.Getenv("PORT")
	if port ==""{
		port = strconv.Itoa(8000)
	}
	  
	// establishes a path for serving static file
	fs := http.FileServer(http.Dir("./templates"))

	// Registering routes and handler that we will implement
	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/", controller.Signin).Methods("GET") // root address renders signin page
	r.HandleFunc("/login", controller.Login).Methods("POST")
	r.HandleFunc("/logout", controller.Logout)
	//r.HandleFunc("/refresh", controller.Refresh)
	r.HandleFunc("/forbidden", controller.Forbidden)
	//r.HandleFunc("/secret", controller.Secret)
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))
	
	// initialize our Postgres database connection
	db := model.InitDB()
	defer db.Close()


	fmt.Printf("Listening and serving on port %s.....", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}