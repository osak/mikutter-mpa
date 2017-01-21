package plugin

import (
	"encoding/json"
	"mpa/model"
	"mpa/route"
)

type EntryView struct {
	Plugin model.Plugin
}

func (e *EntryView) Render(ctx *route.Context) error {
	enc := json.NewEncoder(ctx.ResponseWriter)
	enc.Encode(map[string]interface{}{
		"author":      e.Plugin.Author.Name,
		"name":        e.Plugin.Name,
		"version":     e.Plugin.Version,
		"description": e.Plugin.Description,
		"url":         e.Plugin.Url,
	})
	return nil
}
