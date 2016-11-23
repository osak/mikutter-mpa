package handler

import (
	"encoding/json"
	"model"
	"net/http"
)

type pluginHandler struct {
	dao model.PluginDAO
}

func NewPluginHandler(dao model.PluginDAO) http.Handler {
	return pluginHandler{dao}
}

func (h pluginHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	name := params["name"][0]
	enc := json.NewEncoder(w)
	w.Header().Set("Content-Type", "application/json")

	plugin, err := h.dao.FindPlugin(name)
	if err != nil {
		enc.Encode(map[string]string{"error": err.Error()})
	} else {
		enc.Encode(plugin)
	}
}
