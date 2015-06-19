package master

import (
	"bytes"
	"html/template"
	"io"

	"github.com/ender4021/covenant/model"
	"github.com/ender4021/covenant/service/layout"
)

// New returns a freshly constructed masterLayout
func New(compiledLayout *template.Template, styleSheets []template.HTMLAttr, scripts []template.HTMLAttr) layout.Layout {
	return &masterLayout{compiledLayout: compiledLayout, styleSheets: styleSheets, scripts: scripts}
}

type masterLayout struct {
	compiledLayout *template.Template
	styleSheets    []template.HTMLAttr
	scripts        []template.HTMLAttr
}

// Render applies the layout to the given page and writes the result to the given writer
func (l *masterLayout) Render(w io.Writer, p model.Page) error {
	return l.render(w, &p)
}

func (l *masterLayout) render(w io.Writer, p *model.Page) error {
	//Add this layout's scripts and stylesheets
	p.Scripts = append(l.scripts, p.Scripts...)
	p.StyleSheets = append(l.styleSheets, p.StyleSheets...)
	p.ReduceLinkedFiles()

	return l.compiledLayout.Execute(w, p)
}

// RenderStep applies the layout to the given page and returns a new page model&
func (l *masterLayout) RenderStep(p model.Page) (model.Page, error) {
	buffer := &bytes.Buffer{}
	err := l.render(buffer, &p)

	if err != nil {
		return model.Page{}, err
	}

	return model.Page{Title: p.Title, Body: template.HTML(buffer.String()), Data: p.Data, StyleSheets: p.StyleSheets, Scripts: p.Scripts}, nil
}
