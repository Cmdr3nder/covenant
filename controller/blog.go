package controller

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/ender4021/covenant/model"
	"github.com/ender4021/covenant/service"
	"github.com/ender4021/covenant/service/layout/then"
)

// RegisterBlogController adds routes and initializes constants for routes controlled by the "Blog" controller
func RegisterBlogController(server service.Server) {
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

func getBlogRoot(c model.Context, w http.ResponseWriter, r *http.Request) error {
	l := then.New(service.GetLayout("views_blog_layout"), service.GetRootLayout())
	page := model.Page{Title: "Andrew Bowers: Blog", Body: "Blog Root", Data: model.GetBlog()}

	return l.Render(w, page)
}

func getBlogYear(c model.Context, w http.ResponseWriter, r *http.Request) error {
	l := service.GetRootLayout()
	page := model.Page{Title: "Andrew Bowers: Blog", Body: template.HTML(fmt.Sprintf("Blog Year: %s", c.GetURLParam("year")))}

	return l.Render(w, page)
}

func getBlogMonth(c model.Context, w http.ResponseWriter, r *http.Request) error {
	l := service.GetRootLayout()
	page := model.Page{Title: "Andrew Bowers: Blog", Body: template.HTML(fmt.Sprintf("Blog Month: %s %s", c.GetURLParam("year"), c.GetURLParam("month")))}

	return l.Render(w, page)
}

func getBlogPost(c model.Context, w http.ResponseWriter, r *http.Request) error {
	l := service.GetRootLayout()
	page := model.Page{Title: "Andrew Bowers: Blog", Body: template.HTML(fmt.Sprintf("Blog Post: %s %s %s", c.GetURLParam("year"), c.GetURLParam("month"), c.GetURLParam("guid")))}

	return l.Render(w, page)
}
