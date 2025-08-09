package timeline

import (
	"time"
	postDomain "timeline-service/internal/domain/timeline"
)

type Response struct {
	Status string `json:"status"`
	Post   []Post
}

type Post struct {
	UserID    string    `json:"user_id"`
	Text      string    `json:"text"`
	CreatedAt time.Time `json:"created_at"`
}

func toResponse(posts []postDomain.Post) Response {
	response := make([]Post, 0)
	for _, post := range posts {
		response = append(response, Post{
			UserID:    post.UserID,
			Text:      post.Text,
			CreatedAt: post.CreatedAt,
		})
	}

	return Response{
		Status: "OK",
		Post:   response,
	}
}
