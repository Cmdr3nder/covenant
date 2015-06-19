package model

import (
	"time"

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
	posts = append(posts, blog.NewVideoPost(time.Now(), "a-special-youtube-video", "A Special YouTube Video", "This is a very special youtube video.", "yo7frsh6wtI", true))
	return blog.Blog{RecentPosts: posts}
}

// GetPost constructs a complete post entry for the given uuid or an unfound post entry if uuid was unrecognized
func GetPost(uuid string) blog.Post {
	return blog.NewVideoPost(time.Now(), "a-special-youtube-video", "A Special YouTube Video", "This is a very special youtube video.", "yo7frsh6wtI", true)
}
