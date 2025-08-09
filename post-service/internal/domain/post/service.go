package post

import (
	"context"
	"post-service/internal/infra/queue"
)

type Repository interface {
	// Save guarda un nuevo post.
	Save(ctx context.Context, post *Post) error
	// FindByUserIDs busca posts de una lista de usuarios.
	FindByUserIDs(ctx context.Context, userIDs []string) ([]*Post, error)
}
type Service interface {
	CreatePost(ctx context.Context, post *Post) (*Post, error)
	GetPostsByUsers(ctx context.Context, userIDs []string) ([]*Post, error)
}

type service struct {
	repo     Repository
	notifier queue.Notifier
}

// NewService crea una nueva instancia del servicio de posts.
func NewService(repo Repository, notifier queue.Notifier) Service {
	return &service{
		repo:     repo,
		notifier: notifier,
	}
}
