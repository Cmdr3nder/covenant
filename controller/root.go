package controller

import (
	"net/http"

	"github.com/ender4021/covenant/model"
	"github.com/ender4021/covenant/service"
	"github.com/ender4021/covenant/service/config"
	"github.com/ender4021/covenant/service/layout/then"
)

// RegisterRootController add the "/" route and initializes constants for routes controlled by the "Root" controller
func RegisterRootController(server service.Server, config config.Config) {
	server.Get("/", welcomePage)
}

func welcomePage(c model.Context, w http.ResponseWriter, r *http.Request) error {
	l := then.New(service.GetLayout("views_root_index"), service.GetLayout("views_shared_layout"))

	page := model.Page{Title: "Andrew Bowers"}

	return l.Render(w, page)
}
