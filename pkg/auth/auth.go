package auth

import (
	"context"
	"fmt"

	"github.com/coreos/go-oidc/v3/oidc"
	"golang.org/x/oauth2"
)


type Auth struct {
	*oidc.Provider
	oauth2.Config
}


func (a *Auth) VerifyIDToken(ctx context.Context, token *oauth2.Token)(*oidc.IDToken, error) {
	rawIDToken, ok := token.Extra("id_token").(string)
	if !ok {
		return nil, fmt.Errorf("no id_token in token")
	}
	oidconfiig := &oidc.Config{
		ClientID: a.ClientID,
	}
	return a.Verifier(oidconfiig).Verify(ctx, rawIDToken)
}


func NewAuth(clientID, clientSecret, issuer string) (*Auth, error) {
	provider, err := oidc.NewProvider(context.Background(), issuer)
	if err != nil {
		return nil, err
	}
	config := oauth2.Config{
		ClientID: clientID,
		ClientSecret: clientSecret,
		Endpoint: provider.Endpoint(),
		RedirectURL: "http://localhost:8080/auth/callback",
		Scopes: []string{oidc.ScopeOpenID, "profile", "email"},
	}
	return &Auth{provider, config}, nil
}