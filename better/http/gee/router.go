package gee

import (
    "log"
    "net/http"
)

type router struct {
    handlers map[string]HandlerFunc
}

func NewRoute() *router {
    return & router{handlers : make(map[string] HandlerFunc)}
}

func (r *router) addRoute(method string, path string, handler HandlerFunc) {
    log.Printf("method is %s, router is %s \n", method, path)
    key := method + "-" + path
    r.handlers[key] = handler
}

func(r *router) Handler(ctx *Context) {
    key := ctx.Method + "-" + ctx.Path
    if handler, ok := r.handlers[key]; ok {
        handler(ctx)
    } else {
        ctx.String(http.StatusNotFound, "404 not found %s", ctx.Path)
    }
}