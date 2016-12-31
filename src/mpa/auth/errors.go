package auth

import (
	"errors"
)

var (
	ErrTokenExpired = errors.New("Token expired")
	ErrUnauthorized = errors.New("Unauthorized")
	ErrInvalidToken = errors.New("Invalid token")
)
