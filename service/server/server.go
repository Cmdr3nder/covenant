package server

import (
	"net/http"

	"github.com/ender4021/covenant/model"
)

// Server is the interface for an http server based on our Context interface
type Server interface {
	Get(interface{}, func(model.Context, http.ResponseWriter, *http.Request) error)
	Serve()
}
