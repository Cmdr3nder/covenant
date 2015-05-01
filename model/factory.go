package model

import (
    "github.com/zenazn/goji/web"
)

func GetContext(c web.C) Context {
    return &gojiContext{context: c}
}
