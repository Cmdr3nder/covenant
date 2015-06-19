package page

import (
	"html/template"
)

// Page is a struct that represents the content of a single page.
type Page struct {
	Title       template.HTMLAttr
	Body        template.HTML
	Data        interface{}
	StyleSheets []template.HTMLAttr
	Scripts     []template.HTMLAttr
}

// ReduceLinkedFiles eliminates duplicate entries in the StyleSheets and Scripts slices
func (p *Page) ReduceLinkedFiles() {
	p.StyleSheets = reduce(p.StyleSheets)
	p.Scripts = reduce(p.Scripts)
}

func reduce(items []template.HTMLAttr) []template.HTMLAttr {
	nItemsMap := make(map[template.HTMLAttr]bool)
	var nItemsSlice []template.HTMLAttr

	for _, item := range items {
		if !nItemsMap[item] {
			nItemsMap[item] = true
			nItemsSlice = append(nItemsSlice, item)
		}
	}

	return nItemsSlice
}
