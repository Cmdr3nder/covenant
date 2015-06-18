package service

import (
	"net/http"

	"github.com/zenazn/goji"
	"github.com/zenazn/goji/web"

	"github.com/ender4021/covenant/model"
)

// Server is the interface for an http server based on our Context interface
type Server interface {
	Get(interface{}, func(model.Context, http.ResponseWriter, *http.Request) error)
	Serve()
}

type gojiServer struct{}

func (s *gojiServer) Get(pattern interface{}, fn func(model.Context, http.ResponseWriter, *http.Request) error) {
	goji.Get(pattern, func(c web.C, w http.ResponseWriter, r *http.Request) {
		err := fn(model.GetContext(c), w, r)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})
}

func (s *gojiServer) Serve() {
	goji.Serve()
}
