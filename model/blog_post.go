package model

import "time"

// BlogPost is the interface for a generic blog post
type BlogPost interface {
	Date() time.Time
	UUID() string
	Title() string
}

// NewBlogPost constructs a new dumb blog post
func NewBlogPost(date time.Time, uuid string, title string) BlogPost {
	return &dumbBlogPost{date: date, uuid: uuid, title: title}
}

type dumbBlogPost struct {
	date  time.Time
	uuid  string
	title string
}

// Date gets the value of the blog post's date
func (p *dumbBlogPost) Date() time.Time {
	return p.date
}

// UUID gets the value of the blog post's uuid
func (p *dumbBlogPost) UUID() string {
	return p.uuid
}

// Title gets the value of the blog post's title
func (p *dumbBlogPost) Title() string {
	return p.title
}
