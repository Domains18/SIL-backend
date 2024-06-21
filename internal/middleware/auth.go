package middleware

import (
	"github.com/gorilla/sessions"
	"net/http"
)

var store *sessions.CookieStore

// Checks if user is previously authenticated

func IsAuthenticated(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//
		session, err := store.Get(r, "session-name")
		if err != nil {
			http.Error(w, "Session retrieval failed.", http.StatusInternalServerError)
			return
		}

		if profile, ok := session.Values["profile"]; profile == nil || !ok {
			http.Redirect(w, r, "/", http.StatusSeeOther)
		} else {
			next.ServeHTTP(w, r)
		}
	}
}
