package controller

import (
	"fmt"
	"net/http"

	"github.com/ender4021/covenant/model"
	"github.com/ender4021/covenant/service"
)

// RegisterWorkController adds routes and initializes constants for routes controlled by the "Work" controller
func RegisterWorkController(server service.Server) {
	path := service.GetRouteBuilder()

	path.AppendPart("work")
	server.Get(path.MustCompile(), getWorkRoot)

	detail := path.Fork()
	detail.AppendPart("detail")
	server.Get(detail.MustCompile(), getWorkDetails)

	project := path.Fork()
	project.AppendPart("project")
	server.Get(project.MustCompile(), getWorkProjects)
}

func getWorkRoot(c model.Context, w http.ResponseWriter, r *http.Request) error {
	fmt.Fprintf(w, "Resume Root")
	return nil
}

func getWorkDetails(c model.Context, w http.ResponseWriter, r *http.Request) error {
	fmt.Fprintf(w, "Resume Details")
	return nil
}

func getWorkProjects(c model.Context, w http.ResponseWriter, r *http.Request) error {
	fmt.Fprintf(w, "Resume Projects")
	return nil
}
