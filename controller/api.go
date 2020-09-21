package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/abinash393/voting-app/model"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

var threeDays = time.Hour * 24 * 3
var secret = []byte(os.Getenv("JWT_SECRET"))

// Response struct
type Response struct {
	Ok  bool   `json:"ok,omitempty"`
	Msg string `json:"msg,omitempty"`
}

// Signup controller on "/api/v1/user/signup"
func Signup(res http.ResponseWriter, req *http.Request) {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}

	var userInfo model.User
	json.Unmarshal(body, &userInfo)

	hashPassBytes, err := bcrypt.GenerateFromPassword([]byte(userInfo.Password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}

	_, err = model.DB.Exec(`INSERT INTO users (email, password) VALUES (?, ?)`,
		userInfo.Email,
		string(hashPassBytes),
	)

	res.WriteHeader(http.StatusOK)
	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(Response{true, "Successfully Registered"})
}

// Login controller on "/api/v1/user/signup"
func Login(res http.ResponseWriter, req *http.Request) {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		panic(err.Error())
	}

	var userInfo model.User
	if err := json.Unmarshal(body, &userInfo); err != nil {
		panic("jsonUnmarshal" + err.Error())
	}

	var hashPassword string
	var UserID uint8
	if err := model.DB.QueryRow("SELECT user_id, password FROM `users` WHERE email = ?;",
		userInfo.Email).Scan(&UserID, &hashPassword); err != nil {
		panic("DB QUERY ERR" + err.Error())
	}
	res.Header().Set("Content-Type", "application/json")
	// checking the password
	if err := bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(userInfo.Password)); err == nil {
		v, err := uuid.NewRandom()
		if err != nil {
			panic(err.Error())
		}

		// saving session to redis
		model.Rdb.HMSet(model.Ctx, v.String(), "email", userInfo.Email, "userId", UserID)

		cookie := http.Cookie{
			Path:     "/",
			Name:     "sid",
			Value:    v.String(),
			Expires:  time.Now().Add(threeDays),
			HttpOnly: true,
			SameSite: http.SameSiteLaxMode,
		}
		http.SetCookie(res, &cookie)

		res.WriteHeader(http.StatusOK)
		json.NewEncoder(res).Encode(Response{true, ""})
	} else {
		fmt.Println(err)
		res.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(res).Encode(Response{false, "Wrong Credential"})
	}
}
