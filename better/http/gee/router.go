package gee

import (
	"net/http"
	"strings"
)

type router struct {
	roots    map[string]*node
	handlers map[string]HandlerFunc
}

func NewRoute() *router {
	return &router{
		handlers: make(map[string]HandlerFunc),
		roots:    make(map[string]*node),
	}
}

func (r *router) addRoute(method string, path string, handler HandlerFunc) {

	parts := parsePartner(path)

	key := method + "-" + path

	_, ok := r.roots[method]

	if !ok {
		r.roots[method] = &node{}
	}

	r.roots[method].Insert(path, parts, 0)

	r.handlers[key] = handler
}

func (r *router) Handler(ctx *Context) {

	n, params := r.getRoute(ctx.Method, ctx.Path)

	if n == nil {

		ctx.Handlers = append(ctx.Handlers, func(ctx *Context) {
			ctx.String(http.StatusNotFound, "404 not found %s", ctx.Path)
		})

	} else {

		ctx.Params = params

		key := ctx.Method + "-" + n.path

		ctx.Handlers = append(ctx.Handlers, r.handlers[key])
	}

	ctx.Next()
}

func parsePartner(path string) []string {
	vs := strings.Split(path, "/")

	parts := make([]string, 0)
	for _, item := range vs {
		if item != "" {
			parts = append(parts, item)
			if item[0] == '*' {
				break
			}
		}
	}
	return parts
}

func (r *router) getRoute(method string, path string) (*node, map[string]string) {
	searchParts := parsePartner(path)
	params := make(map[string]string)
	root, ok := r.roots[method]

	if !ok {
		return nil, nil
	}

	n := root.Search(searchParts, 0)

	if n != nil {
		parts := parsePartner(n.path)
		for index, part := range parts {
			if part[0] == ':' {
				params[part[1:]] = searchParts[index]
			}
			if part[0] == '*' && len(part) > 1 {
				params[part[1:]] = strings.Join(searchParts[index:], "/")
				break
			}
		}
		return n, params
	}

	return nil, nil
}
