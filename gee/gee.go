package gee

import (
	"fmt"
	"net/http"
)

type HandleFunc func(http.ResponseWriter,*http.Request)

type Engine struct {
	router map[string]HandleFunc
}

func New()*Engine {
	return &Engine{router: make(map[string]HandleFunc)}
}

func (engine *Engine)addRoute(method , pattern string,handler HandleFunc)  {
	key := method+"-"+pattern
	engine.router[key] = handler
}


func (engine *Engine)GET(pattern string,handler HandleFunc)  {
	engine.addRoute("GET",pattern,handler)
}

func (engine *Engine)POST(pattern string,handler HandleFunc)  {
	engine.addRoute("POST",pattern,handler)
}

func (engine *Engine)Run(addr string)(err error) {
	return http.ListenAndServe(addr,engine)
}

func (engine *Engine)ServeHTTP(w http.ResponseWriter,r *http.Request)  {
	key := r.Method+"-"+r.URL.Path
	if handle,ok := engine.router[key];ok{
		handle(w,r)
	}else {
		fmt.Fprintf(w,"404 NOT FOUND %s\n",r.URL.Path)
	}
}
