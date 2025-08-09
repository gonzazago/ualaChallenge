package inmemory

import (
	"context"
	"sync"
)

// InMemoryFollowRepository almacena las relaciones de seguimiento en un mapa.
// La estructura es: map[followerID] -> map[followingID] -> bool
type InMemoryFollowRepository struct {
	follows map[string]map[string]bool
	mu      sync.RWMutex
}

// NewInMemoryFollowRepository crea una nueva instancia del repositorio.
func NewInMemoryFollowRepository() *InMemoryFollowRepository {
	return &InMemoryFollowRepository{
		follows: make(map[string]map[string]bool),
	}
}

// Follow crea la relación de seguimiento.
func (r *InMemoryFollowRepository) Follow(ctx context.Context, followerID, followingID string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, ok := r.follows[followerID]; !ok {
		r.follows[followerID] = make(map[string]bool)
	}
	r.follows[followerID][followingID] = true
	return nil
}

// GetFollowing devuelve la lista de IDs de usuarios seguidos.
func (r *InMemoryFollowRepository) GetFollowing(ctx context.Context, userID string) ([]string, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	followingMap, ok := r.follows[userID]
	if !ok {
		return []string{}, nil // Devuelve una lista vacía si no sigue a nadie.
	}

	followingList := make([]string, 0, len(followingMap))
	for id := range followingMap {
		followingList = append(followingList, id)
	}
	return followingList, nil
}

// IsFollowing comprueba si un usuario ya sigue a otro.
func (r *InMemoryFollowRepository) IsFollowing(ctx context.Context, followerID, followingID string) (bool, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	if followingMap, ok := r.follows[followerID]; ok {
		if _, isFollowing := followingMap[followingID]; isFollowing {
			return true, nil
		}
	}
	return false, nil
}
