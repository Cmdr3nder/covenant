package model

import (
	"sort"
	"time"

	"github.com/ender4021/covenant/model/blog"
	"github.com/zenazn/goji/web"
)

// GetContext constructs a new Covenant Context for the given Goji Context
func GetContext(c web.C) Context {
	return &gojiContext{context: c}
}

func newLocalDay(year int, month time.Month, day int) time.Time {
	return time.Date(year, month, day, 0, 0, 0, 0, time.Local)
}

var allPosts = map[string]blog.Post{
	"2cellos-thunderstruck":           blog.NewVideoPost(newLocalDay(2014, time.June, 25), "2cellos-thunderstruck", "2CELLOS - Thunderstruck", "", "uT3SBzmDxGk", true),
	"the-expert":                      blog.NewVideoPost(newLocalDay(2014, time.May, 05), "the-expert", "The Expert", "", "BKorP55Aqvg", true),
	"the-emperor-voiced-by-the-joker": blog.NewVideoPost(newLocalDay(2014, time.February, 01), "the-emperor-voiced-by-the-joker", "The Emperor Voiced by Mark Hamill's Joker", "", "agcc7w8YmHo", true),
	"halo-odst-epilogue":              blog.NewVideoPost(newLocalDay(2013, time.October, 19), "halo-odst-epilogue", "Halo: ODST Epilogue", "Just finished Halo: ODST. This is the very cool epilogue.", "DZgTFUDSg2s", true),
	"camlistore-intro":                blog.NewVideoPost(newLocalDay(2013, time.September, 14), "camlistore-intro", "About Camlistore", "Camlistore is a cool idea. Unfortunately this video is really long&#8230;", "yxSzQIwXM1k", true),
	"attack-on-titan-theme":           blog.NewVideoPost(newLocalDay(2013, time.September, 14), "attack-on-titan-theme", "Attack on Titan (Original Opening Song)", "", "bnLgndWTBI0", true),
	"wrath-of-the-lich-king":          blog.NewVideoPost(newLocalDay(2013, time.September, 13), "wrath-of-the-lich-king", "Wrath of the Lich King", "Always enjoyed this trailer despite never getting around to playing WoW.", "BCr7y4SLhck", true),
	"call-of-duty-ghosts":             blog.NewVideoPost(newLocalDay(2013, time.September, 12), "call-of-duty-ghosts", "Call of Duty: Ghosts", "Being mainly a campaign player I am at least intrigued now. However this game is definitely not at the top of my 'to play' list.", "SumIZb6qMJw", true),
	"pumpktris":                       blog.NewVideoPost(newLocalDay(2013, time.September, 4), "pumpktris", "Pumpktris", "I wish I had the time to make something like this.", "8PCp5xk-9Qo", true),
	"man-at-arms-buster-sword":        blog.NewVideoPost(newLocalDay(2013, time.September, 4), "man-at-arms-buster-sword", "Cloud's Buster Sword - MAN AT ARMS", "", "xogheZdAO18", true),
	"man-at-arms-keyblade":            blog.NewVideoPost(newLocalDay(2013, time.September, 4), "man-at-arms-keyblade", "Sora's Keyblade - MAN AT ARMS", "", "bPH8Pe5G1p0", true),
	"man-at-arms-diamond-sword":       blog.NewVideoPost(newLocalDay(2013, time.September, 4), "man-at-arms-diamond-sword", "Diamond Sword (Minecraft) - MAN AT ARMS", "", "aNZRmvELxXM", true),
	"gigi-dagostino-bla-bla-bla":      blog.NewVideoPost(newLocalDay(2013, time.September, 4), "gigi-dagostino-bla-bla-bla", "Gigi D'Agostino - Bla Bla Bla", "", "g6t8g6ka4W0", true),
	"eve-online-ship-size":            blog.NewVideoPost(newLocalDay(2013, time.September, 4), "eve-online-ship-size", "Eve Online - Ship Size Comparison", "", "d8Ke1P3m4nU", true),
	"top-gear-piano":                  blog.NewVideoPost(newLocalDay(2013, time.September, 4), "top-gear-piano", "Top Gear Music on Piano - SNES", "", "_QVKcjpjeM4", true),
	"eve-online-origins":              blog.NewVideoPost(newLocalDay(2013, time.September, 4), "eve-online-origins", "Eve Online - Origins", "", "FZPCiqBLPM8", true),
	"why-x-stands-for-unknown":        blog.NewVideoPost(newLocalDay(2013, time.September, 4), "why-x-stands-for-unknown", "Why is 'X' the unknown?", "", "yo7frsh6wtI", true),
	"using-python-to-code-by-voice":   blog.NewVideoPost(newLocalDay(2013, time.September, 4), "using-python-to-code-by-voice", "Using Python to Code by Voice", "This is so cool. Hopefully he releases this tool soon so that when I go to do something similar I can just extend his solution instead of starting from scratch.", "8SkdfdXWYaI", true),
}

func allPostKeys() []string {
	var postKeys []string

	for k := range allPosts {
		postKeys = append(postKeys, k)
	}

	sort.Strings(postKeys)

	return postKeys
}

// AllPostYears returs a list of all years that have been posted in
func AllPostYears() []int {
	var years []int
	yearMap := make(map[int]bool)

	for _, post := range allPosts {
		year := post.Date().Year()
		if !yearMap[year] {
			yearMap[year] = true
			years = append(years, year)
		}
	}

	sort.Sort(sort.Reverse(sort.IntSlice(years)))

	return years
}

// GetBlog constructs a new Blog context if not yet created, otherwise it returns the current blog context
func GetBlog() blog.Blog {
	var posts []blog.Post

	for _, key := range allPostKeys() {
		posts = append(posts, allPosts[key])
	}

	return blog.Blog{RecentPosts: posts}
}

// GetPost constructs a complete post entry for the given uuid or an unfound post entry if uuid was unrecognized
func GetPost(uuid string) blog.Post {
	return allPosts[uuid]
}
