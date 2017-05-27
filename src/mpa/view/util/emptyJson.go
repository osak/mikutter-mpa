package util

import (
	"mpa/route"
	"encoding/json"
)

type EmptyJsonView struct {
}

// Render implements route.View
func (v *EmptyJsonView) Render(ctx *route.Context) error {
	enc := json.NewEncoder(ctx.ResponseWriter)
	return enc.Encode(map[string]string{})
}