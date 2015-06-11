package model

import "html/template"

// Page is a struct that represents the content of a single page.
type Page struct {
	Title       template.HTMLAttr
	Body        template.HTML
	Data        interface{}
	StyleSheets []template.HTMLAttr
	Scripts     []template.HTMLAttr
}
