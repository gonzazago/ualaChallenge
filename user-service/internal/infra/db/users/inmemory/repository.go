package inmemory

import (
	"context"
	"sync"
	"user-service/internal/domain/users"
	"user-service/internal/infra/db/users/errors"
)

type InMemoryUserRepository struct {
	users map[string]*users.User
	mu    sync.RWMutex
}

func NewInMemoryUserRepository() *InMemoryUserRepository {
	return &InMemoryUserRepository{
		users: make(map[string]*users.User),
	}
}

// Save implementa el método de la interfaz para guardar un usuario en el mapa.
func (r *InMemoryUserRepository) Save(ctx context.Context, user *users.User) (*users.User, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.users[user.ID] = user
	return user, nil
}

// FindByID implementa el método para buscar un usuario por ID en el mapa.
func (r *InMemoryUserRepository) FindByID(ctx context.Context, id string) (*users.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	user, ok := r.users[id]
	if !ok {
		return nil, errors.PersistenceError
	}
	return user, nil
}

// FindByEmail implementa el método para buscar un usuario por email.
func (r *InMemoryUserRepository) FindByEmail(ctx context.Context, email string) (*users.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	for _, user := range r.users {
		if user.Email == email {
			return user, nil
		}
	}

	return nil, errors.ErrUserNotFound
}
