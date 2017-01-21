package route

import (
	"fmt"
	"net/http"
)

type Router struct {
	getControllers  map[string]GetController
	postControllers map[string]PostController
	handlerFuncs    map[string]func(http.ResponseWriter, *http.Request)
}

func NewRouter() *Router {
	return &Router{
		getControllers:  make(map[string]GetController),
		postControllers: make(map[string]PostController),
		handlerFuncs:    make(map[string]func(http.ResponseWriter, *http.Request)),
	}
}

func (router *Router) RegisterGet(path string, controller GetController) {
	router.getControllers[path] = controller
	router.registerHandler(path)
}

func (router *Router) RegisterPost(path string, controller PostController) {
	router.postControllers[path] = controller
	router.registerHandler(path)
}

func (router *Router) registerHandler(path string) {
	if _, ok := router.handlerFuncs[path]; !ok {
		handler := generateHandler(path, router)
		router.handlerFuncs[path] = handler
		http.HandleFunc(path, handler)
	}
}

func generateHandler(path string, router *Router) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		ctx := NewContext(w, req)
		var view View
		var err error
		switch req.Method {
		case "GET":
			if ctl, ok := router.getControllers[path]; ok {
				view, err = ctl.ServeGet(ctx)
			}
		case "POST":
			if ctl, ok := router.postControllers[path]; ok {
				view, err = ctl.ServePost(ctx)
			}
		}
		if err != nil {
			fmt.Printf("Error during request processing: %v\n", err.Error())
			return
		}
		err = view.Render(ctx)
		if err != nil {
			fmt.Printf("Error during rendering: %v\n", err.Error())
		}
	}
}
