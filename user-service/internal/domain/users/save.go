package users

import (
	"context"
	"errors"
	"log"
	"time"
	errorsDb "user-service/internal/infra/db/users/errors"
)

func (s *service) Create(ctx context.Context, user *User) (*User, error) {

	existUser, err := s.repo.FindByEmail(ctx, user.Email)

	if err != nil && !errors.Is(err, errorsDb.ErrUserNotFound) {
		log.Println("Error finding user by email:", err)
		return nil, ErrPersistenceError
	}
	if existUser != nil {
		return nil, ErrMailAlreadyExists
	}

	now := time.Now()

	user.CreatedAt = &now

	user, err = s.repo.Save(ctx, user)

	if err != nil {
		log.Println("Error saving user:", err)
		return nil, ErrPersistenceError
	}
	return user, nil
}
