package auth

import (
	"mpa/route"
	"net/http"
	"strings"
)

type Filter struct {
	TokenDecoder *TokenDecoder
	Secret       []byte
}

func (f *Filter) PreHandle(ctx *route.Context) error {
	auth := ctx.Request.Header.Get("Authorization")
	if auth == "" {
		ctx.ResponseWriter.Header().Set("WWW-Authenticate", `Bearer realm="Login required"`)
		ctx.ResponseWriter.WriteHeader(http.StatusUnauthorized)
		return ErrUnauthorized
	}

	scheme, tokenString := parseAuth(auth)
	if !strings.EqualFold(scheme, "Bearer") {
		ctx.ResponseWriter.WriteHeader(http.StatusBadRequest)
		return ErrUnauthorized
	}
	token, err := f.TokenDecoder.Decode(f.Secret, tokenString)
	if err == ErrTokenExpired {
		ctx.ResponseWriter.Header().Set("WWW-Authenticate", `Bearer realm="Login required"`)
		ctx.ResponseWriter.WriteHeader(http.StatusUnauthorized)
		return ErrUnauthorized
	} else if err != nil {
		ctx.ResponseWriter.WriteHeader(http.StatusBadRequest)
		return ErrUnauthorized
	}
	registerToken(ctx, &token)
	return nil
}

func (f *Filter) PostHandle(ctx *route.Context) error {
	return nil
}

func parseAuth(str string) (scheme, token string) {
	parts := strings.SplitN(str, " ", 2)
	scheme = parts[0]
	token = parts[1]
	return
}
