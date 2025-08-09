package post

import (
	"context"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
	"timeline-service/internal/domain/timeline"
)

func TestPostClient_GetPostsByUsers(t *testing.T) {
	fixedTime := time.Date(2025, 8, 7, 10, 0, 0, 0, time.UTC)
	mockPosts := []timeline.Post{
		{ID: "post-1", UserID: "user-2", Text: "Hello", CreatedAt: fixedTime},
	}

	type want struct {
		posts   []timeline.Post
		wantErr bool
	}
	tests := []struct {
		name   string
		server *httptest.Server
		want   want
	}{
		{
			name: "Success",
			server: httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, "/api/v1/posts", r.URL.Path)
				assert.Equal(t, "user-2,user-3", r.URL.Query().Get("user_ids"))
				w.Header().Set("Content-Type", "application/json")
				json.NewEncoder(w).Encode(mockPosts)
			})),
			want: want{
				posts:   mockPosts,
				wantErr: false,
			},
		},
		{
			name: "Error - Server returns 500",
			server: httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusInternalServerError)
			})),
			want: want{
				posts:   nil,
				wantErr: true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer tt.server.Close()
			client := NewPostClient(tt.server.Client(), tt.server.URL)

			got, err := client.GetPostsByUsers(context.Background(), []string{"user-2", "user-3"})

			if tt.want.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, tt.want.posts, got)
		})
	}
}
