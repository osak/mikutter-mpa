package user

import (
	"mpa/filter"
	"mpa/route"
	view "mpa/view/user"
	"net/http"
)

type CurrentUserController struct{}

// ServeGet implements route.GetController
func (controller *CurrentUserController) ServeGet(ctx *route.Context) (route.View, error) {
	token := filter.GetToken(ctx)
	if token == nil {
		ctx.ResponseWriter.WriteHeader(http.StatusUnauthorized)
		return nil, filter.ErrUnauthenticated
	}

	return &view.UserView{
		User: token.User,
	}, nil
}

// ServePost implements route.PostController
func (controller *CurrentUserController) ServePost(ctx *route.Context) (route.View, error) {
	return nil, route.ErrMethodNotAllowed
}
