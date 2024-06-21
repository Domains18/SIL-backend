package authenticator

import (
	"context"
	"github.com/coreos/go-oidc/v3/oidc"
	"golang.org/x/oauth2"
	"reflect"
	"testing"
)

func TestAuthenticator_VerifyIDToken(t *testing.T) {
	type fields struct {
		Provider *oidc.Provider
		Config   oauth2.Config
	}
	type args struct {
		ctx   context.Context
		token *oauth2.Token
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *oidc.IDToken
		wantErr bool
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Authenticator{
				Provider: tt.fields.Provider,
				Config:   tt.fields.Config,
			}
			got, err := a.VerifyIDToken(tt.args.ctx, tt.args.token)
			if (err != nil) != tt.wantErr {
				t.Errorf("VerifyIDToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("VerifyIDToken() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNew(t *testing.T) {
	tests := []struct {
		name    string
		want    *Authenticator
		wantErr bool
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := New()
			if (err != nil) != tt.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() got = %v, want %v", got, tt.want)
			}
		})
	}
}
