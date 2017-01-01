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

type PluginController struct {
	PluginDAO PluginDAO
}

type PluginEntryController struct {
	PluginDAO PluginDAO
}

func (c *PluginController) ServeGet(ctx *route.Context) error {
	params := ctx.Request.URL.Query()
	filter := params["filter"][0]
	enc := json.NewEncoder(ctx.ResponseWriter)
	ctx.ResponseWriter.Header().Set("Content-Type", "application/json")

	plugins, err := c.PluginDAO.FindPlugins(filter)
	if err != nil {
		enc.Encode(map[string]string{"error": err.Error()})
	} else {
		enc.Encode(plugins)
	}
	return nil
}

func (c *PluginController) ServePost(ctx *route.Context) error {
	dec := json.NewDecoder(io.TeeReader(ctx.Request.Body, os.Stdout))
	plugin := Plugin{}
	err := dec.Decode(&plugin)
	if err != nil {
		return err
	}
	token := auth.GetToken(ctx)
	plugin.UserId = token.User.Id
	err = c.PluginDAO.Create(&plugin)
	if err != nil {
		return err
	}
	ctx.ResponseWriter.Header().Set("Content-Type", "application/json")
	ctx.ResponseWriter.Write([]byte("{}"))
	return nil
}

func (c *PluginEntryController) ServeGet(ctx *route.Context) error {
	components := strings.SplitN(ctx.Request.URL.Path, "/", 4)
	name := components[3]
	enc := json.NewEncoder(ctx.ResponseWriter)
	ctx.ResponseWriter.Header().Set("Content-Type", "application/json")

	plugin, err := c.PluginDAO.FindPlugin(name)
	if err != nil {
		ctx.ResponseWriter.WriteHeader(http.StatusNotFound)
		enc.Encode(map[string]string{"error": err.Error()})
	} else {
		enc.Encode(plugin)
	}
	return nil
}
