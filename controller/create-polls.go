package controller

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/abinash393/voting-app/model"
)

// CreatePoll by given input
func CreatePoll(w http.ResponseWriter, r *http.Request) {
	bodyBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}
	var pollInfo model.Poll
	json.Unmarshal(bodyBytes, &pollInfo)

	session, err := model.Rdb.HGetAll(model.Ctx, r.Context().Value("SID").(string)).Result()
	if err != nil {
		panic(err)
	}

	if _, err := model.DB.Exec(`INSERT INTO polls (title, option1, option2, option3, option4,
		option5, created_by, user_voted) VALUES (?, ?, ?, ?, ?, ?, ?, ?)`,
		pollInfo.Title,
		pollInfo.Option1,
		pollInfo.Option2,
		pollInfo.Option3,
		pollInfo.Option4,
		pollInfo.Option5,
		session["userId"],
		"{}",
	); err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Response{true, ""})
}
