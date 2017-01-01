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

func (fc *FilterChain) WrapGet(controller GetController) GetController {
	return &internalGetHandler{
		filterChain: fc,
		controller:  controller,
	}
}

func (fc *FilterChain) WrapPost(controller PostController) PostController {
	return &internalPostHandler{
		filterChain: fc,
		controller:  controller,
	}
}

type internalGetHandler struct {
	filterChain *FilterChain
	controller  GetController
}

type internalPostHandler struct {
	filterChain *FilterChain
	controller  PostController
}

func (handler *internalGetHandler) ServeGet(ctx *Context) error {
	return processFilter(handler.filterChain, ctx, func(ctx *Context) error {
		return handler.controller.ServeGet(ctx)
	})
}

func (handler *internalPostHandler) ServePost(ctx *Context) error {
	return processFilter(handler.filterChain, ctx, func(ctx *Context) error {
		return handler.controller.ServePost(ctx)
	})
}

func processFilter(filterChain *FilterChain, ctx *Context, callback func(*Context) error) error {
	lastProcessed := -1
	lastCompleted := -1
	for i, f := range *filterChain {
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
		f := (*filterChain)[i]
		f.PostHandle(ctx)
	}
	return nil
}
