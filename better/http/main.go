package main

import (
    "fmt"

    "github.com/oushuifa/golang/better/http/gee"
)

func main() {
    engine := gee.New()
    engine.Get("/", func(ctx *gee.Context) {
        fmt.Fprintf(ctx.Writer, "r.path, %q \n", ctx.Path)
    })
    engine.Post("/hello", func(ctx *gee.Context) {
        for k, v := range ctx.Request.Header {
            fmt.Fprintf(ctx.Writer, "header[%q], vaule: %q", k, v)
        }
    })
    engine.Run(":9999")
}
