package user

import (
	"encoding/json"
	"mpa/auth"
	"mpa/route"
	"net/http"
)

type CurrentUserController struct{}

func (controller CurrentUserController) Serve(ctx *route.Context) error {
	token := auth.GetToken(ctx)
	if token == nil {
		ctx.ResponseWriter.WriteHeader(http.StatusUnauthorized)
		return auth.ErrUnauthorized
	}

	enc := json.NewEncoder(ctx.ResponseWriter)
	enc.Encode(token.User)
	return nil
}
