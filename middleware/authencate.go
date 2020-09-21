package middleware

import (
	"context"
	"net/http"
	"net/url"

	"github.com/abinash393/voting-app/model"
)

// Auth verify token protected routes
func Auth(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("sid")
		if err != nil {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}
		sidCookie, err := url.QueryUnescape(cookie.Value)
		if err != nil {
			panic(err.Error())
		}

		cmd := model.Rdb.HExists(model.Ctx, sidCookie, "email")
		ok, _ := cmd.Result()

		if ok {
			ctx := context.WithValue(
				r.Context(),
				"SID",
				sidCookie,
			)
			f(w, r.WithContext(ctx))
		} else {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
		}
	}
}
