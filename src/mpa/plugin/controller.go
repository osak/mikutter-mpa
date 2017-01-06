package plugin

import (
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"mpa/auth"
	"mpa/route"
	"net/http"
	"os"
	"strings"
)

type PluginController struct {
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
	r, err := ctx.Request.MultipartReader()
	if err != nil {
		return err
	}

	for {
		p, err := r.NextPart()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		if p.FormName() == "plugin-archive" {
			data, err := ioutil.ReadAll(p)
			if err != nil {
				return err
			}
			f, err := saveFile(data)
			if err != nil {
				return err
			}

			spec, err := LoadSpec(f.Name())
			if err != nil {
				return err
			}
			plugin := Plugin{
				Name:        spec.Name,
				Version:     spec.Version,
				Description: spec.Description,
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
	}
	return errors.New("mpa/plugin: plugin-archive entity does not found")
}

func saveFile(data []byte) (*os.File, error) {
	f, err := ioutil.TempFile("", "mpa-plugin")
	if err != nil {
		return nil, err
	}
	defer f.Close()

	_, err = f.Write(data)
	return f, err
}

type PluginEntryController struct {
	PluginDAO PluginDAO
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
