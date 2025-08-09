package users

import (
	"context"
	"errors"
	"log"
	db "user-service/internal/infra/db/users/errors"
)

func (s *service) GetByID(ctx context.Context, id string) (*User, error) {

	user, err := s.repo.FindByID(ctx, id)
	if err != nil {
		log.Println("Error finding user by id:", err)
		if errors.Is(err, db.ErrUserNotFound) {
			return nil, ErrUserNotFound
		}
		return nil, ErrPersistenceError
	}
	return user, nil
}
