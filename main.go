package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	con "github.com/abinash393/voting-app/controller"
	mid "github.com/abinash393/voting-app/middleware"
	"github.com/abinash393/voting-app/model"

	"github.com/go-redis/redis/v8"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func init() {
	if dir, err := filepath.Abs(filepath.Dir(os.Args[0])); err != nil {
		panic(err.Error)
	} else {
		con.PublicDir = filepath.Join(dir, "public")
	}

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
	// init redis server
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	if pong, err := rdb.Ping(model.Ctx).Result(); err != nil {
		log.Println(err.Error())
	} else {
		model.Rdb = rdb
		log.Println(pong)
	}
}

func main() {
	defer model.DB.Close()

	r := mux.NewRouter()
	r.Use(mid.Logger)
	// r.Use(mid.Recover)

	// static
	r.HandleFunc("/", mid.Auth(con.Index)).Methods("GET")
	r.HandleFunc("/signup", con.SignupPage).Methods("GET")
	r.HandleFunc("/login", con.LoginPage).Methods("GET")
	// template
	r.HandleFunc("/polls/view/{id:[0-9]+}", mid.Auth(con.ViewPolls)).Methods("GET")
	r.HandleFunc("/polls/my/{page:[0-9]+}", mid.Auth(con.MyPolls)).Methods("GET")
	r.HandleFunc("/polls/other/{page:[0-9]+}", con.OtherPolls).Methods("GET")
	// rest api
	r.HandleFunc("/api/v1/user/signup", con.Signup).Methods("POST")
	r.HandleFunc("/api/v1/user/login", con.Login).Methods("POST")
	r.HandleFunc("/api/v1/polls/create", mid.Auth(con.CreatePoll)).Methods("POST")
	r.HandleFunc("/api/v1/polls/vote/{page:[0-9]+}/{option}",
		mid.Auth(con.VoteSubmit)).Methods("POST")

	go log.Println(http.ListenAndServeTLS(os.Getenv("SPORT"),
		"/etc/letsencrypt/live/abinash.tech/fullchain.pem",
		"/etc/letsencrypt/live/abinash.tech/privkey.pem",
		r))

	log.Println(http.ListenAndServe(os.Getenv("PORT"), http.HandlerFunc(con.RedirectTLS)))
}
