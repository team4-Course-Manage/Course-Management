package config

import (
	"os"
)

type OAuthConfig struct {
	ClientID     string
	ClientSecret string
	AuthURL      string
	TokenURL     string
	RedirectURL  string
	Scopes       []string
}

var OAuthSettings *OAuthConfig

func LoadOAuthConfig() {
	OAuthSettings = &OAuthConfig{
		ClientID:     os.Getenv("OAUTH_CLIENT_ID"),
		ClientSecret: os.Getenv("OAUTH_CLIENT_SECRET"),
		AuthURL:      os.Getenv("OAUTH_AUTH_URL"),
		TokenURL:     os.Getenv("OAUTH_TOKEN_URL"),
		RedirectURL:  os.Getenv("OAUTH_REDIRECT_URL"),
		Scopes:       []string{"profile", "email"},
	}
}
