package auth

import (
	"errors"
)

var (
	ErrTokenExpired = errors.New("Token expired")
)
