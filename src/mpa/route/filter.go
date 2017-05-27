package route

import ()

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

func (fc *FilterChain) WrapDelete(controller DeleteController) DeleteController {
	return &internalDeleteHandler{
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

type internalDeleteHandler struct {
	filterChain *FilterChain
	controller  DeleteController
}

// ServeGet implements GetController
func (handler *internalGetHandler) ServeGet(ctx *Context) (View, error) {
	return processFilter(handler.filterChain, ctx, func(ctx *Context) (View, error) {
		return handler.controller.ServeGet(ctx)
	})
}

// ServePost implements PostController
func (handler *internalPostHandler) ServePost(ctx *Context) (View, error) {
	return processFilter(handler.filterChain, ctx, func(ctx *Context) (View, error) {
		return handler.controller.ServePost(ctx)
	})
}

func (handler *internalDeleteHandler) ServeDelete(ctx *Context) (View, error) {
	return processFilter(handler.filterChain, ctx, func(ctx *Context) (View, error) {
		return handler.controller.ServeDelete(ctx)
	})
}

func processFilter(filterChain *FilterChain, ctx *Context, callback func(*Context) (View, error)) (View, error) {
	for _, f := range *filterChain {
		err := f.PreHandle(ctx)
		defer f.PostHandle(ctx)
		if err != nil {
			return nil, err
		}
	}
	return callback(ctx)
}
