package route

import (
	"net/http"
)

type Filter interface {
	PreHandle(ctx *Context) error
	PostHandle(ctx *Context) error
}

type FilterChain []Filter

func CreateFilterChain(filters ...Filter) *FilterChain {
	var chain FilterChain = filters
	return &chain
}

func (fc *FilterChain) Wrap(controller Controller) http.Handler {
	return &internalHandler{
		filterChain: fc,
		controller:  controller,
	}
}

type internalHandler struct {
	filterChain *FilterChain
	controller  Controller
}

func (handler *internalHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	lastProcessed := -1
	lastCompleted := -1
	ctx := NewContext(w, req)
	for i, f := range *handler.filterChain {
		lastProcessed = i
		if err := f.PreHandle(ctx); err != nil {
			break
		}
		lastCompleted = i
	}
	if lastCompleted == lastProcessed {
		handler.controller.Serve(ctx)
	}
	for i := lastProcessed; i >= 0; i-- {
		f := (*handler.filterChain)[i]
		f.PostHandle(ctx)
	}
}
