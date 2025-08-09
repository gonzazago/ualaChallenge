package follow

import (
	"context"
	"log"
)

type Service interface {
	Follow(ctx context.Context, followerID, followingID string) error
	GetFollowing(ctx context.Context, userID string) ([]string, error)
}

type Repository interface {
	// Follow crea una relación de seguimiento.
	Follow(ctx context.Context, followerID, followingID string) error
	// GetFollowing devuelve una lista de IDs de los usuarios que un usuario sigue.
	GetFollowing(ctx context.Context, userID string) ([]string, error)
	// IsFollowing comprueba si ya existe una relación de seguimiento.
	IsFollowing(ctx context.Context, followerID, followingID string) (bool, error)
}

type service struct {
	repo Repository
}

// NewService crea una nueva instancia del servicio de seguimiento.
func NewService(repo Repository) Service {
	return &service{repo: repo}
}

// Follow encapsula la lógica para crear una relación de seguimiento.
func (s *service) Follow(ctx context.Context, followerID, followingID string) error {
	if followerID == "" {
		return ErrFollowerIDRequired
	}
	if followingID == "" {
		return ErrFollowingIDRequired
	}
	if followerID == followingID {
		return ErrCannotFollowSelf
	}

	// Verificar si ya lo está siguiendo
	isFollowing, err := s.repo.IsFollowing(ctx, followerID, followingID)
	if err != nil {
		log.Println("Error checking if following is following:", err)
		return ErrPersistenceError
	}
	if isFollowing {
		return ErrAlreadyFollowing
	}

	err = s.repo.Follow(ctx, followerID, followingID)

	if err != nil {
		log.Println("Error following follower", followerID, "following", followingID)
		return ErrPersistenceError
	}

	return nil
}

// GetFollowing devuelve la lista de usuarios seguidos.
func (s *service) GetFollowing(ctx context.Context, userID string) ([]string, error) {
	if userID == "" {
		return nil, ErrFollowerIDRequired
	}
	return s.repo.GetFollowing(ctx, userID)
}
