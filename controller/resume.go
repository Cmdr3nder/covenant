package controller

import (
    "net/http"
    "fmt"

    "github.com/ender4021/covenant/model"
    "github.com/ender4021/covenant/service"
)

func RegisterResumeController(server service.Server) {
    path := service.GetRouteBuilder()

    path.AppendPart("resume")
    server.Get(path.MustCompile(), getResumeRoot)

    detail := path.Fork()
    detail.AppendPart("detail")
    server.Get(detail.MustCompile(), getResumeDetails)

    project := path.Fork()
    project.AppendPart("project")
    server.Get(project.MustCompile(), getResumeProjects)
}

func getResumeRoot(c model.Context, w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Resume Root")
}

func getResumeDetails(c model.Context, w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Resume Details")
}

func getResumeProjects(c model.Context, w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Resume Projects")
}
