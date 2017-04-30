package filter

import (
	"errors"
	"mpa/model"
	"mpa/route"
	"net/http"
	"strings"
)

var (
	ErrUnauthenticated = errors.New("mpa/filter: User is not authenticated")
)

type LoginFilter struct {
	Authenticator *model.Authenticator
	Secret       []byte
}

func (f *LoginFilter) PreHandle(ctx *route.Context) error {
	auth := ctx.Request.Header.Get("Authorization")
	if auth == "" {
		ctx.ResponseWriter.Header().Set("WWW-Authenticate", `Bearer realm="Login required"`)
		ctx.ResponseWriter.WriteHeader(http.StatusUnauthorized)
		return ErrUnauthenticated
	}

	scheme, tokenString := parseAuth(auth)
	if !strings.EqualFold(scheme, "Bearer") {
		ctx.ResponseWriter.WriteHeader(http.StatusBadRequest)
		return ErrUnauthenticated
	}
	authResult, err := f.Authenticator.Authenticate(f.Secret, tokenString)
	if err == model.ErrTokenExpired {
		ctx.ResponseWriter.Header().Set("WWW-Authenticate", `Bearer realm="Login required"`)
		ctx.ResponseWriter.WriteHeader(http.StatusUnauthorized)
		return ErrUnauthenticated
	} else if err != nil {
		ctx.ResponseWriter.WriteHeader(http.StatusBadRequest)
		return ErrUnauthenticated
	}
	registerAuthResult(ctx, &authResult)
	return nil
}

func (f *LoginFilter) PostHandle(ctx *route.Context) error {
	return nil
}

func parseAuth(str string) (scheme, token string) {
	parts := strings.SplitN(str, " ", 2)
	scheme = parts[0]
	token = parts[1]
	return
}

const tokenAttributeName = "auth/result"

func registerAuthResult(ctx *route.Context, token *model.AuthResult) {
	ctx.PutAttribute(tokenAttributeName, token)
}

func GetAuthResult(ctx *route.Context) *model.AuthResult {
	obj := ctx.GetAttribute(tokenAttributeName)
	if token, ok := obj.(*model.AuthResult); ok {
		return token
	} else {
		return nil
	}
}
