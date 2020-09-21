package controller

import (
	"net/http"
	"path/filepath"
)

// PublicDir is the folder where static file lives
var PublicDir = "/home/abinash/Workspace/go/src/voting-app/public"

// Index serve "/"
func Index(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "text/html")
	http.ServeFile(res, req, filepath.Join(PublicDir, "index.html"))
}

// LoginPage serve GET "/login"
func LoginPage(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "text/html")
	http.ServeFile(res, req, filepath.Join(PublicDir, "login.html"))
}

// SignupPage serve GET "/signup"
func SignupPage(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "text/html")
	http.ServeFile(res, req, filepath.Join(PublicDir, "signup.html"))
}
