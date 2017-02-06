package user

import (
	"mpa/route"
	"mpa/model"
	"mpa/filter"
	"strings"
	view "mpa/view/user"
)

type UserController struct {
	UserDAO model.UserDAO
}

// ServeGet implements route.GetController
func (c *UserController) ServeGet(ctx *route.Context) (route.View, error) {
	path := ctx.Request.URL.Path
	i := strings.LastIndex(path, "/")
	id := path[i+1:]
	user, err := c.UserDAO.FindByLogin(id)
	if err != nil {
		return nil, err
	}
	token := filter.GetToken(ctx)
	if model.SameUser(user, token.User) {
		return &view.LoginUserView{
			User: user,
		}, nil
	} else {
		return &view.UserView{
			User: user,
		}, nil
	}
}
