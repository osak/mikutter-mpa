package util

import (
	"mpa/route"
	"net/http"
)

type RedirectView struct {
	Url string
}

// Render implements route.View
func (v *RedirectView) Render(ctx *route.Context) error {
	ctx.ResponseWriter.Header().Set("Location", v.Url)
	ctx.ResponseWriter.WriteHeader(http.StatusFound)
	return nil
}
