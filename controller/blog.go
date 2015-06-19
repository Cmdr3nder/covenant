package controller

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/ender4021/covenant/model"
	"github.com/ender4021/covenant/model/page"
	"github.com/ender4021/covenant/service"
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

	path.AppendPart("(?P<month>(0[1-9])|(1[1-2]))")
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

	page := page.Page{Title: "Andrew Bowers: Blog", Body: "Blog Root", Data: model.GetBlog()}

	return l.Render(w, page)
}

func getBlogYear(c model.Context, w http.ResponseWriter, r *http.Request) error {
	l, err := getCombinedBlogLayout()

	if err != nil {
		return err
	}

	page := page.Page{Title: "Andrew Bowers: Blog", Body: template.HTML(fmt.Sprintf("Blog Year: %s", c.GetURLParam("year"))), Data: model.GetBlog()}

	return l.Render(w, page)
}

func getBlogMonth(c model.Context, w http.ResponseWriter, r *http.Request) error {
	l, err := getCombinedBlogLayout()

	if err != nil {
		return err
	}

	page := page.Page{Title: "Andrew Bowers: Blog", Body: template.HTML(fmt.Sprintf("Blog Month: %s %s", c.GetURLParam("year"), c.GetURLParam("month"))), Data: model.GetBlog()}

	return l.Render(w, page)
}

func getBlogPost(c model.Context, w http.ResponseWriter, r *http.Request) error {
	l, err := getCombinedBlogLayout()

	if err != nil {
		return err
	}

	post := model.GetPost(c.GetURLParam("guid"))

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
	page.Data = model.GetBlog()

	return l.Render(w, page)
}
