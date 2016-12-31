package auth

import (
	"fmt"
	jwt "github.com/dgrijalva/jwt-go"
	"mpa/model"
	"time"
)

type Token struct {
	User model.User
}

func createTokenString(user *model.User, secret []byte) (string, error) {
	now := time.Now()
	claims := jwt.StandardClaims{
		Id:        user.Login,
		IssuedAt:  now.Unix(),
		ExpiresAt: now.AddDate(0, 0, 1).Unix(),
	}
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	return jwtToken.SignedString(secret)
}

type TokenDecoder struct {
	UserDAO model.UserDAO
}

func (dec *TokenDecoder) Decode(secret []byte, tokenString string) (Token, error) {
	jwtToken, err := jwt.Parse(tokenString, func(jwtToken *jwt.Token) (interface{}, error) {
		if _, ok := jwtToken.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", jwtToken.Header["alg"])
		}
		return secret, nil
	})
	if err != nil {
		return Token{}, err
	}

	claims, ok := jwtToken.Claims.(jwt.StandardClaims)
	if !ok {
		return Token{}, fmt.Errorf("Invalid token")
	}
	if err := claims.Valid(); err == nil {
		user, err := dec.UserDAO.FindByLogin(claims.Id)
		if err != nil {
			return Token{}, err
		}
		return Token{
			User: user,
		}, nil
	} else {
		if ve, ok := err.(jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return Token{}, ErrTokenExpired
			}
		}
		return Token{}, err
	}
}
