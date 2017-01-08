package user

import (
	"encoding/json"
	"mpa/filter"
	"mpa/route"
	"net/http"
)

type CurrentUserController struct{}

func (controller *CurrentUserController) ServeGet(ctx *route.Context) error {
	token := filter.GetToken(ctx)
	if token == nil {
		ctx.ResponseWriter.WriteHeader(http.StatusUnauthorized)
		return filter.ErrUnauthenticated
	}

	ctx.ResponseWriter.Header().Set("Content-Type", "application/json")
	enc := json.NewEncoder(ctx.ResponseWriter)
	enc.Encode(token.User)
	return nil
}

func (controller *CurrentUserController) ServePost(ctx *route.Context) error {
	return route.ErrMethodNotAllowed
}
