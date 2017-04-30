package plugin

import (
	"errors"
	"fmt"
	"github.com/satori/go.uuid"
	"io"
	"io/ioutil"
	"mpa/filter"
	"mpa/model"
	"mpa/route"
	view "mpa/view/plugin"
	"os"
	"path/filepath"
	"strings"
)

const StoragePath = "/app/storage"

type PluginController struct {
	PluginDAO model.PluginDAO
}

func (c *PluginController) ServeGet(ctx *route.Context) (route.View, error) {
	params := ctx.Request.URL.Query()
	filter := params["filter"][0]
	ctx.ResponseWriter.Header().Set("Content-Type", "application/json")

	plugins, err := c.PluginDAO.FindByKeyword(filter)
	if err != nil {
		return nil, err
	}
	return &view.SearchView{
		Plugins: plugins,
	}, nil
}

func (c *PluginController) ServePost(ctx *route.Context) (route.View, error) {
	r, err := ctx.Request.MultipartReader()
	if err != nil {
		return nil, err
	}

	for {
		p, err := r.NextPart()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		if p.FormName() == "plugin-archive" {
			data, err := ioutil.ReadAll(p)
			if err != nil {
				return nil, err
			}
			f, err := saveFile(data)
			if err != nil {
				return nil, err
			}
			defer os.Remove(f.Name())

			spec, err := LoadSpec(f.Name())
			if err != nil {
				return nil, err
			}
			authResult := filter.GetAuthResult(ctx)
			if p, err := c.PluginDAO.FindBySlug(spec.Slug); err == nil {
				if !model.SameUser(p.Author, authResult.User) {
					return nil, fmt.Errorf("Plugin already exists but is owned by other user %s", p.Author.Login)
				}
			}
			if _, err := c.PluginDAO.FindBySlugAndVersion(spec.Slug, spec.Version); err == nil {
				return nil, fmt.Errorf("Plugin of the same version and slug already exists")
			}
			uuid := uuid.NewV4()
			plugin := model.Plugin{
				Name:        spec.Name,
				Version:     spec.Version,
				Description: spec.Description,
				Uuid:        uuid,
				Slug:        spec.Slug,
			}
			plugin.Author = authResult.User
			err = c.PluginDAO.Create(&plugin)
			if err != nil {
				return nil, err
			}
			if err = os.Link(f.Name(), filepath.Join(StoragePath, uuid.String())); err != nil {
				return nil, err
			}

			return &view.EntryView{
				Plugin: plugin,
			}, nil
		}
	}
	return nil, errors.New("mpa/plugin: plugin-archive entity does not found")
}

func saveFile(data []byte) (*os.File, error) {
	f, err := ioutil.TempFile(StoragePath, "mpa-plugin-temp")
	if err != nil {
		return nil, err
	}
	defer f.Close()

	_, err = f.Write(data)
	return f, err
}

type PluginEntryController struct {
	PluginDAO model.PluginDAO
	UserDAO   model.UserDAO
}

// ServeGet implements GetController
func (c *PluginEntryController) ServeGet(ctx *route.Context) (route.View, error) {
	components := strings.SplitN(ctx.Request.URL.Path, "/", 4)
	slug := components[3]
	ctx.ResponseWriter.Header().Set("Content-Type", "application/json")

	plugin, err := c.PluginDAO.FindBySlug(slug)
	if err != nil {
		return nil, err
	}
	c.UserDAO.Fill(&plugin.Author)
	return &view.EntryView{
		Plugin: plugin,
	}, nil
}
