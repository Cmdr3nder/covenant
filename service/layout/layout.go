package layout

import (
	"io"

	"github.com/ender4021/covenant/model"
)

// Layout is the Covenant layout mechanism
type Layout interface {
	Render(io.Writer, model.Page) error
	RenderStep(model.Page) (model.Page, error)
}
