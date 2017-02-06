package user

import (
	"encoding/json"
	"mpa/model"
	"mpa/route"
)

type LoginUserView struct {
	User model.User
}

// Render implements route.View
func (v *LoginUserView) Render(ctx *route.Context) error {
	enc := json.NewEncoder(ctx.ResponseWriter)
	return enc.Encode(map[string]interface{}{
		"id": v.User.Login,
		"name": v.User.Name,
		"loginToken": v.User.LoginToken,
	})
}
