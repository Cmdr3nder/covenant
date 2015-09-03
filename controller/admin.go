package controller

import (
	"fmt"
	"net/http"

	"github.com/ender4021/covenant/service/config"
	"github.com/ender4021/covenant/service/server"
	"github.com/goji/httpauth"
)

// RegisterAdminController sets up the authenticated pages
func RegisterAdminController(s server.Server, c config.Config) {
	admin := server.New()
	s.Handle("/admin/*", admin)
	admin.Use(httpauth.SimpleBasicAuth(c.GetString("admin_name"), c.GetString("admin_password")))

	admin.Get("/admin/ttt", ttt)
}

func ttt(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}
