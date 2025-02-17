package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey_MissingAuthHeader(t *testing.T) {
	// Create empty headers (no Authorization header included)
	headers := http.Header{}

	// Call the function
	apiKey, err := GetAPIKey(headers)

	// Verify that the API key is empty
	if apiKey != "" {
		t.Errorf("expected empty API key, got %q", apiKey)
	}

	// Verify that the error returned is ErrNoAuthHeaderIncluded
	if err != ErrNoAuthHeaderIncluded {
		t.Errorf("expected error %v, got %v", ErrNoAuthHeaderIncluded, err)
	}
}
