package controller

import (
    "net/http"
    "fmt"

    "github.com/ender4021/covenant/model"
    "github.com/ender4021/covenant/service"
)

func RegisterHelloController(server service.Server) {
    server.Get("/hello/:name", hello)
}

func hello(c model.Context, w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, %s!", c.GetUrlParam("name"))
}
