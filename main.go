package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/abinash393/voting-app/model"

	"github.com/abinash393/voting-app/controller"
	"github.com/abinash393/voting-app/middleware"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func init() {
	// initializing .env file
	godotenv.Load("./config/.env")

	db, err := sql.Open("mysql", os.Getenv("DB_URI"))
	if err != nil {
		panic(err.Error())
	}
	db.SetConnMaxLifetime(time.Minute)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	// Initialize the first connection to the database, to see if everything works correctly.
	if err := db.Ping(); err != nil {
		panic(err.Error())
	} else {
		log.Println("DataBase Connected")
		model.DB = db
	}
}

func main() {
	defer model.DB.Close()
	router := mux.NewRouter()
	router.Use(middleware.Logger)

	router.HandleFunc("/", controller.Index).Methods("GET")
	router.HandleFunc("/auth", controller.Auth).Methods("GET")
	router.HandleFunc("/api/v1/user/signup", controller.Signup).Methods("POST")
	router.HandleFunc("/api/v1/user/login", controller.Login).Methods("POST")

	log.Println(http.ListenAndServe(os.Getenv("PORT"), router))
}
