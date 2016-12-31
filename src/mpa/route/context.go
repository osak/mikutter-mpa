package route

import (
	"net/http"
)

type Context struct {
	ResponseWriter http.ResponseWriter
	Request        *http.Request
	attributes     map[string]interface{}
}

func NewContext(w http.ResponseWriter, req *http.Request) *Context {
	return &Context{
		ResponseWriter: w,
		Request:        req,
		attributes:     make(map[string]interface{}),
	}
}

func (ctx *Context) PutAttribute(name string, val interface{}) {
	ctx.attributes[name] = val
}

func (ctx *Context) GetAttribute(name string) interface{} {
	return ctx.attributes[name]
}
