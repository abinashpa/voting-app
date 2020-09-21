package controller

import (
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
	"strconv"

	"github.com/abinash393/voting-app/model"
	"github.com/gorilla/mux"
)

// Poll type
type Poll struct {
	Title, Option1, Option2, Option3, Option4, Option5                          string
	Option1Vote, Option2Vote, Option3Vote, Option4Vote, Option5Vote, TotalVotes int
	IsOpen, IsVoted, PollNum                                                    int
}

// ViewPolls render poll on given ID
func ViewPolls(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	session, err := model.Rdb.HGetAll(model.Ctx, r.Context().Value("SID").(string)).Result()
	if err != nil {
		panic(err)
	}

	var p Poll
	if err := model.DB.QueryRow(fmt.Sprintf(`SELECT title, option1, option2, option3, option4, option5,
	option1_vote, option2_vote, option3_vote, option4_vote, option5_vote, total_votes,
	BIN(poll_status), JSON_CONTAINS_PATH(user_voted, 'one', '$."%s"') FROM polls WHERE poll_id = ?`, session["userId"]),
		vars["id"]).Scan(&p.Title,
		&p.Option1, &p.Option2, &p.Option3, &p.Option4,
		&p.Option5, &p.Option1Vote, &p.Option2Vote, &p.Option3Vote, &p.Option4Vote,
		&p.Option5Vote, &p.TotalVotes, &p.IsOpen, &p.IsVoted,
	); err != nil {
		panic(err.Error())
	}
	p.PollNum, _ = strconv.Atoi(vars["id"])

	w.Header().Set("Content-Type", "text/html")
	if p.IsVoted == 1 {
		http.ServeFile(w, r, filepath.Join(PublicDir, "already-voted.html"))
	} else if p.IsOpen == 0 {
		http.ServeFile(w, r, filepath.Join(PublicDir, "voteing-ended.html"))
	} else {
		tmpl := template.Must(template.ParseFiles(
			filepath.Join(PublicDir, "vote-poll.html")))
		tmpl.Execute(w, p)
	}
}
