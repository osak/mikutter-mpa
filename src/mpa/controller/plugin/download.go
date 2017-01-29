package plugin

import (
	"mpa/model"
	"mpa/route"
	view "mpa/view/plugin"
	"strings"
)

type DownloadController struct {
	PluginDAO model.PluginDAO
}

// ServeGet implements GetController
func (c *DownloadController) ServeGet(ctx *route.Context) (route.View, error) {
	s := strings.LastIndex(ctx.Request.URL.Path, "/")
	base := ctx.Request.URL.Path[s+1 : len(ctx.Request.URL.Path)]
	slug := base[0 : len(base)-4]

	plugin, err := c.PluginDAO.FindBySlug(slug)
	if err != nil {
		return nil, err
	}
	return &view.DownloadView{
		Plugin: plugin,
	}, nil
}
