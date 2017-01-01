package plugin

import (
	"encoding/json"
	"io"
	"mpa/auth"
	"mpa/route"
	"net/http"
	"os"
	"strings"
)

type pluginController struct {
	dao PluginDAO
}

type pluginSearchController struct {
	dao PluginDAO
}

func NewPluginController(dao PluginDAO) route.Controller {
	return &pluginController{dao}
}

func (c *pluginController) ServeGet(ctx *route.Context) error {
	components := strings.SplitN(ctx.Request.URL.Path, "/", 4)
	name := components[3]
	enc := json.NewEncoder(ctx.ResponseWriter)
	ctx.ResponseWriter.Header().Set("Content-Type", "application/json")

	plugin, err := c.dao.FindPlugin(name)
	if err != nil {
		ctx.ResponseWriter.WriteHeader(http.StatusNotFound)
		enc.Encode(map[string]string{"error": err.Error()})
	} else {
		enc.Encode(plugin)
	}
	return nil
}

func (c *pluginController) ServePost(ctx *route.Context) error {
	return nil
}

func NewPluginSearchController(dao PluginDAO) route.Controller {
	return &pluginSearchController{dao}
}

func (c *pluginSearchController) ServeGet(ctx *route.Context) error {
	params := ctx.Request.URL.Query()
	filter := params["filter"][0]
	enc := json.NewEncoder(ctx.ResponseWriter)
	ctx.ResponseWriter.Header().Set("Content-Type", "application/json")

	plugins, err := c.dao.FindPlugins(filter)
	if err != nil {
		enc.Encode(map[string]string{"error": err.Error()})
	} else {
		enc.Encode(plugins)
	}
	return nil
}

func (c *pluginSearchController) ServePost(ctx *route.Context) error {
	dec := json.NewDecoder(io.TeeReader(ctx.Request.Body, os.Stdout))
	plugin := Plugin{}
	err := dec.Decode(&plugin)
	if err != nil {
		return err
	}
	token := auth.GetToken(ctx)
	plugin.UserId = token.User.Id
	err = c.dao.Create(&plugin)
	if err != nil {
		return err
	}
	ctx.ResponseWriter.Header().Set("Content-Type", "application/json")
	ctx.ResponseWriter.Write([]byte("{}"))
	return nil
}
