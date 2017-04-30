package model

import (
	"errors"
	"github.com/google/uuid"
)

var (
	ErrTokenExpired = errors.New("mpa/model: Token expired")
	ErrInvalidToken = errors.New("mpa/model: Invalid token string")
)

func CreateTokenString(user User, secret []byte) (string, error) {
	if generated, err := uuid.NewRandom(); err != nil {
		return "", err
	} else {
		return generated.String(), nil
	}
}

type AuthResult struct {
	User User
}

type Authenticator struct {
	UserDAO UserDAO
}

func (auth *Authenticator) Authenticate(secret []byte, tokenString string) (AuthResult, error) {
	if user, err := auth.UserDAO.FindByLoginToken(tokenString); err != nil {
		return AuthResult{}, err
	} else {
		return AuthResult{user}, nil
	}
}
