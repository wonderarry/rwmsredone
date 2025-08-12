package domain

import "errors"

var (
	ErrForbidden         = errors.New("forbidden")
	ErrNotFound          = errors.New("not found")
	ErrInvalidTransition = errors.New("invalid transition")
)
