package handlers

import (
	"errors"
	"net/http"

	"github.com/Domains18/food-delivery/pkg/authenticator"
	"github.com/gorilla/sessions"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
)

func CallBackHandler(auth *authenticator.Authenticator, store *sessions.CookieStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, span := otel.Tracer("auth-callback-service").Start(r.Context(), "CallBackHandler")
		defer span.End()

		// Retrieve the session from the store.
		session, err := store.Get(r, "session-name")
		if err != nil {
			http.Error(w, "Session retrieval failed.", http.StatusInternalServerError)
			span.RecordError(err)

			return
		}

		// Check if the state matches.
		if r.URL.Query().Get("state") != session.Values["state"] {
			http.Error(w, "Invalid state parameter.", http.StatusBadRequest)
			span.RecordError(errors.New("Invalid state parameter"))

			return
		}

		token, err := auth.Exchange(r.Context(), r.URL.Query().Get("code"))
		if err != nil {
			http.Error(w, "Failed to exchange an authorization code for a token.", http.StatusUnauthorized)
			span.RecordError(err)

			return
		}

		idToken, err := auth.VerifyIDToken(r.Context(), token)
		if err != nil {
			http.Error(w, "Failed to verify ID Token.", http.StatusInternalServerError)
			span.RecordError(err)

			return
		}

		var profile map[string]interface{}
		if err := idToken.Claims(&profile); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			span.RecordError(err)

			return
		}

		// Set the access_token and profile in the session.
		session.Values["access_token"] = token.AccessToken
		session.Values["profile"] = profile

		if err := session.Save(r, w); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		span.SetAttributes(attribute.String("session.access_token", "set"))
		span.SetAttributes(attribute.String("session.profile", "set"))
		http.Redirect(w, r, "/user.html", http.StatusTemporaryRedirect)
	}
}
