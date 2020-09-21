package controller

import (
	"fmt"
	"net/http"

	"github.com/abinash393/voting-app/model"

	"github.com/gorilla/mux"
)

// VoteSubmit controller registers a vote
func VoteSubmit(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	session, err := model.Rdb.HGetAll(model.Ctx, r.Context().Value("SID").(string)).Result()
	if err != nil {
		panic(err)
	}

	var totalVotes, isOpen int
	model.DB.QueryRow("SELECT total_votes, BIN(poll_status) FROM polls WHERE poll_id = ?",
		vars["page"]).Scan(&totalVotes, &isOpen)

	if totalVotes > 12 {
		if _, err := model.DB.Exec(`UPDATE polls SET poll_status=0 WHERE poll_id = ?`,
			vars["page"]); err != nil {
			panic(err)
		}
		w.WriteHeader(http.StatusGone)
		return
	}

	if isOpen == 0 {
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}

	if _, err := model.DB.Exec(fmt.Sprintf(`UPDATE polls SET 
	user_voted = JSON_MERGE_PATCH(user_voted,'{"%s":"%s"}'), total_votes=total_votes+1, 
	%s_vote=%s_vote+1 WHERE poll_id = ?`,
		session["userId"], vars["option"], vars["option"],
		vars["option"]), vars["page"]); err != nil {
		panic(err.Error())
	}

	http.Redirect(w, r, "/polls/view"+vars["page"], http.StatusSeeOther)
}
