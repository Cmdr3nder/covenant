package controller

import (
    "net/http"
    "fmt"

    "github.com/ender4021/covenant/model"
    "github.com/ender4021/covenant/service"
)

func RegisterBlogController(server service.Server) {
    path := service.GetRouteBuilder()

    path.AppendPart("blog")
    server.Get(path.MustCompile(), getBlogRoot)

    path.AppendPart("(?P<year>(19|20)[0-9]{2})")
    server.Get(path.MustCompile(), getBlogYear)

    path.AppendPart("(?P<month>(0[1-9])|(1[1-2]))")
    server.Get(path.MustCompile(), getBlogMonth)

    path.AppendPart("(?P<guid>([0-9]|[a-z]|[A-Z]|-|_)+)")
    server.Get(path.MustCompile(), getBlogPost)
}

func getBlogRoot(c model.Context, w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Blog Root")
}

func getBlogYear(c model.Context, w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Blog Year: %s", c.GetUrlParam("year"))
}

func getBlogMonth(c model.Context, w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Blog Month: %s %s", c.GetUrlParam("year"), c.GetUrlParam("month"))
}

func getBlogPost(c model.Context, w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Blog Post: %s %s %s", c.GetUrlParam("year"), c.GetUrlParam("month"), c.GetUrlParam("guid"))
}
