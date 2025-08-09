package post

import "errors"

var (
	ErrUserIDRequired   = errors.New("el userID es requerido")
	ErrTextRequired     = errors.New("el texto del post es requerido")
	ErrTextTooLong      = errors.New("el texto del post excede los 280 caracteres")
	ErrPersistenceError = errors.New("persistence error")
)
