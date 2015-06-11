package master

import (
	"bytes"
	"html/template"
	"io"

	"github.com/ender4021/covenant/model"
	"github.com/ender4021/covenant/service/layout"
)

// New returns a freshly constructed masterLayout
func New(compiledLayout *template.Template) layout.Layout {
	return &masterLayout{compiledLayout: compiledLayout}
}

type masterLayout struct {
	compiledLayout *template.Template
}

func (l *masterLayout) Render(w io.Writer, p model.Page) error {
	return l.compiledLayout.Execute(w, p)
}

func (l *masterLayout) RenderStep(p model.Page) (model.Page, error) {
	buffer := &bytes.Buffer{}
	err := l.Render(buffer, p)

	if err != nil {
		return model.Page{}, err
	}

	return model.Page{Title: p.Title, Body: template.HTML(buffer.String()), Data: p.Data, StyleSheets: p.StyleSheets, Scripts: p.Scripts}, nil
}
