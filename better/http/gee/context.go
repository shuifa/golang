package gee

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type H map[string]interface{}

type Context struct {
	Path       string
	Method     string
	Writer     http.ResponseWriter
	Request    *http.Request
	StatusCode int
	Params     map[string]string
	Handlers   []HandlerFunc
	Index      int
}

func NewContext(w http.ResponseWriter, r *http.Request) *Context {
	return &Context{
		Path:    r.URL.Path,
		Method:  r.Method,
		Writer:  w,
		Request: r,
		Params:  make(map[string]string),
		Index:   -1,
	}
}

func (ctx *Context) PostForm(key string) string {
	return ctx.Request.FormValue(key)
}

func (ctx *Context) Query(key string) string {
	return ctx.Request.URL.Query().Get(key)
}

func (ctx *Context) Status(code int) {
	ctx.StatusCode = code
	ctx.Writer.WriteHeader(code)
}

func (ctx *Context) SetHeader(key string, value string) {
	ctx.Writer.Header().Set(key, value)
}

func (ctx *Context) String(code int, format string, values ...interface{}) {
	ctx.Writer.Header().Set("Content-type", "text/plain")
	ctx.StatusCode = code
	ctx.Writer.Write([]byte(fmt.Sprintf(format, values...)))
}

func (ctx *Context) JSON(code int, object interface{}) {
	ctx.Writer.WriteHeader(code)
	ctx.SetHeader("Content-Type", "application/json")
	ctx.Status(code)
	encoder := json.NewEncoder(ctx.Writer)
	if err := encoder.Encode(object); err != nil {
		http.Error(ctx.Writer, err.Error(), http.StatusInternalServerError)
	}
}

func (ctx *Context) Data(code int, data []byte) {
	ctx.Status(code)
	ctx.Writer.Write(data)
}

func (ctx *Context) HTML(code int, html string) {
	ctx.SetHeader("Content-Type", "text/html")
	ctx.Status(code)
	ctx.Writer.Write([]byte(html))
}

func (ctx *Context) Param(key string) string {
	value, _ := ctx.Params[key]
	return value
}

func (ctx *Context) Next() {
	ctx.Index++
	for ; ctx.Index < len(ctx.Handlers); ctx.Index++ {
		ctx.Handlers[ctx.Index](ctx)
	}
}

func (ctx *Context) Fail(code int, message string) {
	ctx.String(code, "only for v2 route %s", message)
	return
}
