package handler

import (
	"encoding/json"
	"model"
	"net/http"
	"strings"
)

type pluginHandler struct {
	dao model.PluginDAO
}

type pluginSearchHandler struct {
	dao model.PluginDAO
}

func NewPluginHandler(dao model.PluginDAO) http.Handler {
	return pluginHandler{dao}
}

func (h pluginHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	components := strings.SplitN(r.URL.Path, "/", 4)
	name := components[3]
	enc := json.NewEncoder(w)
	w.Header().Set("Content-Type", "application/json")

	plugin, err := h.dao.FindPlugin(name)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		enc.Encode(map[string]string{"error": err.Error()})
	} else {
		enc.Encode(plugin)
	}
}

func NewPluginSearchHandler(dao model.PluginDAO) http.Handler {
	return pluginSearchHandler{dao}
}

func (h pluginSearchHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	filter := params["filter"][0]
	enc := json.NewEncoder(w)
	w.Header().Set("Content-Type", "application/json")

	plugins, err := h.dao.FindPlugins(filter)
	if err != nil {
		enc.Encode(map[string]string{"error": err.Error()})
	} else {
		enc.Encode(plugins)
	}
}
