package token

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSalesforceTokenProvider_GetAccessToken(t *testing.T) {
	// Mock server
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"access_token": "mock_access_token"}`))
	}))
	defer mockServer.Close()

	// Initialize token provider with mock server URL
	tokenProvider := &SalesforceTokenProvider{
		ClientID:     "mock-client-id",
		ClientSecret: "mock-client-secret",
		AuthURL:      mockServer.URL,
	}

	// Test GetAccessToken
	token, err := tokenProvider.GetAccessToken()
	if err != nil {
		t.Errorf("Failed to get access token: %v", err)
	}

	if token != "mock_access_token" {
		t.Errorf("Expected access token 'mock_access_token', got '%s'", token)
	}
}

func TestSalesforceTokenProvider_GetAccessToken_Error(t *testing.T) {
	// Mock server returning an error response
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(`{"error": "invalid_client"}`))
	}))
	defer mockServer.Close()

	tokenProvider := &SalesforceTokenProvider{
		ClientID:     "mock-client-id",
		ClientSecret: "mock-client-secret",
		AuthURL:      mockServer.URL,
	}

	// Test GetAccessToken
	_, err := tokenProvider.GetAccessToken()
	if err == nil {
		t.Error("Expected error for invalid client, but got nil")
	}
}
