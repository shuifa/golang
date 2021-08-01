package gee

import (
    "net/http"
)

// HandlerFunc defines the request handler used by gee
type HandlerFunc func(ctx *Context)

type Engine struct {
    router *router
}

func New() *Engine {
    return &Engine{router: NewRoute()}
}

func (e *Engine) addRoute(method string, path string, handler HandlerFunc) {
    key := method + "-" + path
    e.router.handlers[key] = handler
}

func (e *Engine) Get(path string, handler HandlerFunc) {
    e.addRoute("GET", path, handler)
}

func (e *Engine) Post(path string, handler HandlerFunc) {
    e.addRoute("POST", path, handler)
}

func (e *Engine) Run(addr string) error {
    return http.ListenAndServe(addr, e)
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    ctx := NewContext(w, r)
    e.router.Handler(ctx)
}
