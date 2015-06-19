package blog

import (
	"html/template"
	"time"

	"github.com/ender4021/covenant/model/page"
)

// VideoPost is a blog post that primarily focuses on some linked video content
type VideoPost struct {
}

// Date returns the date of the post
func (p *VideoPost) Date() time.Time {
	return time.Now()
}

// UUID returns the uuid of the post
func (p *VideoPost) UUID() string {
	return "a-special-youtube-video"
}

// Title returns the title of the post
func (p *VideoPost) Title() string {
	return "A Special YouTube Video"
}

// Data returns the extra data for the post
func (p *VideoPost) Data() interface{} {
	return &VideoPostData{IsYouTube: true, VideoID: "yo7frsh6wtI"}
}

// AsPage constructs a page instance that can be used with the post's layout
func (p *VideoPost) AsPage() page.Page {
	return page.Page{Title: template.HTMLAttr(p.Title()), Body: "This is a very special youtube video.", Data: p.Data()}
}

// LayoutID returns the string that represents the layout to use for the post
func (p *VideoPost) LayoutID() string {
	return "views_blog_video"
}

// VideoPostData is the extra information a page requires when rendering a video post
type VideoPostData struct {
	IsYouTube bool
	VideoID   string
}
