package model

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	//"strconv"
	"github.com/joho/godotenv"
)

// The "db" package level variable will hold the reference to our database instance
var db *sql.DB

//InitDB establishes the database connection
func InitDB() *sql.DB {
	err:= godotenv.Load()
	if err!= nil{
		log.Println("Error loading .env file")
	}
	host := os.Getenv("host")
	//port,_ := strconv.Atoi(os.Getenv("dbport"))
	user:= os.Getenv("user")
	password:= os.Getenv("password")
	dbname:= os.Getenv("dbname")
	//connectionURL := os.Getenv("DATABASE_URL")
	

	// Connect to the postgres db
	//you might have to change the connection string to add your database credentials
	psqlinfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, 5432 , user ,password, dbname)
	//db, err = sql.Open("postgres", connectionURL)
	db, err = sql.Open("postgres", psqlinfo)

	log.Println("connecting to db")
	if err != nil {
		log.Println("Database could not connect",err)
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err) 
	}
	log.Println("Connected to database")
	return db
}
