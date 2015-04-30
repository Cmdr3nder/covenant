package controller

import (
    "net/http"
    "fmt"

    "github.com/ender4021/covenant/model"
    "github.com/ender4021/covenant/service"
)

var g = service.CovenantServer_Goji{}

func RegisterHelloController() {
    g.Get("/hello/:name", hello)
}

func hello(c model.CovenantContext, w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, %s!", c.GetUrlParam("name"))
}
