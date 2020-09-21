package controller

import (
	"html/template"
	"net/http"
	"path/filepath"
	"strconv"

	"github.com/abinash393/voting-app/model"
	"github.com/gorilla/mux"
)

// PollInfo contains Title and ID
type PollInfo struct {
	Title  string
	PollID int
}

// PollView template struct
type PollView struct {
	PageName    string
	Arr         [5]PollInfo
	Prev, Next  int
	ShowPrevBtn bool
}

// OtherPolls render template on given page
func OtherPolls(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	page, _ := strconv.Atoi(vars["page"])
	skip := page
	skip--
	skip *= 5
	rows, err := model.DB.Query(`SELECT title, poll_id FROM polls LIMIT ? , 5`, skip)
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	var d PollView
	var i int
	for rows.Next() {
		var p PollInfo

		if err := rows.Scan(&p.Title, &p.PollID); err != nil {
			panic(err.Error())
		}
		d.Arr[i] = p
		i++
	}

	tmpl := template.Must(template.ParseFiles(
		filepath.Join(PublicDir, "polls-list.html")))

	d.PageName = "other"
	d.ShowPrevBtn = page != 1
	d.Next = page + 1
	d.Prev = page - 1

	tmpl.Execute(w, d)
}

// MyPolls give created polls by user
func MyPolls(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	page, _ := strconv.Atoi(vars["page"])
	skip := page
	skip--
	skip *= 5

	session, err := model.Rdb.HGetAll(model.Ctx, r.Context().Value("SID").(string)).Result()
	if err != nil {
		panic(err)
	}

	rows, err := model.DB.Query(`SELECT title, poll_id FROM polls WHERE created_by = ? LIMIT ? , 5`,
		session["userId"], skip)
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	var d PollView
	var i int
	for rows.Next() {
		var p PollInfo

		if err := rows.Scan(&p.Title, &p.PollID); err != nil {
			panic(err.Error())
		}
		d.Arr[i] = p
		i++
	}

	tmpl := template.Must(template.ParseFiles(
		filepath.Join(PublicDir, "polls-list.html")))

	d.PageName = "my"
	d.ShowPrevBtn = page != 1
	d.Next = page + 1
	d.Prev = page - 1

	tmpl.Execute(w, d)
}
