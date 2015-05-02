package controller

import (
    "bytes"
    "regexp"
    "net/http"
    "fmt"

    "github.com/ender4021/covenant/model"
    "github.com/ender4021/covenant/service"
)

func RegisterBlogController(server service.Server) {
    path := bytes.Buffer{}
    end := "/?$"

    path.WriteString("^/blog")
    server.Get(regexp.MustCompile(path.String() + end), getBlogRoot)

    path.WriteString("/(?P<year>(19|20)[0-9]{2})")
    server.Get(regexp.MustCompile(path.String() + end), getBlogYear)

    path.WriteString("/(?P<month>(0[1-9])|(1[1-2]))")
    server.Get(regexp.MustCompile(path.String() + end), getBlogMonth)

    path.WriteString("/(?P<guid>([0-9]|[a-z]|[A-Z]|-|_)+)")
    server.Get(regexp.MustCompile(path.String() + end), getBlogPost)
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
