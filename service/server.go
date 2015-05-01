package service

import (
    "net/http"

    "github.com/zenazn/goji"
    "github.com/zenazn/goji/web"

    "github.com/ender4021/covenant/model"
)

type Server interface {
    Get(interface{}, func(model.Context, http.ResponseWriter, *http.Request))
    Serve()
}

type gojiServer struct {

}

func (s *gojiServer) Get(pattern interface{}, fn func(model.Context, http.ResponseWriter, *http.Request)) {
    goji.Get(pattern, func (c web.C, w http.ResponseWriter, r *http.Request) {
        fn(model.GetContext(c), w, r)
    })
}

func (s *gojiServer) Serve() {
    goji.Serve()
}
