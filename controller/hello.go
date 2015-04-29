package controller

import (
    "net/http"
    "fmt"

    "github.com/zenazn/goji"
    "github.com/zenazn/goji/web"
)

func RegisterHelloController() {
    goji.Get("/hello/:name", hello)
}

func hello(c web.C, w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, %s!", c.URLParams["name"])
}
