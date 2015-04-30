package service

import (
    "net/http"

    "github.com/zenazn/goji"
    "github.com/zenazn/goji/web"

    "github.com/ender4021/covenant/model"
)

type CovenantServer interface {
    Get(interface{}, func(model.CovenantContext, http.ResponseWriter, *http.Request))
}

type CovenantServer_Goji struct {

}

func (s *CovenantServer_Goji) Get(pattern interface{}, fn func(model.CovenantContext, http.ResponseWriter, *http.Request)) {
    goji.Get(pattern, func (c web.C, w http.ResponseWriter, r *http.Request) {
        fn(&model.CovenantContext_Goji{GojiContext: c}, w, r)
    })
}
