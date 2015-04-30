package model

import (
    "github.com/zenazn/goji/web"
)

type CovenantContext interface {
    GetUrlParam(string) string
    SetUrlParam(string, string)
    GetEnv(interface{}) interface{}
    SetEnv(interface{}, interface{})
}

type CovenantContext_Goji struct {
    GojiContext web.C
}

func (c *CovenantContext_Goji) GetUrlParam(key string) string {
    return c.GojiContext.URLParams[key]
}

func (c *CovenantContext_Goji) SetUrlParam(key string, value string) {
    c.GojiContext.URLParams[key] = value
}

func (c *CovenantContext_Goji) GetEnv(key interface{}) interface{} {
    return c.GojiContext.Env[key]
}

func (c *CovenantContext_Goji) SetEnv(key interface{}, value interface{}) {
    c.GojiContext.Env[key] = value
}
