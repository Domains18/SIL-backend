package handlers

import (
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"net/http"
	"net/url"
	"os"
)

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
