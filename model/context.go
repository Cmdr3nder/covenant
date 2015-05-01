package model

import (
    "github.com/zenazn/goji/web"
)

type Context interface {
    GetUrlParam(string) string
    SetUrlParam(string, string)
    GetEnv(interface{}) interface{}
    SetEnv(interface{}, interface{})
}

type gojiContext struct {
    context web.C
}

func (c *gojiContext) GetUrlParam(key string) string {
    return c.context.URLParams[key]
}

func (c *gojiContext) SetUrlParam(key string, value string) {
    c.context.URLParams[key] = value
}

func (c *gojiContext) GetEnv(key interface{}) interface{} {
    return c.context.Env[key]
}

func (c *gojiContext) SetEnv(key interface{}, value interface{}) {
    c.context.Env[key] = value
}
