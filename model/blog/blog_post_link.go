package blog

import (
	"html/template"
	"time"

	"github.com/ender4021/covenant/model/page"
)

// NewLinkPost builds a new video post from the given details
func NewLinkPost(date time.Time, uuid string, title string, comment string, href string, image string) LinkPost {
	return LinkPost{PostedAt: date, Unique: uuid, Header: title, Text: comment, PostData: LinkPostData{Address: href, Image: image}}
}

// LinkPost is a blog post that primarily focuses on linked article/page content
type LinkPost struct {
	PostedAt time.Time
	Unique   string
	Header   string
	PostData LinkPostData
	Text     string
}

// Date returns the date of the post
func (p *LinkPost) Date() time.Time {
	return p.PostedAt
}

// UUID returns the uuid of the post
func (p *LinkPost) UUID() string {
	return p.Unique
}

// Title returns the title of the post
func (p *LinkPost) Title() string {
	return p.Header
}

// Comment returns the text of the post
func (p *LinkPost) Comment() string {
	return p.Text
}

// Data returns the extra data for the post
func (p *LinkPost) Data() interface{} {
	return p.PostData
}

// AsPage constructs a page instance that can be used with the post's layout
func (p *LinkPost) AsPage() page.Page {
	return page.Page{Title: template.HTMLAttr(p.Header), Body: template.HTML(p.Text), Data: p.PostData}
}

// LayoutID returns the string that represents the layout to use for the post
func (p *LinkPost) LayoutID() string {
	return "views_blog_link"
}

// Type returns the type of post to be inserted into the DB
func (p *LinkPost) Type() string {
	return "link"
}

// LinkPostData is the extra information a page requires when rendering a video post
type LinkPostData struct {
	Address string
	Image   string
}
