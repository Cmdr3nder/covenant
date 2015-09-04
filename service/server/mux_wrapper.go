package server

import (
	"net/http"

	"github.com/ender4021/covenant/model"
	"github.com/zenazn/goji/web"
)

// MuxWrapper is a wrapper for web.Mux
type MuxWrapper struct {
	Mux *web.Mux
}

// Get wraps web.Mux.Get
func (s *MuxWrapper) Get(pattern interface{}, fn func(model.Context, http.ResponseWriter, *http.Request) error) {
	s.Mux.Get(pattern, func(c web.C, w http.ResponseWriter, r *http.Request) {
		err := fn(model.GetContext(c), w, r)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})
}

// Post wraps web.Mux.Post
func (s *MuxWrapper) Post(pattern interface{}, fn func(model.Context, http.ResponseWriter, *http.Request) error) {
	s.Mux.Post(pattern, func(c web.C, w http.ResponseWriter, r *http.Request) {
		err := fn(model.GetContext(c), w, r)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})
}

// Use wraps web.Mux.Use
func (s *MuxWrapper) Use(middleware interface{}) {
	s.Mux.Use(middleware)
}
