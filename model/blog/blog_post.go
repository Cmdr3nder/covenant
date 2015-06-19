package blog

import (
	"time"

	"github.com/ender4021/covenant/model/page"
)

// Post is the interface for a generic blog post
type Post interface {
	Date() time.Time
	UUID() string
	Title() string
	Data() interface{}
	AsPage() page.Page
	LayoutID() string
}
