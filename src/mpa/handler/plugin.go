package handler

import (
	"fmt"
	"net/http"
)

type handler struct {
	dao *PluginDAO
}

func NewHandler(dao *PluginDAO) Handler {
	return handler{dao}
}

func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "{}")
}
