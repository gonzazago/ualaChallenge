package post

import (
	"post-service/internal/domain/post"
	"time"
)

type PostResponse struct {
	ID        string    `json:"id"`
	UserID    string    `json:"user_id"`
	Text      string    `json:"text"`
	CreatedAt time.Time `json:"created_at"`
}

func toResponse(post *post.Post) PostResponse {
	return PostResponse{
		ID:        post.ID,
		UserID:    post.UserID,
		Text:      post.Text,
		CreatedAt: post.CreatedAt,
	}
}
