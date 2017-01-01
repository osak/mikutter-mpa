package route

import (
	"fmt"
	"net/http"
)

type Router struct {
	controllers map[string]Controller
}

func NewRouter() *Router {
	return &Router{
		make(map[string]Controller),
	}
}

func (router *Router) Register(path string, controller Controller) {
	router.controllers[path] = controller
	http.Handle(path, router)
}

func (router *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	controller := router.controllers[req.URL.Path]
	ctx := NewContext(w, req)
	var err error
	switch req.Method {
	case "GET":
		err = controller.ServeGet(ctx)
	case "POST":
		err = controller.ServePost(ctx)
	}
	if err != nil {
		fmt.Printf("Error during request processing: %v\n", err.Error())
	}
}
