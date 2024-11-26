package auth

import (
	"net/http"
	"testing"
)

func TestGetApiKey(t *testing.T) {
	headers := http.Header{}
	headers.Set("Authorization", "ApiKey your-api-key")

	apiKey, err := GetAPIKey(headers)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if apiKey != "your-api-key" {
		t.Errorf("Expected 'your-api-key', got %v", apiKey)
	}
}
