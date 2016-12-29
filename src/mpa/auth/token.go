package auth

import (
	"fmt"
	jwt "github.com/dgrijalva/jwt-go"
	"mpa/model"
)

type Token struct {
	User model.User
}

func (token *Token) encode(secret []byte) (string, error) {
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.MapClaims{
		"login": token.User.Login,
	})
	return jwtToken.SignedString(secret)
}

func decodeToken(secret []byte, tokenString string) (Token, error) {
	jwtToken, err := jwt.Parse(tokenString, func(jwtToken *jwt.Token) (interface{}, error) {
		if _, ok := jwtToken.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", jwtToken.Header["alg"])
		}
		return secret, nil
	})
	if err != nil {
		return Token{}, err
	}

	claims, ok := jwtToken.Claims.(jwt.MapClaims)
	if ok && jwtToken.Valid {
		return Token{
			User: model.User{
				Login: claims["login"].(string),
			},
		}, nil
	} else {
		return Token{}, fmt.Errorf("Invalid token")
	}
}
