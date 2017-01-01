package resource

import (
	"mpa/route"
	"net/http"
)

type Resource interface {
	route.GetController
	route.PostController
	Name() string
}

func Register(resource Resource) {
}
