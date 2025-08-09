package post

import (
	"post-service/internal/domain/post"
	"strings"
)

type CreatePostRequest struct {
	UserID string `json:"user_id"`
	Text   string `json:"text"`
}

func (c CreatePostRequest) Validate() (bool, string) {

	if strings.TrimSpace(c.UserID) == "" {
		return false, "user id is required"
	}
	if strings.TrimSpace(c.Text) == "" {
		return false, "text is required"
	}
	if len(c.Text) > 280 {
		return false, "text is too long, not exceed 280 characters"
	}
	return true, ""
}

func (c CreatePostRequest) toDomain() *post.Post {
	return &post.Post{
		UserID: c.UserID,
		Text:   c.Text,
	}
}
