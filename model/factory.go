package model

import (
	"github.com/ender4021/covenant/model/blog"
	"github.com/zenazn/goji/web"
)

// GetContext constructs a new Covenant Context for the given Goji Context
func GetContext(c web.C) Context {
	return &gojiContext{context: c}
}

// GetBlog constructs a new Blog context if not yet created, otherwise it returns the current blog context
func GetBlog() blog.Blog {
	var posts []blog.Post
	posts = append(posts, &blog.VideoPost{})
	return blog.Blog{RecentPosts: posts}
}
