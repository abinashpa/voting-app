package controller

import (
	"net/http"
	"path/filepath"
)

// PublicDir is the folder where static file lives
var PublicDir = "/home/abinash/Workspace/go/src/voting-app/public"

// Index serve "/"
func Index(res http.ResponseWriter, req *http.Request) {

	res.WriteHeader(200)
	http.ServeFile(res, req, filepath.Join(PublicDir, "index.html"))
}

// Auth server "/auth"
func Auth(res http.ResponseWriter, req *http.Request) {
	res.WriteHeader(200)
	http.ServeFile(res, req, filepath.Join(PublicDir, "auth.html"))
}
