package handler

import (
	"fmt"
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
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "{}")
}
