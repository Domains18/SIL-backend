package handlers

import (
	"crypto/rand"
	"encoding/base64"
	"net/http"

	"github.com/Domains18/food-delivery/pkg/authenticator"
	"github.com/gorilla/sessions"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
)

func LoginHandler(auth *authenticator.Authenticator, store *sessions.CookieStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, span := otel.Tracer("user-login-service").Start(r.Context(), "LoginHandler")
		defer span.End()
		state, err := generateRandomState()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			span.RecordError(err)

			return
		}

		session, err := store.Get(r, "session-name")
		if err != nil {
			http.Error(w, "Session retrieval failed.", http.StatusInternalServerError)
			span.RecordError(err)

			return
		}

		session.Values["state"] = state
		if err := session.Save(r, w); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			span.RecordError(err)
			return
		}
		span.SetAttributes(attribute.String("session.state", state))
		http.Redirect(w, r, auth.AuthCodeURL(state), http.StatusTemporaryRedirect)
	}
}

func generateRandomState() (string, error) {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}

	state := base64.StdEncoding.EncodeToString(b)

	return state, nil
}
