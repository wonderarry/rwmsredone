package domain

import "errors"

var (
	ErrForbidden         = errors.New("forbidden")
	ErrNotFound          = errors.New("not found")
	ErrInvalidTransition = errors.New("invalid transition")
	ErrNotImplemented    = errors.New("not implemented")

	ErrAccAlreadyExists   = errors.New("account already exists")
	ErrInvalidCredentials = errors.New("invalid credentials")

	ErrUnauthorized = errors.New("unauthorized")

	ErrUnknownRole = errors.New("unknown role")
)
