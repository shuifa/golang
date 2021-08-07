package gee

import (
    "log"
    "time"
)

func Logger() HandlerFunc {

    return func(ctx *Context) {

        t := time.Now()

        time.Sleep(time.Second)
        log.Printf("%#v", ctx.Request)

        ctx.Next()

        log.Printf("[%d] %s %v", ctx.StatusCode, ctx.Request.URL, time.Since(t))
    }

}
