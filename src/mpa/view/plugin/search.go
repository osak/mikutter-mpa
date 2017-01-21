package plugin

import (
	"encoding/json"
	"mpa/model"
	"mpa/route"
)

type SearchView struct {
	Plugins []model.Plugin
}

// Render implements route.View
func (v *SearchView) Render(ctx *route.Context) error {
	arr := make([]map[string]interface{}, len(v.Plugins))
	for i, p := range v.Plugins {
		arr[i] = map[string]interface{}{
			"author":      p.Author.Name,
			"name":        p.Name,
			"version":     p.Version,
			"description": p.Description,
			"url":         p.Url,
		}
	}
	enc := json.NewEncoder(ctx.ResponseWriter)
	enc.Encode(arr)
	return nil
}
