package inmemory

import (
	"context"
	"post-service/internal/domain/post"
	"sync"
)

// InMemoryPostRepository almacena los posts en una lista.
type InMemoryPostRepository struct {
	posts []*post.Post
	mu    sync.RWMutex
}

// NewInMemoryPostRepository crea una nueva instancia del repositorio.
func NewInMemoryPostRepository() *InMemoryPostRepository {
	return &InMemoryPostRepository{
		posts: make([]*post.Post, 0),
	}
}

// Save guarda un nuevo post.
func (r *InMemoryPostRepository) Save(ctx context.Context, p *post.Post) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.posts = append(r.posts, p)
	return nil
}

// FindByUserIDs busca posts de una lista de usuarios.
func (r *InMemoryPostRepository) FindByUserIDs(ctx context.Context, userIDs []string) ([]*post.Post, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	// Crear un set de userIDs para búsquedas rápidas.
	userIDsSet := make(map[string]bool)
	for _, id := range userIDs {
		userIDsSet[id] = true
	}

	var foundPosts []*post.Post
	for _, p := range r.posts {
		if _, ok := userIDsSet[p.UserID]; ok {
			foundPosts = append(foundPosts, p)
		}
	}

	return foundPosts, nil
}
