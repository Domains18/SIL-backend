package handlers

import (
	"net/http"

	"github.com/Domains18/SIL-backend/pkg/auth"
	"github.com/gorilla/sessions"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
)



func CallBackHandler(auth auth.Auth, store *sessions.CookieStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request){
		_, span := otel.Tracer("auth-callback-service").Start(r.Context(), "CallBackHandler")
		defer span.End()

		session, err := store.Get(r, "session-name")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			span.RecordError(err)
			return
		}
		if r.URL.Query().Get("state") != session.Values["state"] {
			http.Error(w, "state did not match", http.StatusBadRequest)
			span.RecordError(err)
			return
		}

		token, err := auth.Exchange(r.Context(), r.URL.Query().Get("code"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			span.RecordError(err)
			return
		}

		idToken, err := auth.VerifyIDToken(r.Context(), token)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			span.RecordError(err)
			return
		}

		var profile map[string]interface{}
		if err := idToken.Claims(&profile); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			span.RecordError(err)
			return
		}
		session.Values["access_token"] = token.AccessToken
		session.Values["profile"] = profile
		
		if err := session.Save(r, w); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			span.RecordError(err)
			return
		}
		span.SetAttributes(attribute.String("session.access_token", "set"))
		span.SetAttributes(attribute.String("session.profile", "set"))
		http.Redirect(w, r, "/user.html", http.StatusTemporaryRedirect)
	}
}