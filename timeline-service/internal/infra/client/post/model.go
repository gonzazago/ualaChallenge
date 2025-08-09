package post

import (
	"time"
	"timeline-service/internal/domain/timeline"
)

type Post struct {
	ID        string    `json:"id"`
	UserID    string    `json:"user_id"`
	Text      string    `json:"text"`
	CreatedAt time.Time `json:"created_at"`
}

func (post *Post) toDomain() timeline.Post {
	return timeline.Post{
		ID:        post.ID,
		UserID:    post.UserID,
		Text:      post.Text,
		CreatedAt: post.CreatedAt,
	}
}
