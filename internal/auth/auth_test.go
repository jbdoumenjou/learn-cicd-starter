package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name    string
		headers http.Header
		want    string
		wantErr bool
	}{
		{
			name: "no auth header",
			headers: http.Header{
				"Authorization": []string{},
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "malformed auth header",
			headers: http.Header{
				"Authorization": []string{"Bearer"},
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "valid auth header",
			headers: http.Header{
				"Authorization": []string{"ApiKey 123"},
			},
			want:    "123",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetAPIKey(tt.headers)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAPIKey() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetAPIKey() got = %v, want %v", got, tt.want)
			}
		})
	}
}
