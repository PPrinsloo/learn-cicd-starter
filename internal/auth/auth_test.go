package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name           string
		authHeader     string
		expectedKey    string
		expectingError bool
	}{
		{
			name:           "Valid Authorization Header",
			authHeader:     "ApiKey 12345",
			expectedKey:    "12345",
			expectingError: false,
		},
		{
			name:           "Invalid Authorization Header Format",
			authHeader:     "Bearer 12345",
			expectedKey:    "",
			expectingError: true,
		},
		{
			name:           "Missing Authorization Header",
			authHeader:     "",
			expectedKey:    "",
			expectingError: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			headers := http.Header{}
			if tc.authHeader != "" {
				headers.Set("Authorization", tc.authHeader)
			}

			key, err := GetAPIKey(headers)
			if (err != nil) != tc.expectingError {
				t.Errorf("expected error: %v, got: %v", tc.expectingError, err)
			}
			if key != tc.expectedKey {
				t.Errorf("expected key: %s, got: %s", tc.expectedKey, key)
			}
		})
	}
}
