package domain

import "errors"

var (
	ErrUserAlreadyExists = errors.New("a user with this email has already registered")
	ErrPasswordMismatch  = errors.New("password and confirm password must be the same")
	ErrInternalServer    = errors.New("something went wrong on server")
)
