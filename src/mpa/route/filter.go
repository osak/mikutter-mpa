package route

import (
	"fmt"
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

func (fc *FilterChain) Wrap(controller Controller) Controller {
	return &internalHandler{
		filterChain: fc,
		controller:  controller,
	}
}

type internalHandler struct {
	filterChain *FilterChain
	controller  Controller
}

func (handler *internalHandler) ServeGet(ctx *Context) error {
	return handler.processFilter(ctx, func(ctx *Context) error {
		return handler.controller.ServeGet(ctx)
	})
}

func (handler *internalHandler) ServePost(ctx *Context) error {
	return handler.processFilter(ctx, func(ctx *Context) error {
		return handler.controller.ServePost(ctx)
	})
}

func (handler *internalHandler) processFilter(ctx *Context, callback func(*Context) error) error {
	lastProcessed := -1
	lastCompleted := -1
	for i, f := range *handler.filterChain {
		lastProcessed = i
		if err := f.PreHandle(ctx); err != nil {
			break
		}
		lastCompleted = i
	}
	if lastCompleted == lastProcessed {
		err := callback(ctx)
		if err != nil {
			fmt.Printf("Error: %v\n", err.Error())
		}
	}
	for i := lastProcessed; i >= 0; i-- {
		f := (*handler.filterChain)[i]
		f.PostHandle(ctx)
	}
	return nil
}
