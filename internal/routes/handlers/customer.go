package handlers

import (
	"errors"
	"github.com/gorilla/sessions"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"html/template"
	"net/http"
)

var store = sessions.NewCookieStore([]byte("mike"))

func UserHandler(w http.ResponseWriter, r *http.Request) {
	_, span := otel.Tracer("user-profile-service").Start(r.Context(), "UserProfileHandler")
	defer span.End()

	session, err := store.Get(r, "auth-session")
	if err != nil {
		http.Error(w, "Session retrieval failed.", http.StatusInternalServerError)
		return
	}

	profile, ok := session.Values["profile"]
	if !ok {
		http.Error(w, "Profile not found in session.", http.StatusInternalServerError)
		span.RecordError(errors.New("Profile not found in session"))

		return
	}

	span.SetAttributes(attribute.String("session.profile", "retrieved"))
	renderHTML(w, http.StatusOK, "user.html", profile)
}

func renderHTML(w http.ResponseWriter, status int, filename string, data interface{}) {
	tmpl, err := template.ParseFiles(filename)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}
	w.WriteHeader(status)
	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}