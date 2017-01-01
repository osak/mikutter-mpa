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
	http.HandleFunc(path, generateHandler(controller))
}

func generateHandler(controller Controller) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
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
}
