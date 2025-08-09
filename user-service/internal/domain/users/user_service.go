package users

import (
	"context"
)

// Service define la interfaz para el servicio de usuario.
// Esta capa se sitúa entre los handlers y los repositorios.
type Service interface {
	Create(ctx context.Context, user *User) (*User, error)
	GetByID(ctx context.Context, id string) (*User, error)
}

type Repository interface {
	Save(ctx context.Context, user *User) (*User, error)
	FindByID(ctx context.Context, id string) (*User, error)
	FindByEmail(ctx context.Context, email string) (*User, error)
}

type service struct {
	repo Repository
}

// NewService crea una nueva instancia del servicio de usuario.
func NewService(repo Repository) Service {
	return &service{
		repo: repo,
	}
}
