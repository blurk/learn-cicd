package auth

import (
	"errors"
	"net/http"
	"strings"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	testCases := map[string]struct {
		headers     http.Header
		expectedKey string
		expectedErr error
	}{
		"Valid API Key": {
			headers:     http.Header{"Authorization": []string{"ApiKey valid_api_key"}},
			expectedKey: "valid_api_key",
			expectedErr: nil,
		},
		"No Authorization Header": {
			headers:     http.Header{},
			expectedKey: "",
			expectedErr: ErrNoAuthHeaderIncluded,
		},
		"Malformed Header": {
			headers:     http.Header{"Authorization": []string{"Bearer invalid"}},
			expectedKey: "",
			expectedErr: errors.New("malformed authorization header"),
		},
		"Missing API Key": {
			headers:     http.Header{"Authorization": []string{"ApiKey"}},
			expectedKey: "",
			expectedErr: errors.New("malformed authorization header"),
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			key, err := GetAPIKey(tc.headers)

			if err != nil && !strings.Contains(err.Error(), tc.expectedErr.Error()) {
				t.Errorf("GetAPIKey() error mismatch (-got: %v +want: %v)\n", err.Error(), tc.expectedErr.Error())
			}

			if !strings.Contains(key, tc.expectedKey) {
				t.Errorf("GetAPIKey() key mismatch (-got: %v +want: %v)\n", key, tc.expectedKey)
			}
		})
	}
}
