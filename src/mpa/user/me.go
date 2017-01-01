package user

import (
	"encoding/json"
	"mpa/auth"
	"mpa/route"
	"net/http"
)

type CurrentUserController struct{}

func (controller *CurrentUserController) ServeGet(ctx *route.Context) error {
	token := auth.GetToken(ctx)
	if token == nil {
		ctx.ResponseWriter.WriteHeader(http.StatusUnauthorized)
		return auth.ErrUnauthorized
	}

	ctx.ResponseWriter.Header().Set("Content-Type", "application/json")
	enc := json.NewEncoder(ctx.ResponseWriter)
	enc.Encode(token.User)
	return nil
}

func (controller *CurrentUserController) ServePost(ctx *route.Context) error {
	return route.ErrMethodNotAllowed
}
