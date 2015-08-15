package controller

import (
	"net/http"
	"strconv"
	"time"

	"github.com/ender4021/covenant/model"
	blogModels "github.com/ender4021/covenant/model/blog"
	"github.com/ender4021/covenant/model/page"
	"github.com/ender4021/covenant/service"
	"github.com/ender4021/covenant/service/blog"
	"github.com/ender4021/covenant/service/layout"
	"github.com/ender4021/covenant/service/layout/then"
	"github.com/ender4021/covenant/service/server"
)

// RegisterBlogController adds routes and initializes constants for routes controlled by the "Blog" controller
func RegisterBlogController(server server.Server) {
	path := service.GetRouteBuilder()

	path.AppendPart("blog")
	server.Get(path.MustCompile(), getBlogRoot)

	path.AppendPart("(?P<year>(19|20)[0-9]{2})")
	server.Get(path.MustCompile(), getBlogYear)

	path.AppendPart("(?P<month>(0[1-9])|(1[0-2]))")
	server.Get(path.MustCompile(), getBlogMonth)

	path.AppendPart("(?P<guid>([0-9]|[a-z]|[A-Z]|-|_)+)")
	server.Get(path.MustCompile(), getBlogPost)
}

func getCombinedBlogLayout() (layout.Layout, error) {
	blogLayout, err := service.GetLayout("views_blog_layout")

	if err != nil {
		return nil, err
	}

	rootLayout, err := service.GetRootLayout()

	if err != nil {
		return nil, err
	}

	return then.New(blogLayout, rootLayout), nil
}

func getBlogRoot(c model.Context, w http.ResponseWriter, r *http.Request) error {
	l, err := getCombinedBlogLayout()

	if err != nil {
		return err
	}

	archiveLayout, err := service.GetLayout("views_blog_archive")

	if err != nil {
		return err
	}

	years, err := blog.Years()

	if err != nil {
		return err
	}

	page, err := archiveLayout.RenderStep(page.Page{Data: years})

	if err != nil {
		return err
	}

	page.Title = "Andrew Bowers: Blog"
	page.Data, err = blog.Context()

	if err != nil {
		return err
	}

	return l.Render(w, page)
}

type blogYearView struct {
	Months  []time.Month
	Year    int
	IsValid bool
}

func getBlogYear(c model.Context, w http.ResponseWriter, r *http.Request) error {
	l, err := getCombinedBlogLayout()

	if err != nil {
		return err
	}

	yearLayout, err := service.GetLayout("views_blog_year")

	if err != nil {
		return err
	}

	year, err := strconv.Atoi(c.GetURLParam("year"))

	if err != nil {
		return err
	}

	months, err := blog.Months(year)

	if err != nil {
		return err
	}

	page, err := yearLayout.RenderStep(page.Page{Data: blogYearView{Year: year, Months: months, IsValid: len(months) > 0}})

	if err != nil {
		return err
	}

	page.Title = "Andrew Bowers: Blog"
	page.Data, err = blog.Context()

	if err != nil {
		return err
	}

	return l.Render(w, page)
}

type blogMonthView struct {
	Year    int
	Month   time.Month
	Posts   []blogModels.Post
	IsValid bool
}

func getBlogMonth(c model.Context, w http.ResponseWriter, r *http.Request) error {
	l, err := getCombinedBlogLayout()

	if err != nil {
		return err
	}

	monthLayout, err := service.GetLayout("views_blog_month")

	if err != nil {
		return err
	}

	year, err := strconv.Atoi(c.GetURLParam("year"))

	if err != nil {
		return err
	}

	month, err := strconv.Atoi(c.GetURLParam("month"))

	if err != nil {
		return err
	}

	monthPosts, err := blog.MonthPosts(year, time.Month(month))

	if err != nil {
		return err
	}

	page, err := monthLayout.RenderStep(page.Page{Data: blogMonthView{Year: year, Month: time.Month(month), Posts: monthPosts, IsValid: len(monthPosts) > 0}})

	if err != nil {
		return err
	}

	blogContext, err := blog.Context()

	if err != nil {
		return err
	}

	page.Title = "Andrew Bowers: Blog"
	page.Data = blogContext

	return l.Render(w, page)
}

func getBlogPost(c model.Context, w http.ResponseWriter, r *http.Request) error {
	l, err := getCombinedBlogLayout()

	if err != nil {
		return err
	}

	post, err := blog.RetrievePost(c.GetURLParam("guid"))

	//verify year and month match c.GetURLParam("year"), c.GetURLParam("month")

	postLayout, err := service.GetLayout(post.LayoutID())

	if err != nil {
		return err
	}

	page, err := postLayout.RenderStep(post.AsPage())

	if err != nil {
		return err
	}

	page.Title = "Andrew Bowers: Blog"
	page.Data, err = blog.Context()

	if err != nil {
		return err
	}

	return l.Render(w, page)
}
