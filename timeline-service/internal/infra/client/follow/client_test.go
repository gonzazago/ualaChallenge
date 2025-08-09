package follow

import (
	"context"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestFollowClient_GetFollowing(t *testing.T) {
	type want struct {
		following *Followers
		wantErr   bool
	}
	tests := []struct {
		name   string
		server *httptest.Server
		want   want
	}{
		{
			name: "Success",
			server: httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, "/api/v1/users/user-1/following", r.URL.Path)
				w.Header().Set("Content-Type", "application/json")
				json.NewEncoder(w).Encode(
					Followers{Followers: []string{"user-2", "user-3"}},
				)
			})),
			want: want{
				following: &Followers{Followers: []string{"user-2", "user-3"}},
				wantErr:   false,
			},
		},
		{
			name: "Error - Server returns 500",
			server: httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusInternalServerError)
			})),
			want: want{
				following: &Followers{},
				wantErr:   true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer tt.server.Close()
			client := NewFollowClient(tt.server.Client(), tt.server.URL)

			got, err := client.GetFollowing(context.Background(), "user-1")

			if tt.want.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, tt.want.following, got)
		})
	}
}
