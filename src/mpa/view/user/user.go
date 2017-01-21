package user

import (
	"encoding/json"
	"mpa/model"
	"mpa/route"
)

type UserView struct {
	User model.User
}

// Render implements route.View
func (v *UserView) Render(ctx *route.Context) error {
	enc := json.NewEncoder(ctx.ResponseWriter)
	enc.Encode(v.User)
	return nil
}
