package gee

import (
	"log"
	"net/http"
	"strings"
)

// HandlerFunc defines the request handler used by gee
type HandlerFunc func(ctx *Context)

type Engine struct {
	*RouterGroup
	router *router
	groups []*RouterGroup
}

type RouterGroup struct {
	Middlewares []HandlerFunc
	Prefix      string
	Engine      *Engine
	parent      *RouterGroup
}

func New() *Engine {

	engine := &Engine{router: NewRoute()}

	engine.RouterGroup = &RouterGroup{Engine: engine}

	engine.groups = []*RouterGroup{engine.RouterGroup}

	return engine
}

func (e *Engine) addRoute(method string, path string, handler HandlerFunc) {
	// key := method + "-" + path
	// e.router.handlers[key] = handler
	e.router.addRoute(method, path, handler)
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

	var middlewares []HandlerFunc

	for _, group := range e.groups {
		if strings.HasPrefix(r.URL.Path, group.Prefix) {
			middlewares = append(middlewares, group.Middlewares...)
		}
	}

	ctx := NewContext(w, r)
	ctx.Handlers = middlewares
	e.router.Handler(ctx)
}

func (g *RouterGroup) Group(prefix string) *RouterGroup {

	engine := g.Engine

	newGroup := &RouterGroup{
		Middlewares: nil,
		Prefix:      g.Prefix + prefix,
		Engine:      engine,
		parent:      g,
	}

	engine.groups = append(engine.groups, newGroup)

	return newGroup
}

func (g *RouterGroup) addRoute(method string, comp string, handler HandlerFunc) {
	pattern := g.Prefix + comp
	log.Printf("Route %4s - %s", method, pattern)
	g.Engine.addRoute(method, pattern, handler)
}

// GET defines the method to add GET request
func (g *RouterGroup) GET(pattern string, handler HandlerFunc) {
	g.addRoute("GET", pattern, handler)
}

// POST defines the method to add POST request
func (g *RouterGroup) POST(pattern string, handler HandlerFunc) {
	g.addRoute("POST", pattern, handler)
}

func (g *RouterGroup) Use(middleware ...HandlerFunc) {
	g.Middlewares = append(g.Middlewares, middleware...)
}

