package controller

import (
	"html/template"
	"net/http"
	"time"

	"github.com/ender4021/covenant/model"
	"github.com/ender4021/covenant/model/blog"
	"github.com/ender4021/covenant/model/page"
	"github.com/ender4021/covenant/service"
	blogService "github.com/ender4021/covenant/service/blog"
	"github.com/ender4021/covenant/service/config"
	"github.com/ender4021/covenant/service/server"
	"github.com/goji/httpauth"
)

// RegisterAdminController sets up the authenticated pages
func RegisterAdminController(s server.Server, c config.Config) {
	admin := server.NewMux()
	s.Handle("/admin/*", admin)
	admin.Use(httpauth.SimpleBasicAuth(c.GetString("admin_name"), c.GetString("admin_password")))

	admin.Get("/admin/blog", blogAdmin)
	admin.Get("/admin/blog/post/video", dumbGetter("views_admin_blog_video", "Create Video Post"))
	admin.Post("/admin/blog/post/video", createVideoSubmit)
	admin.Get("/admin/blog/post/link", dumbGetter("views_admin_blog_link", "Create Link Post"))
	admin.Post("/admin/blog/post/link", createLinkSubmit)
}

func blogAdmin(c model.Context, w http.ResponseWriter, r *http.Request) error {
	l, err := service.GetRootLayout()

	if err != nil {
		return err
	}

	adminBlogLayout, err := service.GetLayout("views_admin_blog_index")

	if err != nil {
		return err
	}

	page, err := adminBlogLayout.RenderStep(page.Page{})

	if err != nil {
		return err
	}

	page.Title = "Blog Administration"

	return l.Render(w, page)
}

func dumbGetter(layout string, title string) func(model.Context, http.ResponseWriter, *http.Request) error {
	return func(c model.Context, w http.ResponseWriter, r *http.Request) error {
		l, err := service.GetRootLayout()

		if err != nil {
			return err
		}

		videoLayout, err := service.GetLayout(layout)

		if err != nil {
			return err
		}

		page, err := videoLayout.RenderStep(page.Page{})

		if err != nil {
			return err
		}

		page.Title = template.HTMLAttr("Blog Administration: " + title)

		return l.Render(w, page)
	}
}

func createVideoSubmit(c model.Context, w http.ResponseWriter, r *http.Request) error {
	uuid := r.FormValue("uuid")
	title := r.FormValue("title")
	comment := r.FormValue("comment")
	address := r.FormValue("address")
	image := r.FormValue("image")

	videoPost := blog.NewLinkPost(time.Now(), uuid, title, comment, address, image)

	err := blogService.InsertPost(&videoPost)

	if err != nil {
		return err
	}

	http.Redirect(w, r, "/admin/blog", http.StatusFound)
	return nil
}

func createLinkSubmit(c model.Context, w http.ResponseWriter, r *http.Request) error {
	uuid := r.FormValue("uuid")
	title := r.FormValue("title")
	comment := r.FormValue("comment")
	videoID := r.FormValue("videoID")
	videoProvider := r.FormValue("videoProvider")

	videoPost := blog.NewVideoPost(time.Now(), uuid, title, comment, videoID, videoProvider)

	err := blogService.InsertPost(&videoPost)

	if err != nil {
		return err
	}

	http.Redirect(w, r, "/admin/blog", http.StatusFound)
	return nil
}
