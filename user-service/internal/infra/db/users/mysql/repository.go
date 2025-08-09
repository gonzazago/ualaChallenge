package mysql

import (
	"context"
	"database/sql"
	"user-service/internal/domain/users"
	"user-service/internal/infra/db/users/errors"
)

type UserRepository struct {
	db *sql.DB
}

func NewMySQLUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Save(ctx context.Context, userRequest *users.User) (*users.User, error) {
	query := "INSERT INTO users (id, username, email, created_at) VALUES (?, ?, ?, ?)"
	_, err := r.db.ExecContext(ctx, query, userRequest.ID, userRequest.Username, userRequest.Email, userRequest.CreatedAt)
	if err != nil {
		return nil, errors.PersistenceError
	}
	return &users.User{}, nil
}

func (r *UserRepository) FindByID(ctx context.Context, id string) (*users.User, error) {
	query := "SELECT id, username, email, created_at FROM users WHERE id = ?"
	row := r.db.QueryRowContext(ctx, query, id)

	var user users.User
	err := row.Scan(&user.ID, &user.Username, &user.Email, &user.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.ErrUserNotFound
		}
		return nil, errors.PersistenceError
	}
	return &user, nil
}
