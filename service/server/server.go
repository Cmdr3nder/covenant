package server

import (
	"net/http"

	"github.com/ender4021/covenant/model"
	"github.com/zenazn/goji/web"
)

// Server is the interface for an http server based on our Context interface
type Server interface {
	Get(interface{}, func(model.Context, http.ResponseWriter, *http.Request) error)
	Handle(interface{}, MuxWrapper)
	Serve()
}

// NewMux returns a wrapper for web.New()
func NewMux() MuxWrapper {
	return MuxWrapper{Mux: web.New()}
}
