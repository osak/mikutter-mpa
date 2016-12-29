package route

import (
	"net/http"
)

type Filter interface {
	PreHandle(w http.ResponseWriter, req *http.Request) bool
	PostHandle(w http.ResponseWriter, req *http.Request) bool
}

type FilterChain []Filter

func CreateFilterChain(filters ...Filter) FilterChain {
	return filters
}

func (fc FilterChain) Wrap(handler func(http.ResponseWriter, *http.Request)) http.Handler {
	return &internalHandler{
		filterChain: fc,
		rawHandler:  handler,
	}
}

type internalHandler struct {
	filterChain FilterChain
	rawHandler  func(http.ResponseWriter, *http.Request)
}

func (handler *internalHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	lastProcessed := 0
	lastCompleted := 0
	for i, f := range handler.filterChain {
		lastProcessed = i
		if !f.PreHandle(w, req) {
			break
		}
		lastCompleted = i
	}
	if lastCompleted == lastProcessed {
		handler.rawHandler(w, req)
	}
	for i := lastProcessed; i >= 0; i-- {
		f := handler.filterChain[i]
		f.PostHandle(w, req)
	}
}
