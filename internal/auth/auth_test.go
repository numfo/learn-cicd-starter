package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name       string
		headers    http.Header
		wantAPIKey string
		wantErrMsg string
	}{
		{
			name:       "No Authorization Header",
			headers:    http.Header{},
			wantAPIKey: "",
			wantErrMsg: "no authorization header included",
		},
		{
			name: "Malformed Authorization Header",
			headers: http.Header{
				"Authorization": []string{"Bearer abc123"},
			},
			wantAPIKey: "",
			wantErrMsg: "malformed authorization header",
		},
		{
			name: "Valid Authorization Header",
			headers: http.Header{
				"Authorization": []string{"ApiKey my-secret-key"},
			},
			wantAPIKey: "my-secret-key",
			wantErrMsg: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			apiKey, err := GetAPIKey(tt.headers)
			if apiKey != tt.wantAPIKey {
				t.Errorf("expected API key '%s', got '%s'", tt.wantAPIKey, apiKey)

			}
			if tt.wantErrMsg == "" && err != nil {
				t.Errorf("expected no error, got '%v'", err)
			} else if tt.wantErrMsg != "" {
				if err == nil {
					t.Errorf("expected error '%s', got nil", tt.wantErrMsg)
				} else if err.Error() != tt.wantErrMsg {
					t.Errorf("expected error '%s', got '%s'", tt.wantErrMsg, err.Error())
		}
			}
		})
	}
}
