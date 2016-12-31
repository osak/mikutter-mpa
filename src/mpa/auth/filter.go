package auth

import (
	"net/http"
	"strings"
)

type Filter struct {
	TokenDecoder *TokenDecoder
	Secret       []byte
}

func (f *Filter) PreHandle(w http.ResponseWriter, req *http.Request) bool {
	auth := req.Header.Get("Authorization")
	if auth == "" {
		w.Header().Set("WWW-Authenticate", `Bearer realm="Login required"`)
		w.WriteHeader(http.StatusUnauthorized)
		return false
	}

	scheme, tokenString := parseAuth(auth)
	if strings.EqualFold(scheme, "Bearer") {
		w.WriteHeader(http.StatusBadRequest)
		return false
	}
	_, err := f.TokenDecoder.Decode(f.Secret, tokenString)
	if err == ErrTokenExpired {
		w.Header().Set("WWW-Authenticate", `Bearer realm="Login required"`)
		w.WriteHeader(http.StatusUnauthorized)
		return false
	} else if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return false
	}

	return true
}

func (f *Filter) PostHandle(w http.ResponseWriter, req *http.Request) bool {
	return true
}

func parseAuth(str string) (scheme, token string) {
	parts := strings.SplitN(str, " ", 2)
	scheme = parts[0]
	token = parts[1]
	return
}
