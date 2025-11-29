package restclient

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/assert"
)

func TestMessageAPI_Get(t *testing.T) {
	// Create a mock server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/messages/123" {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"text": "Hello World"}`))
			return
		}
		if r.URL.Path == "/messages/404" {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer server.Close()

	// Override the URL for testing
	originalURL := messageAPIURL
	messageAPIURL = server.URL + "/messages/%s"
	defer func() { messageAPIURL = originalURL }()

	client := resty.New()
	api := NewMessageAPI(client)

	tests := []struct {
		name    string
		id      string
		want    *Message
		wantErr bool
	}{
		{
			name:    "Success",
			id:      "123",
			want:    &Message{Text: "Hello World"},
			wantErr: false,
		},
		{
			name:    "Not Found",
			id:      "404",
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := api.Get(context.Background(), tt.id)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, got)
			}
		})
	}
}
