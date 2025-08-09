package post

import (
	"context"
	"log"
	"post-service/internal/infra/queue"
)

// CreatePost create a new post.
func (s *service) CreatePost(ctx context.Context, post *Post) (*Post, error) {

	if err := s.repo.Save(ctx, post); err != nil {
		log.Println("Error creating post:", err)
		return nil, ErrPersistenceError
	}

	go func() {
		err := s.notifier.Send(context.Background(), queue.Event{
			Name: "PostCreated",
			Payload: map[string]interface{}{
				"post_id": post.ID,
				"user_id": post.UserID,
			},
		})
		if err != nil {
			log.Println("notifier.Send", err)
		}
	}()
	return post, nil
}
