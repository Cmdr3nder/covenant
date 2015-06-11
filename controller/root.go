package controller

import (
	"net/http"

	"github.com/ender4021/covenant/model"
	"github.com/ender4021/covenant/service"
	"github.com/ender4021/covenant/service/layout"
	"github.com/ender4021/covenant/service/layout/then"
)

// RegisterRootController add the "/" route and initializes constants for routes controlled by the "Root" controller
func RegisterRootController(server service.Server, config service.Config) {
	t := readTemplates(config)

	server.Get("/", welcomePage(t))
}

func readTemplates(config service.Config) layout.Layout {
	return then.New(service.GetLayout("views_root_index"), service.GetLayout("views_shared_layout"))
}

func welcomePage(l layout.Layout) func(model.Context, http.ResponseWriter, *http.Request) {
	return func(c model.Context, w http.ResponseWriter, r *http.Request) {
		page := model.Page{Title: "This Is The Title", Body: "This is the Body and it", Data: "doesn't have to be a string."}
		err := l.Render(w, page)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}
