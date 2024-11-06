package services

import (
	"Course-Management/config"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

type AuthService struct{}

func NewAuthService() *AuthService {
	return &AuthService{}
}

// Login handles the OAuth authentication process
func (s *AuthService) Login(username, password string) (string, error) {
	// Prepare the request payload for OAuth server
	payload := map[string]string{
		"grant_type":    "password", // Use "password" grant type if allowed
		"username":      username,
		"password":      password,
		"client_id":     config.OAuthSettings.ClientID,
		"client_secret": config.OAuthSettings.ClientSecret,
	}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		return "", fmt.Errorf("failed to marshal request payload: %w", err)
	}

	// Send the request to the OAuth server
	resp, err := http.Post(config.OAuthSettings.TokenURL, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return "", fmt.Errorf("failed to contact OAuth server: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(resp.Body)
		return "", errors.New(fmt.Sprintf("OAuth server error: %s", body))
	}

	// Parse the response body
	var tokenResponse struct {
		AccessToken string `json:"access_token"`
		TokenType   string `json:"token_type"`
		ExpiresIn   int    `json:"expires_in"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&tokenResponse); err != nil {
		return "", fmt.Errorf("failed to parse token response: %w", err)
	}

	// Return the access token
	return tokenResponse.AccessToken, nil
}
