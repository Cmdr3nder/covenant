package blog

import (
	"time"

	blogModels "github.com/ender4021/covenant/model/blog"
)

func retrieveVideoPosts() map[string]blogModels.Post {
	return nil
}

func retrieveAllPosts() map[string]blogModels.Post {
	return retrieveVideoPosts()
}

// RetrievePost gets the post with the given uuid
func RetrievePost(uuid string) blogModels.Post {
	return nil
}

// MonthPosts gets all posts for the given month in the given year
func MonthPosts(year int, month time.Month) []blogModels.Post {
	return nil
}

// RecentPosts the last "last" posts for the blog
func RecentPosts(last int) []blogModels.Post {
	return nil
}

// Years returns the years that have posts for this blog
func Years() []int {
	return nil
}

// Months returns the months that have posts for this blog in the given year
func Months(year int) []time.Month {
	return nil
}

// Context returns a context for the blog
func Context() blogModels.Blog {
	return blogModels.Blog{RecentPosts: RecentPosts(10)}
}
