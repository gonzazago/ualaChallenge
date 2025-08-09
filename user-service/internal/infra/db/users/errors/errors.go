package errors

import "errors"

var (
	PersistenceError      = errors.New("persistence error")
	ErrUserNotFound       = errors.New("usuario no encontrado")
	ErrEmailAlreadyExists = errors.New("email already exists")
)
