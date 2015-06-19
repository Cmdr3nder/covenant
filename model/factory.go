package model

import (
	"time"

	"github.com/zenazn/goji/web"
)

// GetContext constructs a new Covenant Context for the given Goji Context
func GetContext(c web.C) Context {
	return &gojiContext{context: c}
}

// GetBlog constructs a new Blog context if not yet created, otherwise it returns the current blog context
func GetBlog() Blog {
	var posts []BlogPost
	posts = append(posts, NewBlogPost(time.Now(), "simple-uuid", "title string"))
	return Blog{RecentPosts: posts}
}
