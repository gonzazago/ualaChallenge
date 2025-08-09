package users

import "errors"

var (
	ErrUserNotFound      = errors.New("user not found")
	ErrMailAlreadyExists = errors.New("mail already exists")
	ErrPersistenceError  = errors.New("persistence error")
)
