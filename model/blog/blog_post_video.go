package blog

import (
	"html/template"
	"time"

	"github.com/ender4021/covenant/model/page"
)

// NewVideoPost builds a new video post from the given details
func NewVideoPost(date time.Time, uuid string, title string, comment string, videoID string, isYouTube bool) VideoPost {
	return VideoPost{PostedAt: date, Unique: uuid, Header: title, Text: comment, PostData: VideoPostData{IsYouTube: isYouTube, VideoID: videoID}}
}

// VideoPost is a blog post that primarily focuses on some linked video content
type VideoPost struct {
	PostedAt time.Time
	Unique   string
	Header   string
	PostData VideoPostData
	Text     string
}

// Date returns the date of the post
func (p *VideoPost) Date() time.Time {
	return p.PostedAt
}

// UUID returns the uuid of the post
func (p *VideoPost) UUID() string {
	return p.Unique
}

// Title returns the title of the post
func (p *VideoPost) Title() string {
	return p.Header
}

// Data returns the extra data for the post
func (p *VideoPost) Data() interface{} {
	return p.PostData
}

// AsPage constructs a page instance that can be used with the post's layout
func (p *VideoPost) AsPage() page.Page {
	return page.Page{Title: template.HTMLAttr(p.Header), Body: template.HTML(p.Text), Data: p.PostData}
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
