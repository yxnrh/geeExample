package router

import (
	"geeExample/gee/context"
	"net/http"
)

type HandleFunc func(c *context.Context)

type router struct {
	handler map[string]HandleFunc
}

func newRouter()*router {
	return &router{handler: make(map[string]HandleFunc)}
}

func (router *router)addRoute(method , pattern string,handler HandleFunc)  {
	key := method+"-"+pattern
	router.handler[key] = handler
}

func (router *router)handle(c *context.Context)  {
	key := c.Method+"-"+c.Path
	if handle,ok := router.handler[key];ok{
		handle(c)
	}else {
		c.String(http.StatusNotFound,"404 NOT FOUND: %s\n",c.Path)
	}
}
