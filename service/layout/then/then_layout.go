package then

import (
	"io"

	"github.com/ender4021/covenant/model/page"
	"github.com/ender4021/covenant/service/layout"
)

// New returns a freshly constructed masterLayout
func New(first layout.Layout, second layout.Layout) layout.Layout {
	return &thenLayout{first: first, second: second}
}

type thenLayout struct {
	first  layout.Layout
	second layout.Layout
}

func (l *thenLayout) Render(w io.Writer, p page.Page) error {
	p, err := l.first.RenderStep(p)

	if err != nil {
		return err
	}

	return l.second.Render(w, p)
}

func (l *thenLayout) RenderStep(p page.Page) (page.Page, error) {
	p, err := l.first.RenderStep(p)

	if err != nil {
		return p, err
	}

	return l.second.RenderStep(p)
}
