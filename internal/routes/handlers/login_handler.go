package handlers

import (
	"crypto/rand"
	"encoding/base64"
	"net/http"
	"net/url"
	"os"

	"github.com/Domains18/SIL-backend/pkg/auth"
	"github.com/gorilla/sessions"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
)

func LoginHandler(auth  authenticator.Authenticator,  store *sessions.CookieStore) http.HandlerFunc {
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

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	_, span := otel.Tracer("auth-service").Start(r.Context(), "LogoutHandler")
	defer span.End()

	logoutURL, err := url.Parse("https://" + os.Getenv("AUTH0_DOMAIN") + "/v2/logout")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		span.RecordError(err)

		return
	}
	scheme := "http"
	if r.TLS != nil {
		scheme = "https"
	}

	returnTo, err := url.Parse(scheme + "://" + r.Host)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		span.RecordError(err)
		return
	}

	parameters := url.Values{}
	parameters.Add("returnTo", returnTo.String())
	parameters.Add("client_id", os.Getenv("AUTH0_CLIENT_ID"))
	logoutURL.RawQuery = parameters.Encode()

	span.SetAttributes(attribute.String("logout.redirect", logoutURL.String()))
	http.Redirect(w, r, logoutURL.String(), http.StatusTemporaryRedirect)
}