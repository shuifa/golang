package main

import (
    "fmt"
    "html"
    "log"
    "net/http"
)

func main() {
    fmt.Println("launch server at 8080 port")

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        _, err := fmt.Fprintf(w, "hello %q", html.EscapeString(r.URL.Path))
        if err != nil {
            log.Fatal(err)
        }
    })

    err := http.ListenAndServe(":8080", nil)
    if err != nil {
        log.Fatal(err)
    }

}
