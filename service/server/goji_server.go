package server

import (
	"net/http"

	"github.com/ender4021/covenant/model"
	"github.com/zenazn/goji"
	"github.com/zenazn/goji/web"
)

type gojiServer struct{}

// NewGoji creates a new Goji based Server instance
func NewGoji() Server {
	return &gojiServer{}
}

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
