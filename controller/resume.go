package controller

import (
    "net/http"
    "fmt"

    "github.com/ender4021/covenant/model"
    "github.com/ender4021/covenant/service"
)

func RegisterResumeController(server service.Server) {
    server.Get("/resume", getResumeRoot)
    server.Get("/resume/detail", getResumeDetails)
    server.Get("/resume/project", getResumeProjects)
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
