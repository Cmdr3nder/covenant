package controller

import (
	"net/http"

	"github.com/ender4021/covenant/model"
	"github.com/ender4021/covenant/model/page"
	"github.com/ender4021/covenant/service"
	"github.com/ender4021/covenant/service/config"
	"github.com/ender4021/covenant/service/layout"
	"github.com/ender4021/covenant/service/layout/then"
	"github.com/ender4021/covenant/service/server"
)

// RegisterRootController add the "/" route and initializes constants for routes controlled by the "Root" controller
func RegisterRootController(server server.Server, config config.Config) {
	server.Get("/", welcomePage)
}

func getRootLayout() (layout.Layout, error) {
	rootLayout, err := service.GetLayout("views_root_index")

	if err != nil {
		return nil, err
	}

	sharedLayout, err := service.GetRootLayout()

	if err != nil {
		return nil, err
	}

	return then.New(rootLayout, sharedLayout), nil
}

func welcomePage(c model.Context, w http.ResponseWriter, r *http.Request) error {
	l, err := getRootLayout()

	if err != nil {
		return err
	}

	page := page.Page{Title: "Andrew Bowers"}

	return l.Render(w, page)
}
