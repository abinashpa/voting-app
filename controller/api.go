package controller

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"golang.org/x/crypto/bcrypt"

	"github.com/abinash393/voting-app/model"
)

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
		panic(err.Error())
	}

	var hashPassword string
	if err := model.DB.QueryRow(`SELECT password FROM users WHERE email = $`, userInfo.Email).Scan(&hashPassword); err != nil {
		panic(err.Error())
	}

	res.Header().Set("Content-Type", "application/json")
	if err := bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(userInfo.Password)); err != nil {

		res.WriteHeader(http.StatusOK)
		json.NewEncoder(res).Encode(Response{true, ""})
	} else {
		res.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(res).Encode(Response{})
	}
}
