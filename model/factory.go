package model

import "github.com/zenazn/goji/web"

// GetContext constructs a new Covenant Context for the given Goji Context
func GetContext(c web.C) Context {
	return &gojiContext{context: c}
}
