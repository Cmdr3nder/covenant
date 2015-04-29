package controller

import (
    "regexp"
    "net/http"
    "fmt"

    "github.com/zenazn/goji"
    "github.com/zenazn/goji/web"
)

func RegisterStaticFileController() {
    cssRegex := regexp.MustCompile("^/css/(?P<fileName>.+\\.css)$")
    jsRegex := regexp.MustCompile("^/js/(?P<fileName>.+\\.js)$")

    goji.Get(cssRegex, getCssFile)
    goji.Get(jsRegex, getJsFile)
    goji.Get("/favicon.ico", getFavicon)
}

func getCssFile(c web.C, w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Rendering \"css\" %s %s", c.URLParams["fileName"], r.FormValue("kingdom"))
}

func getJsFile(c web.C, w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Rendering \"js\" %s", c.URLParams["fileName"])
}

func getFavicon(c web.C, w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Oops Favicon Not Found Lolz")
}
