package user

import (
	"mpa/model"
	"mpa/route"
	"mpa/filter"
	"net/http"
	"mpa/view/util"
)

type TokenController struct {
	UserDAO model.UserDAO
}

// ServeDelete implements route.DeleteController
func (c *TokenController) ServeDelete(ctx *route.Context) (route.View, error) {
	authResult := filter.GetAuthResult(ctx)
	if authResult == nil {
		ctx.ResponseWriter.WriteHeader(http.StatusUnauthorized)
		return nil, filter.ErrUnauthenticated
	}
	if err := c.UserDAO.DeleteLoginToken(authResult.User); err != nil {
		return nil, err
	}
	return &util.EmptyJsonView{}, nil
}
