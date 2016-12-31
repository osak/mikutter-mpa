package route

import (
	"net/http"
)

type Controller interface {
	Serve(*Context) error
}

func Wrap(handler func(http.ResponseWriter, *http.Request)) Controller {
	return &internalController{handler}
}

func Unwrap(controller Controller) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		ctx := NewContext(w, req)
		controller.Serve(ctx)
	}
}

type internalController struct {
	handler func(http.ResponseWriter, *http.Request)
}

func (controller *internalController) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	ctx := NewContext(w, req)
	controller.Serve(ctx)
}

func (controller *internalController) Serve(ctx *Context) error {
	controller.handler(ctx.ResponseWriter, ctx.Request)
	return nil
}
