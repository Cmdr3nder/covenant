package controller

import (
    "net/http"
    "fmt"

    "github.com/ender4021/covenant/model"
    "github.com/ender4021/covenant/service"
)

func RegisterStaticFileController(server service.Server) {
    cssRegex := service.GetRouteBuilder().AppendPart("css").AppendPart("(?P<fileName>.+\\.css)")
    server.Get(cssRegex.MustCompile(), getCssFile)

    jsRegex := service.GetRouteBuilder().AppendPart("js").AppendPart("(?P<fileName>.+\\.js)")
    server.Get(jsRegex.MustCompile(), getJsFile)

    server.Get("/favicon.ico", getFavicon)
}

func getCssFile(c model.Context, w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Rendering \"css\" %s %s", c.GetUrlParam("fileName"), r.FormValue("kingdom"))
}

func getJsFile(c model.Context, w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Rendering \"js\" %s", c.GetUrlParam("fileName"))
}

func getFavicon(c model.Context, w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Oops Favicon Not Found Lolz")
}
