package layout

import (
	"io"

	"github.com/ender4021/covenant/model/page"
)

// Layout is the Covenant layout mechanism
type Layout interface {
	Render(io.Writer, page.Page) error
	RenderStep(page.Page) (page.Page, error)
}
